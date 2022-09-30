package usecase

import (
	"encoding/json"
	"math/rand"
	"time"

	"fake.com/padel-api/internal/matches"
	"fake.com/padel-api/internal/models"
	"fake.com/padel-api/internal/tournaments"
	"k8s.io/klog/v2"
)

type tournamentsUC struct {
	tournamentsRepo tournaments.Repository
	matchesRepo     matches.Repository
}

func NewTournamentsUseCase(tournamentsRepo tournaments.Repository, matchesRepo matches.Repository) tournaments.UseCase {
	return &tournamentsUC{tournamentsRepo: tournamentsRepo, matchesRepo: matchesRepo}
}

// Create news
func (u *tournamentsUC) Create(body []byte) (*models.Tournament, error) {
	klog.Info("# UseCaseTournmanet - Create")
	var tournamentJSON models.TournamentJSON
	json.Unmarshal(body, &tournamentJSON)

	tournament := models.NewTournament(tournamentJSON.Icon, tournamentJSON.Name, tournamentJSON.Description, tournamentJSON.Rounds, tournamentJSON.ActualRounds, createPlayersTForTournament(tournamentJSON.Players))

	t, err := u.tournamentsRepo.Create(tournament)

	if err != nil {
		klog.Error(err)
		return nil, err
	}

	createFirstRound(u, t)
	return t, nil
}

func (u *tournamentsUC) Update(body []byte, tournamentID int) (*models.Tournament, error) {
	klog.Info("# UseCaseTournmanet - Update")
	var tournamentJSON models.TournamentJSON
	json.Unmarshal(body, &tournamentJSON)

	tournament := models.NewTournament(tournamentJSON.Icon, tournamentJSON.Name, tournamentJSON.Description, tournamentJSON.Rounds, tournamentJSON.ActualRounds, createPlayersTForTournament(tournamentJSON.Players))

	return u.tournamentsRepo.Update(tournament, tournamentID)
}

func (u *tournamentsUC) Delete(tournamentID int) error {
	klog.Info("# UseCaseTournmanet - Delete")
	matches, _ := u.matchesRepo.GetMatchesByTournamentId(tournamentID)

	for _, m := range matches {
		u.matchesRepo.Delete(m.MatchId)
	}

	return u.tournamentsRepo.Delete(tournamentID)
}

func (u *tournamentsUC) GetTournaments() ([]models.Tournament, error) {
	return u.tournamentsRepo.GetTournaments()
}

func (u *tournamentsUC) GetTournament(playerID int) (*models.Tournament, error) {
	return u.tournamentsRepo.GetTournament(playerID)
}

func (u *tournamentsUC) NextRound(tournamentID int) (string, error) {
	klog.Info("# UseCaseTournmanet - NextRound")
	t, err := u.tournamentsRepo.GetTournament(tournamentID)

	if err != nil {
		klog.Error(err)
		return "Error resolving Tournament", err
	}

	tournamentAttrs := t.GetAttrs()

	if tournamentAttrs.ActualRounds == tournamentAttrs.Rounds {
		klog.Info("Last Round Created: Creating Results")
		// TODO Crear logica de creacion de resultados
		// TODO setear partidos a CLOSED
		return "Rounds Ended, Results Created", nil
	} else {
		tournamentAttrs.ActualRounds = tournamentAttrs.ActualRounds + 1
		//Verify all matches in this torunament are played
		matchesPending, err := u.matchesRepo.GetMatchesByTournamentIdAndStatus(tournamentID, "Pending")

		if err != nil {
			klog.Error(err)
			return "Error resolving tournament matches", err
		}

		matchesPlayed, err := u.matchesRepo.GetMatchesByTournamentIdAndStatus(tournamentID, "Played")

		if err != nil {
			klog.Error(err)
			return "Error resolving tournament matches", err
		}

		if len(matchesPending) == 0 {
			createNewRound(u, t)
			for _, m := range matchesPlayed {
				mAttrs := m.GetAttrs()
				mAttrs.Status = "Closed"
				m.SetAttrs(mAttrs)
				u.matchesRepo.Update(&m, m.MatchId)
			}
		} else {
			return "Matches Pending", nil
		}
	}

	return "Round Created", nil
}

func createPlayersTForTournament(playersId []int) []models.PlayerT {
	klog.Info("# UseCaseTournmanet - createPlayersTForTournament")
	var players []models.PlayerT
	for _, playerId := range playersId {
		players = append(players, *models.NewPlayerT(playerId))
	}

	return players
}

func createFirstRound(u *tournamentsUC, tournament *models.Tournament) error {
	klog.Info("# UseCaseTournmanet - createFirstRound")
	rand.Seed(time.Now().Unix())
	var randomIndex int
	var couples []models.PlayerT
	var auxPlayers []models.PlayerT
	tournamentAtrs := tournament.GetAttrs()
	matchesCount := len(tournamentAtrs.Players) / 4

	for i := 0; i < matchesCount; i++ {
		couples = couples[:0]
		for j := 0; j < 4; j++ {
			randomIndex = rand.Intn(len(tournamentAtrs.Players))
			p := tournamentAtrs.Players[randomIndex]
			tournamentAtrs.Players = append(tournamentAtrs.Players[:randomIndex], tournamentAtrs.Players[randomIndex+1:]...)
			couples = append(couples, p)
		}
		// create match: edit status playersT
		createMatchAndEdit(u, couples, tournament.TournamentID)
		auxPlayers = append(auxPlayers, couples...)
	}
	tournamentAtrs.Players = append(tournamentAtrs.Players, auxPlayers...)
	tournamentAtrs.ActualRounds = 1
	tournament.SetAttrs(tournamentAtrs)

	u.tournamentsRepo.Update(tournament, tournament.TournamentID)
	return nil
}

func createMatchAndEdit(u *tournamentsUC, couples []models.PlayerT, tournamentId int) {
	klog.Info("# UseCaseTournmanet - createMatchAndEdit")
	klog.Infof("Creating Match for Tournament(%d) with Couple(%d,%d) and Couple(%d,%d)",
		tournamentId,
		couples[0].PlayerID,
		couples[1].PlayerID,
		couples[2].PlayerID,
		couples[3].PlayerID)

	match := models.NewMatch(couples[0].PlayerID,
		couples[1].PlayerID,
		couples[2].PlayerID,
		couples[3].PlayerID,
		"Pending",
		tournamentId,
		models.Result{SetsCounter: 3,
			CoupleOneSets: []int{0, 0, 0},
			CoupleTwoSets: []int{0, 0, 0},
		})

	couples[0].Couples = append(couples[0].Couples, couples[1].PlayerID)
	couples[1].Couples = append(couples[1].Couples, couples[0].PlayerID)
	couples[2].Couples = append(couples[2].Couples, couples[3].PlayerID)
	couples[3].Couples = append(couples[3].Couples, couples[2].PlayerID)

	_, err := u.matchesRepo.Create(match)

	if err != nil {
		klog.Error(err)
	}
}

func createNewRound(u *tournamentsUC, t *models.Tournament) {
	klog.Info("# UseCaseTournmanet - createNewRound")
	rand.Seed(time.Now().Unix())
	var randomIndex int
	var p models.PlayerT
	tAttrs := t.GetAttrs()
	var couples []models.PlayerT
	var auxPlayers []models.PlayerT
	matchesCount := len(tAttrs.Players) / 4

	// Verificamos si todos han jugado las mismas rondas o algunos no han jugado esta ultima
	if tAttrs.FullRounds {
		klog.Info("CREATE ROUND: All Players have same rounds played")
		for i := 0; i < matchesCount; i++ {
			couples = couples[:0]
			for j := 0; j < 4; j++ {
				// Seteo de los jugadores sin mirar conflicto de parejas (son la primera indexacion de pareja, con lo que no hay restriccion)
				if j == 0 || j == 2 {
					randomIndex = rand.Intn(len(tAttrs.Players))
					p = tAttrs.Players[randomIndex]
					tAttrs.Players = append(tAttrs.Players[:randomIndex], tAttrs.Players[randomIndex+1:]...)
				} else {
					// Verify if that couple already play together and avoid that situation
					for k := 0; k < len(tAttrs.Players); k++ {
						p = tAttrs.Players[k]
						if !alreadyPlayTogether(couples[len(couples)-1].Couples, p.PlayerID) {
							tAttrs.Players = append(tAttrs.Players[:k], tAttrs.Players[k+1:]...)
							break
						}
					}
				}
				couples = append(couples, p)
			}
			// create match: edit status playersT
			createMatchAndEdit(u, couples, t.TournamentID)
			auxPlayers = append(auxPlayers, couples...)
		}
	} else {
		klog.Info("CREATE ROUND: Some Players have less rounds played")
		playersWithLessRoundsPlayed := getPlayersWithLessRounds(tAttrs.Players, tAttrs.ActualRounds)

		for i := 0; i < matchesCount; i++ {
			couples = couples[:0]
			for j := 0; j < 4; j++ {
				if j == 0 || j == 2 {
					if len(playersWithLessRoundsPlayed) > 0 {
						randomIndex = rand.Intn(len(playersWithLessRoundsPlayed))
						p = playersWithLessRoundsPlayed[randomIndex]
						index := getIndexFromOriginPlayersSlice(tAttrs.Players, p)
						tAttrs.Players = append(tAttrs.Players[:index], tAttrs.Players[index+1:]...)
						playersWithLessRoundsPlayed = append(playersWithLessRoundsPlayed[:randomIndex], playersWithLessRoundsPlayed[randomIndex+1:]...)
					} else {
						randomIndex = rand.Intn(len(tAttrs.Players))
						p = tAttrs.Players[randomIndex]
						tAttrs.Players = append(tAttrs.Players[:randomIndex], tAttrs.Players[randomIndex+1:]...)
					}
					couples = append(couples, p)
				} else {
					foundCouple := false
					for k := 0; k < len(playersWithLessRoundsPlayed); k++ {
						p = playersWithLessRoundsPlayed[k]
						if !alreadyPlayTogether(couples[len(couples)-1].Couples, p.PlayerID) {
							playersWithLessRoundsPlayed = append(playersWithLessRoundsPlayed[:k], playersWithLessRoundsPlayed[k+1:]...)
							index := getIndexFromOriginPlayersSlice(tAttrs.Players, p)
							tAttrs.Players = append(tAttrs.Players[:index], tAttrs.Players[index+1:]...)
							foundCouple = true
							break
						}
					}

					if !foundCouple {
						for k := 0; k < len(tAttrs.Players); k++ {
							p = tAttrs.Players[k]
							if !alreadyPlayTogether(couples[len(couples)-1].Couples, p.PlayerID) {
								tAttrs.Players = append(tAttrs.Players[:k], tAttrs.Players[k+1:]...)
								foundCouple = true
								break
							}
						}
					}

					couples = append(couples, p)
				}
			}
			// create match: edit status playersT
			createMatchAndEdit(u, couples, t.TournamentID)
			auxPlayers = append(auxPlayers, couples...)
		}
	}

	tAttrs.Players = append(tAttrs.Players, auxPlayers...)
	tAttrs.ActualRounds++
	t.SetAttrs(tAttrs)

	u.tournamentsRepo.Update(t, t.TournamentID)
}

func alreadyPlayTogether(couples []int, playerId int) bool {
	klog.Info("# UseCaseTournmanet - alreadyPlayTogether")
	for _, id := range couples {
		if id == playerId {
			return true
		}
	}

	return false
}

func getPlayersWithLessRounds(players []models.PlayerT, roundsPlayed int) []models.PlayerT {
	klog.Info("# UseCaseTournmanet - getPlayersWithLessRounds")
	var playersWithLessRoundsPlayed []models.PlayerT

	for _, p := range players {
		if p.RoundsPlayed != roundsPlayed {
			playersWithLessRoundsPlayed = append(playersWithLessRoundsPlayed, p)
		}
	}

	return playersWithLessRoundsPlayed
}

func getIndexFromOriginPlayersSlice(players []models.PlayerT, p models.PlayerT) int {
	klog.Info("# UseCaseTournmanet - getIndexFromOriginPlayersSlice")
	for index, plyr := range players {
		if p.PlayerID == plyr.PlayerID {
			klog.Info("ENDING # UseCaseTournmanet - getIndexFromOriginPlayersSlice")
			return index
		}
	}

	return -1
}

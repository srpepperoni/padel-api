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
	var tournamentJSON models.TournamentJSON
	json.Unmarshal(body, &tournamentJSON)

	tournament := models.NewTournament(tournamentJSON.Icon, tournamentJSON.Name, tournamentJSON.Description, tournamentJSON.Rounds, tournamentJSON.ActualRounds, createPlayersTForTournament(tournamentJSON.Players))

	return u.tournamentsRepo.Update(tournament, tournamentID)
}

func (u *tournamentsUC) Delete(tournamentID int) error {
	return u.tournamentsRepo.Delete(tournamentID)
}

func (u *tournamentsUC) GetTournaments() ([]models.Tournament, error) {
	return u.tournamentsRepo.GetTournaments()
}

func (u *tournamentsUC) GetTournament(playerID int) (*models.Tournament, error) {
	return u.tournamentsRepo.GetTournament(playerID)
}

func (u *tournamentsUC) NextRound(tournamentID int) (string, error) {
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
		matchesPlayed, err := u.matchesRepo.GetMatchesByTournamentIdAndStatus(tournamentID, "Played")

		if err != nil {
			klog.Error(err)
			return "Error resolving tournament matches", err
		}

		if len(*matchesPending) == 0 {
			createNewRound(u, t)
			for _, m := range *matchesPlayed {
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
	var players []models.PlayerT
	for _, playerId := range playersId {
		players = append(players, *models.NewPlayerT(playerId))
	}

	return players
}

func createFirstRound(u *tournamentsUC, tournament *models.Tournament) error {
	rand.Seed(time.Now().Unix())
	var randomIndex int
	var couples []models.PlayerT
	var auxPlayers []models.PlayerT
	tournamentAtrs := tournament.GetAttrs()
	matchesCount := len(tournamentAtrs.Players) / 4

	players := tournamentAtrs.Players
	for i := 0; i < matchesCount; i++ {
		couples = couples[:0]
		for j := 0; j < 4; j++ {
			randomIndex = rand.Intn(len(players))
			p := players[randomIndex]
			players = append(players[:randomIndex], players[randomIndex+1:]...)
			couples = append(couples, p)
		}
		// create match: edit status playersT
		createMatchAndEdit(u, couples, tournament.TournamentID)
		auxPlayers = append(auxPlayers, couples...)
	}
	players = append(players, auxPlayers...)
	tournamentAtrs.ActualRounds = 1
	tournament.SetAttrs(tournamentAtrs)

	u.tournamentsRepo.Update(tournament, tournament.TournamentID)
	return nil
}

func createMatchAndEdit(u *tournamentsUC, couples []models.PlayerT, tournamentId int) {
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
	rand.Seed(time.Now().Unix())
	var randomIndex int
	var randomIndexAux int
	var limitLoop int
	var firstArrayFinished = false
	tAttrs := t.GetAttrs()
	var couples []models.PlayerT
	var auxPlayers []models.PlayerT
	matchesCount := len(tAttrs.Players) / 4

	// Verificamos si todos han jugado las mismas rondas o algunos no han jugado esta ultima
	if haveAllPlayersSameRounds(tAttrs) {
		klog.Info("CREATE ROUND: All Players have same rounds played")
		for i := 0; i < matchesCount; i++ {
			couples = couples[:0]
			// Bucle de creacion de parejas por partido
			for j := 0; j < 4; j++ {
				// Seteo de los jugadores sin mirar conflicto de parejas (son la primera indexacion de pareja, con lo que no hay restriccion)
				if j == 0 || j == 2 {
					randomIndex = rand.Intn(len(tAttrs.Players))
					p := tAttrs.Players[randomIndex]
					tAttrs.Players = append(tAttrs.Players[:randomIndex], tAttrs.Players[randomIndex+1:]...)
					couples = append(couples, p)
				} else {
					// Verify if that couple already play together and avoid that situation
					randomIndex = rand.Intn(len(tAttrs.Players))
					p := tAttrs.Players[randomIndex]
					for alreadyPlayTogether(couples[len(couples)-1].Couples, p.PlayerID) {
						randomIndex++
						if randomIndex > len(tAttrs.Players)-1 {
							randomIndex = 0
						}
						p = tAttrs.Players[randomIndex]
					}
					tAttrs.Players = append(tAttrs.Players[:randomIndex], tAttrs.Players[randomIndex+1:]...)
					couples = append(couples, p)
				}
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
						p := playersWithLessRoundsPlayed[randomIndex]
						index := getIndexFromOriginPlayersSlice(tAttrs.Players, p)
						if index < 0 {
							klog.Error("Player not Found")
							return
						}
						tAttrs.Players = append(tAttrs.Players[:index], tAttrs.Players[index+1:]...)
						playersWithLessRoundsPlayed = append(playersWithLessRoundsPlayed[:randomIndex], playersWithLessRoundsPlayed[randomIndex+1:]...)
						couples = append(couples, p)
					} else {
						randomIndex = rand.Intn(len(tAttrs.Players))
						p := tAttrs.Players[randomIndex]
						tAttrs.Players = append(tAttrs.Players[:randomIndex], tAttrs.Players[randomIndex+1:]...)
						couples = append(couples, p)
					}
				} else {
					if len(playersWithLessRoundsPlayed) > 0 {
						randomIndexAux = rand.Intn(len(playersWithLessRoundsPlayed))
						randomIndex = rand.Intn(len(tAttrs.Players))
						p := playersWithLessRoundsPlayed[randomIndex]
						limitLoop = randomIndexAux
						firstArrayFinished = false
						for alreadyPlayTogether(couples[len(couples)-1].Couples, p.PlayerID) {
							if firstArrayFinished == false {
								randomIndexAux++

								if randomIndexAux > len(playersWithLessRoundsPlayed)-1 {
									randomIndexAux = 0
								}

								if randomIndexAux == limitLoop {
									firstArrayFinished = true
								} else {
									p = playersWithLessRoundsPlayed[randomIndexAux]
								}
							} else {
								randomIndex++
								if randomIndex > len(tAttrs.Players)-1 {
									randomIndex = 0
								}
								p = tAttrs.Players[randomIndex]
								if couples[len(couples)-1].PlayerID == p.PlayerID {
									randomIndex++
									p = tAttrs.Players[randomIndex]
								}
							}
						}

						playersWithLessRoundsPlayed = append(playersWithLessRoundsPlayed[:randomIndexAux], playersWithLessRoundsPlayed[randomIndexAux+1:]...)
						tAttrs.Players = append(tAttrs.Players[:randomIndex], tAttrs.Players[randomIndex+1:]...)
						couples = append(couples, p)
					} else {
						// Verify if that couple already play together and avoid that situation
						randomIndex = rand.Intn(len(tAttrs.Players))
						p := tAttrs.Players[randomIndex]
						for alreadyPlayTogether(couples[len(couples)-1].Couples, p.PlayerID) {
							randomIndex++
							if randomIndex > len(tAttrs.Players)-1 {
								randomIndex = 0
							}
							p = tAttrs.Players[randomIndex]
						}
						tAttrs.Players = append(tAttrs.Players[:randomIndex], tAttrs.Players[randomIndex+1:]...)
						couples = append(couples, p)
					}
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

func haveAllPlayersSameRounds(attrs *models.TournamentAttrs) bool {
	for _, p := range attrs.Players {
		klog.Info(p.RoundsPlayed)
		if p.RoundsPlayed != attrs.ActualRounds {
			return false
		}
	}

	return true
}

func alreadyPlayTogether(couples []int, playerId int) bool {
	for _, id := range couples {
		if id == playerId {
			return true
		}
	}

	return false
}

func getPlayersWithLessRounds(players []models.PlayerT, roundsPlayed int) []models.PlayerT {
	var playersWithLessRoundsPlayed []models.PlayerT

	for _, p := range players {
		if p.RoundsPlayed != roundsPlayed {
			playersWithLessRoundsPlayed = append(playersWithLessRoundsPlayed, p)
		}
	}

	return playersWithLessRoundsPlayed
}

func getIndexFromOriginPlayersSlice(players []models.PlayerT, p models.PlayerT) int {
	for index, plyr := range players {
		if p.PlayerID == plyr.PlayerID {
			return index
		}
	}
	return -1
}

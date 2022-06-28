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

	createRound(u, t)

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

func createPlayersTForTournament(playersId []int) []models.PlayerT {
	var players []models.PlayerT
	for _, playerId := range playersId {
		players = append(players, *models.NewPlayerT(playerId))
	}

	return players
}

func createRound(u *tournamentsUC, tournament *models.Tournament) error {
	rand.Seed(time.Now().Unix())
	var randomIndex int
	var couples []models.PlayerT
	var auxPlayers []models.PlayerT
	tournamentAtrs := tournament.GetAttrs()
	matchesCount := len(tournamentAtrs.Players) / 4

	if tournamentAtrs.ActualRounds == 0 {
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
	} else {
		// TODO algoritmo de emparejamiento TOP
		klog.Info("dasdasdf")
	}
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

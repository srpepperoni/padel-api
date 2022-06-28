package usecase

import (
	"text/template"

	"fake.com/padel-api/internal/matches"
	"fake.com/padel-api/internal/models"
	"fake.com/padel-api/internal/players"
	"fake.com/padel-api/internal/templates"
	"fake.com/padel-api/internal/tournaments"
	"k8s.io/klog/v2"
)

const (
	basePrefix = "api-players:"
)

type templatesUC struct {
	playersRepo     players.Repository
	matchesRepo     matches.Repository
	tournamentsRepo tournaments.Repository
}

func NewTemplatesUseCase(playersRepo players.Repository, matchesRepo matches.Repository, tournamentsRepo tournaments.Repository) templates.UseCase {
	return &templatesUC{playersRepo: playersRepo, matchesRepo: matchesRepo, tournamentsRepo: tournamentsRepo}
}

func (u *templatesUC) GetTemplate() (*template.Template, error) {
	t := template.Must(template.ParseFiles("./internal/templates/resources/index.html"))
	return t, nil
}

func (u *templatesUC) GetTemplateNewUser() (*template.Template, []models.PlayerJSON, error) {
	players, err := u.playersRepo.GetPlayers()
	var playerJson []models.PlayerJSON

	if err != nil {
		klog.Error(err)
	}

	for _, p := range players {
		playerJson = append(playerJson, *p.ToPlayerJSON())
	}

	t := template.Must(template.ParseFiles("./internal/templates/resources/new-user.html"))
	return t, playerJson, nil
}

func (u *templatesUC) GetTemplateNewMatch() (*template.Template, []models.MatchForTemplate, error) {
	var matchesForTemplate []models.MatchForTemplate

	matches, err := u.matchesRepo.GetMatches()

	if err != nil {
		klog.Error(err)
	}

	for _, m := range matches {
		matchAttrs := m.GetAttrs()

		t, err := u.tournamentsRepo.GetTournament(matchAttrs.TournamentID)
		if err != nil {
			klog.Error(err)
		}

		tournamentAttrs := t.GetAttrs()

		coupleOne, coupleTwo := getCouplesAsPlayerJSONArray(u, matchAttrs.CoupleOne, matchAttrs.CoupleTwo)

		matchesForTemplate = append(matchesForTemplate, models.MatchForTemplate{TournamentName: tournamentAttrs.Name,
			Status:    matchAttrs.Status,
			Result:    matchAttrs.Result,
			CoupleOne: coupleOne,
			CoupleTwo: coupleTwo,
		})
	}

	t := template.Must(template.ParseFiles("./internal/templates/resources/new-match.html"))

	return t, matchesForTemplate, nil
}

func (u *templatesUC) GetTemplateNewTournament() (*template.Template, []models.PlayerJSON, error) {

	players, err := u.playersRepo.GetPlayers()
	var playerJson []models.PlayerJSON

	if err != nil {
		klog.Error(err)
	}

	for _, p := range players {
		playerJson = append(playerJson, *p.ToPlayerJSON())
	}

	t := template.Must(template.ParseFiles("./internal/templates/resources/new-tournaments.html"))
	return t, playerJson, nil
}

func getCouplesAsPlayerJSONArray(u *templatesUC, coupleOne []int, coupleTwo []int) ([]models.PlayerJSON, []models.PlayerJSON) {

	var CoupleOne []models.PlayerJSON
	var CoupleTwo []models.PlayerJSON

	if p, err := u.playersRepo.GetPlayer(coupleOne[0]); err != nil {
		klog.Error(err)
	} else {
		playerAttrs := p.GetAttrs()
		CoupleOne = append(CoupleOne, models.PlayerJSON{PlayerName: playerAttrs.PlayerName, Name: playerAttrs.Name, LastName: playerAttrs.LastName})
	}

	if p, err := u.playersRepo.GetPlayer(coupleOne[1]); err != nil {
		klog.Error(err)
	} else {
		playerAttrs := p.GetAttrs()
		CoupleOne = append(CoupleOne, models.PlayerJSON{PlayerName: playerAttrs.PlayerName, Name: playerAttrs.Name, LastName: playerAttrs.LastName})
	}

	if p, err := u.playersRepo.GetPlayer(coupleTwo[0]); err != nil {
		klog.Error(err)
	} else {
		playerAttrs := p.GetAttrs()
		CoupleTwo = append(CoupleOne, models.PlayerJSON{PlayerName: playerAttrs.PlayerName, Name: playerAttrs.Name, LastName: playerAttrs.LastName})
	}

	if p, err := u.playersRepo.GetPlayer(coupleTwo[1]); err != nil {
		klog.Error(err)
	} else {
		playerAttrs := p.GetAttrs()
		CoupleTwo = append(CoupleOne, models.PlayerJSON{PlayerName: playerAttrs.PlayerName, Name: playerAttrs.Name, LastName: playerAttrs.LastName})
	}

	return CoupleOne, CoupleTwo
}

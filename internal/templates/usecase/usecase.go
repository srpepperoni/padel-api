package usecase

import (
	"fmt"
	"text/template"

	"fake.com/padel-api/internal/matches"
	"fake.com/padel-api/internal/models"
	"fake.com/padel-api/internal/players"
	"fake.com/padel-api/internal/templates"
)

const (
	basePrefix = "api-players:"
)

type templatesUC struct {
	playersRepo players.Repository
	matchesRepo matches.Repository
}

func NewTemplatesUseCase(playersRepo players.Repository, matchesRepo matches.Repository) templates.UseCase {
	return &templatesUC{playersRepo: playersRepo, matchesRepo: matchesRepo}
}

func (u *templatesUC) GetTemplate() (*template.Template, error) {
	t := template.Must(template.ParseFiles("./internal/templates/resources/index.html"))
	return t, nil
}

func (u *templatesUC) GetTemplateNewUser() (*template.Template, error) {
	t := template.Must(template.ParseFiles("./internal/templates/resources/new-user.html"))
	return t, nil
}

func (u *templatesUC) GetTemplateNewMatch() (*template.Template, *models.MatchesAndPlayers, error) {

	var players *[]models.Player
	var matches *[]models.Match
	var err error

	if players, err = u.playersRepo.GetPlayers(); err != nil {
		fmt.Println(err)
	}

	u.matchesRepo.GetMatches()

	if matches, err = u.matchesRepo.GetMatches(); err != nil {
		fmt.Println(err)
	}

	var matchesAndPlayers = models.MatchesAndPlayers{Players: *players, Matches: *matches}

	t := template.Must(template.ParseFiles("./internal/templates/resources/new-match.html"))

	return t, &matchesAndPlayers, nil
}

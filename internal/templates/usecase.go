package templates

import (
	"text/template"

	"fake.com/padel-api/internal/models"
)

type UseCase interface {
	GetTemplate() (*template.Template, error)
	GetTemplateNewUser() (*template.Template, []models.PlayerJSON, error)
	GetTemplateNewMatch() (*template.Template, *models.MatchesAndPlayers, error)
	GetTemplateNewTournament() (*template.Template, []models.PlayerJSON, error)
}

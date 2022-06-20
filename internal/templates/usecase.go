package templates

import (
	"text/template"

	"fake.com/padel-api/internal/models"
)

type UseCase interface {
	GetTemplate() (*template.Template, error)
	GetTemplateNewUser() (*template.Template, error)
	GetTemplateNewMatch() (*template.Template, *models.MatchesAndPlayers, error)
}

package http

import (
	"net/http"

	"fake.com/padel-api/internal/templates"
)

type templatesHandlers struct {
	templatesUC templates.UseCase
}

func NewTemplatesHandlers(templatesUC templates.UseCase) templates.Handlers {
	return &templatesHandlers{templatesUC: templatesUC}
}

func (h templatesHandlers) GetTemplate(w http.ResponseWriter, r *http.Request) {
	t, err := h.templatesUC.GetTemplate()
	err = t.Execute(w, nil)

	if err != nil {
		panic(err)
	}
}

func (h templatesHandlers) GetTemplateNewUser(w http.ResponseWriter, r *http.Request) {
	t, err := h.templatesUC.GetTemplateNewUser()
	err = t.Execute(w, nil)

	if err != nil {
		panic(err)
	}
}

func (h templatesHandlers) GetTemplateNewMatch(w http.ResponseWriter, r *http.Request) {
	t, matchesAndPlayers, err := h.templatesUC.GetTemplateNewMatch()

	err = t.Execute(w, matchesAndPlayers)

	if err != nil {
		panic(err)
	}
}

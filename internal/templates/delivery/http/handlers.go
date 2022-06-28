package http

import (
	"net/http"

	"fake.com/padel-api/internal/templates"
	"k8s.io/klog/v2"
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
	t, p, err := h.templatesUC.GetTemplateNewUser()
	err = t.Execute(w, p)

	if err != nil {
		panic(err)
	}
}

func (h templatesHandlers) GetTemplateNewMatch(w http.ResponseWriter, r *http.Request) {
	t, matchesForTemplate, err := h.templatesUC.GetTemplateNewMatch()

	err = t.Execute(w, matchesForTemplate)

	if err != nil {
		panic(err)
	}
}

func (h templatesHandlers) GetTemplateNewTournament(w http.ResponseWriter, r *http.Request) {
	t, Players, err := h.templatesUC.GetTemplateNewTournament()
	klog.Info(Players)
	err = t.Execute(w, Players)

	if err != nil {
		panic(err)
	}
}

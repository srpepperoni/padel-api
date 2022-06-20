package http

import (
	"net/http"

	"fake.com/padel-api/internal/templates"
	"github.com/gorilla/mux"
)

func MapTemplatesRoutes(router *mux.Router, templatesHandlers templates.Handlers) {
	router.HandleFunc("/", templatesHandlers.GetTemplate).Methods(http.MethodGet)
	router.HandleFunc("/new-user", templatesHandlers.GetTemplateNewUser).Methods(http.MethodGet)
	router.HandleFunc("/new-match", templatesHandlers.GetTemplateNewMatch).Methods(http.MethodGet)
}

package templates

import (
	"net/http"
)

type Handlers interface {
	GetTemplate(w http.ResponseWriter, r *http.Request)
	GetTemplateNewUser(w http.ResponseWriter, r *http.Request)
	GetTemplateNewMatch(w http.ResponseWriter, r *http.Request)
}

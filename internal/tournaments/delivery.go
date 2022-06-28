package tournaments

import (
	"net/http"
)

type Handlers interface {
	Create(w http.ResponseWriter, r *http.Request)
	Update(w http.ResponseWriter, r *http.Request)
	Delete(w http.ResponseWriter, r *http.Request)
	GetTournament(w http.ResponseWriter, r *http.Request)
	GetTournaments(w http.ResponseWriter, r *http.Request)
	NextRound(w http.ResponseWriter, r *http.Request)
}

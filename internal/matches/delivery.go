package matches

import (
	"net/http"
)

type Handlers interface {
	Create(w http.ResponseWriter, r *http.Request)
	Update(w http.ResponseWriter, r *http.Request)
	Delete(w http.ResponseWriter, r *http.Request)
	GetMatch(w http.ResponseWriter, r *http.Request)
	GetMatches(w http.ResponseWriter, r *http.Request)
	GetMatchesByTournamentId(w http.ResponseWriter, r *http.Request)
	SetResult(w http.ResponseWriter, r *http.Request)
}

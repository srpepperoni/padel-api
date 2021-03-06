package players

import (
	"net/http"
)

type Handlers interface {
	Create(w http.ResponseWriter, r *http.Request)
	Update(w http.ResponseWriter, r *http.Request)
	Delete(w http.ResponseWriter, r *http.Request)
	GetPlayer(w http.ResponseWriter, r *http.Request)
	GetPlayers(w http.ResponseWriter, r *http.Request)
}

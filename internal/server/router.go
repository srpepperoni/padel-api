package router

import (
	"net/http"

	datastore "fake.com/padel-api/pkg/db/postgres"
	"github.com/gorilla/mux"

	playersHttp "fake.com/padel-api/internal/players/delivery/http"
	playersRepository "fake.com/padel-api/internal/players/repository"
	playersUseCase "fake.com/padel-api/internal/players/usecase"

	matchesHttp "fake.com/padel-api/internal/matches/delivery/http"
	matchesRepository "fake.com/padel-api/internal/matches/repository"
	matchesUseCase "fake.com/padel-api/internal/matches/usecase"

	tournamentsHttp "fake.com/padel-api/internal/tournaments/delivery/http"
	tournamentsRepository "fake.com/padel-api/internal/tournaments/repository"
	tournamentsUseCase "fake.com/padel-api/internal/tournaments/usecase"

	templatesHttp "fake.com/padel-api/internal/templates/delivery/http"
	templatesUseCase "fake.com/padel-api/internal/templates/usecase"
)

func NewRouter() *mux.Router {
	db, err := datastore.NewPostgresDB()
	if err != nil {
		panic(err)
	}
	router := mux.NewRouter()

	//Init repositories
	playersRepo := playersRepository.NewPlayersRepository(db)
	matchesRepo := matchesRepository.NewMatchesRepository(db)
	tournamentsRepo := tournamentsRepository.NewTournamentsRepository(db)

	//Init useCases
	playersUC := playersUseCase.NewPlayersUseCase(playersRepo)
	matchesUC := matchesUseCase.NewMatchsUseCase(matchesRepo)
	templatesUC := templatesUseCase.NewTemplatesUseCase(playersRepo, matchesRepo)
	tournamentsUC := tournamentsUseCase.NewTournamentsUseCase(tournamentsRepo)

	//Init Handlers
	playersHandlers := playersHttp.NewPlayersHandlers(playersUC)
	matchesHandlers := matchesHttp.NewMatchesHandlers(matchesUC)
	templatesHandlers := templatesHttp.NewTemplatesHandlers(templatesUC)
	tournamentsHandlers := tournamentsHttp.NewTournamentsHandlers(tournamentsUC)

	playersHttp.MapPlayersRoutes(router, playersHandlers)
	matchesHttp.MapMatchesRoutes(router, matchesHandlers)
	templatesHttp.MapTemplatesRoutes(router, templatesHandlers)
	tournamentsHttp.MapTournamentsRoutes(router, tournamentsHandlers)

	css := "/internal/templates/resources/css"
	js := "/internal/templates/resources/js"
	router.PathPrefix(css).Handler(http.StripPrefix(css, http.FileServer(http.Dir("."+css))))
	router.PathPrefix(js).Handler(http.StripPrefix(js, http.FileServer(http.Dir("."+js))))

	return router
}

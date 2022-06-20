package router

import (
	datastore "fake.com/padel-api/pkg/db/postgres"
	"github.com/gorilla/mux"

	playersHttp "fake.com/padel-api/internal/players/delivery/http"
	playersRepository "fake.com/padel-api/internal/players/repository"
	playersUseCase "fake.com/padel-api/internal/players/usecase"

	matchesHttp "fake.com/padel-api/internal/matches/delivery/http"
	matchesRepository "fake.com/padel-api/internal/matches/repository"
	matchesUseCase "fake.com/padel-api/internal/matches/usecase"

	templatesHttp "fake.com/padel-api/internal/templates/delivery/http"
	templatesUseCase "fake.com/padel-api/internal/templates/usecase"
)

func NewRouter() *mux.Router {
	DB, _ := datastore.NewPostgresDB()
	router := mux.NewRouter()

	//Init repositories
	playersRepo := playersRepository.NewPlayersRepository(DB)
	matchesRepo := matchesRepository.NewMatchesRepository(DB)

	//Init useCases
	playersUC := playersUseCase.NewPlayersUseCase(playersRepo)
	matchesUC := matchesUseCase.NewMatchsUseCase(matchesRepo)
	templatesUC := templatesUseCase.NewTemplatesUseCase(playersRepo, matchesRepo)

	//Init Handlers
	playersHandlers := playersHttp.NewPlayersHandlers(playersUC)
	matchesHandlers := matchesHttp.NewMatchesHandlers(matchesUC)
	templatesHandlers := templatesHttp.NewTemplatesHandlers(templatesUC)

	playersHttp.MapPlayersRoutes(router, playersHandlers)
	matchesHttp.MapMatchesRoutes(router, matchesHandlers)
	templatesHttp.MapTemplatesRoutes(router, templatesHandlers)

	return router
}

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

	httpSwagger "github.com/swaggo/http-swagger" // http-swagger middleware
)

func NewRouter() *mux.Router {
	DB, _ := datastore.NewPostgresDB()
	router := mux.NewRouter()

	//Init repositories
	playersRepo := playersRepository.NewPlayersRepository(DB)
	matchesRepo := matchesRepository.NewMatchesRepository(DB)
	tournamentsRepo := tournamentsRepository.NewTournamentsRepository(DB)

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

	router.PathPrefix("/swagger").HandlerFunc(swaggerHandler) // swagger config

	router.PathPrefix("/docs").Handler(http.StripPrefix("/docs", http.FileServer(http.Dir("./docs")))) //swagger config statics path
	router.PathPrefix("/internal/templates/resources/css").Handler(http.StripPrefix("/internal/templates/resources/css", http.FileServer(http.Dir("./internal/templates/resources/css"))))
	router.PathPrefix("/internal/templates/resources/js").Handler(http.StripPrefix("/internal/templates/resources/js", http.FileServer(http.Dir("./internal/templates/resources/js"))))

	return router
}

// setting new url for swagger (default is docs.json)
func swaggerHandler(w http.ResponseWriter, r *http.Request) {
	swaggerFileUrl := "http://localhost:8000/docs/swagger.json"
	handler := httpSwagger.Handler(httpSwagger.URL(swaggerFileUrl))
	handler.ServeHTTP(w, r)
}

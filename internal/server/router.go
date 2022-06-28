package server

import (
	"net/http"

	"fake.com/padel-api/config"
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

const (
	docsPath = "/docs"
	docsDir  = "./docs"
	cssPath  = "/internal/templates/resources/css"
	cssDir   = "./internal/templates/resources/css"
	jsPath   = "/internal/templates/resources/js"
	jsDir    = "./internal/templates/resources/js"
)

func NewRouter(cfg *config.Config) *mux.Router {
	db, _ := datastore.NewPostgresDB(cfg)
	router := mux.NewRouter()

	//Init repositories
	playersRepo := playersRepository.NewPlayersRepository(db)
	matchesRepo := matchesRepository.NewMatchesRepository(db)
	tournamentsRepo := tournamentsRepository.NewTournamentsRepository(db)

	//Init useCases
	playersUC := playersUseCase.NewPlayersUseCase(playersRepo)
	matchesUC := matchesUseCase.NewMatchUseCase(matchesRepo, tournamentsRepo)
	templatesUC := templatesUseCase.NewTemplatesUseCase(playersRepo, matchesRepo, tournamentsRepo)
	tournamentsUC := tournamentsUseCase.NewTournamentsUseCase(tournamentsRepo, matchesRepo)

	//Init Handlers
	playersHandlers := playersHttp.NewPlayersHandlers(playersUC)
	matchesHandlers := matchesHttp.NewMatchesHandlers(matchesUC)
	templatesHandlers := templatesHttp.NewTemplatesHandlers(templatesUC)
	tournamentsHandlers := tournamentsHttp.NewTournamentsHandlers(tournamentsUC)

	playersHttp.MapPlayersRoutes(router, playersHandlers)
	matchesHttp.MapMatchesRoutes(router, matchesHandlers)
	templatesHttp.MapTemplatesRoutes(router, templatesHandlers)
	tournamentsHttp.MapTournamentsRoutes(router, tournamentsHandlers)

	router.PathPrefix("/swagger").HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		handler := httpSwagger.Handler(httpSwagger.URL(cfg.Server.Swagger))
		handler.ServeHTTP(w, r)
	}) // swagger config

	router.PathPrefix(docsPath).Handler(http.StripPrefix(docsPath, http.FileServer(http.Dir(docsDir)))) //swagger config statics path
	router.PathPrefix(cssPath).Handler(http.StripPrefix(cssPath, http.FileServer(http.Dir(cssDir))))
	router.PathPrefix(jsPath).Handler(http.StripPrefix(jsPath, http.FileServer(http.Dir(jsDir))))

	return router
}

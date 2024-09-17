package application

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	"github.com/Mishany52/testTaskSWoyo/handler"
	pathstore "github.com/Mishany52/testTaskSWoyo/pathStore"
)

func loadRoutes() *chi.Mux {
	router := chi.NewRouter()
	router.Use(middleware.Logger)

	//Думаю, что такой подход передачи в обработчик дурной тон
	pathStore := pathstore.NewPathStoreMap()
	shortPathHandler := handler.NewShortPath(pathStore)

	router.Get("/test", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})
	
	router.Post("/", shortPathHandler.Create)
	router.Get("/{urlKey}", shortPathHandler.GetByShortPath)

	return router
}


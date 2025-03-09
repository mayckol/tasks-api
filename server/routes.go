package server

import (
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	httpSwagger "github.com/swaggo/http-swagger/v2"
	_ "tasks-api/docs"
	"tasks-api/internal/infra/web"
)

type HandlersContainer struct {
	UserHandler web.UserHandler
}

func StartHttpHandler(hc *HandlersContainer, port int) *chi.Mux {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS", "PATCH"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type"},
		AllowCredentials: true,
		MaxAge:           300,
	}))

	r.Route("/api", func(r chi.Router) {
		r.Route("/v1", func(r chi.Router) {
			r.Get("/docs/*", httpSwagger.Handler(httpSwagger.URL(fmt.Sprintf("http://localhost:%d/api/v1/docs/doc.json", port))))
			// to avoid create two handlers like Auth and User, I chose to create only UserHandler to focus on the functional requirements of the task
			r.Route("/user", func(r chi.Router) {
				r.Post("/", hc.UserHandler.NewUser)
				r.Post("/signin", hc.UserHandler.Signin)
			})
		})
	})
	return r
}

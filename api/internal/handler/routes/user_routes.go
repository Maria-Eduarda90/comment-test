package routes

import (
	"api/internal/handler/userhandler"

	"github.com/go-chi/chi/v5"
)

func InitUserRoutes(router chi.Router, h userhandler.UserHandler) {
	router.Route("/user", func(r chi.Router){
		r.Post("/", h.CreateUser)
	})
}
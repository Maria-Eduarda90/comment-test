package routes

import (
	"api/internal/config/env"
	"api/internal/handler/userhandler"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/jwtauth/v5"
)

func InitUserRoutes(router chi.Router, h userhandler.UserHandler) {
	router.Post("/user", h.CreateUser)
	router.Route("/user", func(r chi.Router){
		r.Use(jwtauth.Verifier(env.Env.TokenAutn))
		r.Use(jwtauth.Authenticator(env.Env.TokenAutn))
		
		r.Patch("/{id}", h.UpdateUser)
		r.Get("/{id}", h.GetUserByID)
		r.Delete("/{id}", h.DeleteUser)
		r.Get("/", h.FindManyUsers)
		r.Patch("/password/{id}", h.UpdateUserPassword)
	})

	router.Route("/auth", func(r chi.Router){
		r.Post("/login", h.Login)
	})
}
package main

import (
	"api/internal/config/env"
	"api/internal/config/logger"
	"api/internal/database"
	"api/internal/database/sqlc"
	"api/internal/handler/routes"
	"api/internal/handler/userhandler"
	"api/internal/repository/userepository"
	"api/internal/service/userservices"
	"fmt"
	"log/slog"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func main() {
	logger.InitLogger()
	slog.Info("starting api")

	_, err := env.LoadingConfig(".")
	if err != nil {
		slog.Error("failed to load environment variables", "err", err, slog.String("package", "main"))

		return
	}

	dbConnection, err := database.NewDBConnection()
	if err != nil {
		slog.Error("error to connect to database", "err", err, slog.String("package", "main"))

		return
	}

	router := chi.NewRouter()
	queries := sqlc.New(dbConnection)

	userRepo := userepository.NewUserRepository(dbConnection, queries)
	newUserService := userservices.NewUserService(userRepo)
	newUserHandler := userhandler.NewUserHandler(newUserService)


	routes.InitUserRoutes(router, newUserHandler)

	port := fmt.Sprintf(":%s", env.Env.GoPort)
	slog.Info(fmt.Sprintf("server running on port %s", port))
	err = http.ListenAndServe(port, router)

	if err != nil {
		slog.Error("error to start server", "err", err, slog.String("package", "main"))
	}

	// testHash()
}

// func testHash(){
// 	password := "Meyh123456@"
// 	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
// 	if err != nil {
// 		fmt.Println("Error generating hash:", err)
// 		return
// 	}
// 	fmt.Println("Generated Hash:", string(hash))
// }
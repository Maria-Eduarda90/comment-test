package main

import (
	"api/internal/config/logger"
	"log/slog"
)

type user struct {
	Name	 string `json:"name"`
	Age		 int `json:"Age"`
	Password string `json:"password"`
}

func (u user) LogUser() slog.Value {
	return slog.GroupValue(
		slog.String("name", u.Name),
		slog.Int("age", u.Age),
		slog.String("password", "HIDDEN"),
	)
}

func main() {
	logger.InitLogger()

	user := user{
		Name:	"Maria",
		Age:	20,
		Password: "meyh123",
	}

	slog.Info("starting api")
	slog.Info("creating user", "user", user.LogUser())
}
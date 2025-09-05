package main

import (
	"log/slog"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/joaoguilherme2909/crudUsers/api"
	"github.com/joaoguilherme2909/crudUsers/store"
)

func main() {
	if err := run(); err != nil {
		slog.Error("Failed to execute code", "error", err)
		return
	}
	slog.Info("All systems offline")
}

func run() error {

	db := store.UserRepo{}

	id, _ := uuid.NewRandom()

	db[id] = store.User{
		FirstName: "Joao",
		LastName:  "Guilherme",
		Bio:       "A Full stack developer",
		Id:        id,
	}

	handler := api.NewHandler(db)

	s := http.Server{
		Addr:         ":8080",
		Handler:      handler,
		ReadTimeout:  10 * time.Second,
		IdleTimeout:  time.Minute,
		WriteTimeout: 10 * time.Second,
	}

	if err := s.ListenAndServe(); err != nil {
		return err
	}

	return nil
}

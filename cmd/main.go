package main

import (
	"context"
	"homework-3/config"
	"homework-3/internal/pkg/db"
	"homework-3/internal/pkg/repository/postgresql"
	"homework-3/internal/pkg/routers"
	"homework-3/internal/pkg/server"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

const port = ":9000"

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	config := config.NewConfig()
	database, err := db.NewDB(ctx, config)
	if err != nil {
		log.Fatal(err)
	}
	defer database.GetPool(ctx).Close()

	authorRepo := postgresql.NewAuthors(database)
	bookRepo := postgresql.NewBooks(database)

	s := server.NewServer(authorRepo, bookRepo)
	router := mux.NewRouter()
	routers.CreateAuthorRouter(router, s)
	routers.CreateBookSubRouter(router, *s)
	http.Handle("/", router)

	if err := http.ListenAndServe(port, nil); err != nil {
		log.Fatal(err)
	}
}

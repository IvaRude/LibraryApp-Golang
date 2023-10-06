package server

import "homework-3/internal/pkg/repository/postgresql"

type Server struct {
	AuthorRepo *postgresql.AuthorRepo
	BookRepo   *postgresql.BookRepo
}

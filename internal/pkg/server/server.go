package server

import (
	"homework-3/internal/pkg/repository"
)

type Server struct {
	AuthorRepo repository.AuthorsRepo
	BookRepo   repository.BooksRepo
}

func NewServer(authorRepo repository.AuthorsRepo, bookRepo repository.BooksRepo) *Server {
	return &Server{AuthorRepo: authorRepo, BookRepo: bookRepo}
}

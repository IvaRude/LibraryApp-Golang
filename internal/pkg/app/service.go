package app

type App struct {
	AuthorRepo AuthorsRepo
	BookRepo   BooksRepo
}

func NewApp(authorRepo AuthorsRepo, bookRepo BooksRepo) *App {
	return &App{AuthorRepo: authorRepo, BookRepo: bookRepo}
}

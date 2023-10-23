package app

type App struct {
	AuthorRepo      AuthorsRepo
	BookRepo        BooksRepo
	HandlerSender   Sender
	MessageReceiver Receiver
}

func NewApp(authorRepo AuthorsRepo, bookRepo BooksRepo, sender Sender) *App {
	return &App{AuthorRepo: authorRepo, BookRepo: bookRepo, HandlerSender: sender}
}

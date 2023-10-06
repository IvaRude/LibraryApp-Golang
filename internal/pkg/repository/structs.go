package repository

type Author struct {
	Id    int64  `db:"id"`
	Name  string `db:"name"`
	Books []Book
}

type Book struct {
	Id       int64  `db:"id"`
	Name     string `db:"name"`
	AuthorId int64  `db:"author_id"`
}

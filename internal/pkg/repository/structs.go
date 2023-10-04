package repository

type Author struct {
	Id int64 `db:"id"`
	Name string `db:"name"`
}
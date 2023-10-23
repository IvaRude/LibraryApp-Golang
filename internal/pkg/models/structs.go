package models

import (
	"time"
)

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

type Item struct {
	ID       int64  `json:"id"`
	Name     string `json:"name"`
	AuthorID int64  `json:"author_id"`
}

type Request struct {
	Method string
	Body   Item
}

type HandlerMessage struct {
	Timestamp time.Time
	EventType string
	Req       Request
}

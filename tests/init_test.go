package tests

import "homework-3/tests/postgres"

var (
	db *postgres.TDB
)

func init() {
	db = postgres.NewFromEnv()
}

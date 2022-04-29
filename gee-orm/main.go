package main

import (
	"database/sql"
	"gee-orm/session"
)

func main() {
	TestDB, _ := sql.Open("sqlite3", "../gee.db")
	session.New(TestDB)
}

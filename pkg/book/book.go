package book

import "database/sql"

type Book struct {
	DB *sql.DB
}

func CreateTable(db *sql.DB) *Book {
	statement, _ := db.Prepare(`
		CREATE TABLE IF NOT EXISTS "book" (
			"id"	INTEGER NOT NULL UNIQUE,
			"title"	TEXT NOT NULL,
			"isbn" TEXT NOT NULL,
			"author" TEXT NOT NULL,
			"publisher" TEXT NOT NULL,
			PRIMARY KEY("id" AUTOINCREMENT)
		);
	`)
	statement.Exec()
	return &Book{DB: db}
}

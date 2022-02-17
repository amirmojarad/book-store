package main

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

/*
	Book {
		id,
		title,
		isbn,
		author,
		publisher,
		date,
		pages,
		language
	}

	Order {
		customerID,
		BookID,
		Date,
		price
	}


	Favorites{
		costumerID,
		bookID
	}

	Customer {
		id,
		firstName,
		lastName,
		email,
		password
	}

*/

func main() {
	db, _ := sql.Open("sqlite3", "./db/bookstore.db")
	// statement, _ := db.Prepare(`
	// 	CREATE TABLE "book" (
	// 		"id"	INTEGER NOT NULL UNIQUE,
	// 		"title"	TEXT NOT NULL,
	// 		"isbn" TEXT NOT NULL,
	// 		"author" TEXT NOT NULL,
	// 		"publisher" TEXT NOT NULL,
	// 		PRIMARY KEY("id" AUTOINCREMENT)
	// 	);
	// `)
	// statement, _ := db.Prepare(`
	// 	CREATE TABLE "order" (
	// 		"id" INTEGER NOT NULL UNIQUE,
	// 		"book_id"	INTEGER NOT NULL,
	// 		"customer_id"	INTEGER NOT NULL,
	// 		"date" TEXT NOT NULL,
	// 		"price" TEXT NOT NULL,
	// 		PRIMARY KEY("id" AUTOINCREMENT)
	// 	);
	// `)
	// statement, _ := db.Prepare(`
	// 	CREATE TABLE "favorites" (
	// 		"id" INTEGER NOT NULL UNIQUE,
	// 		"book_id"	INTEGER NOT NULL ,
	// 		"customer_id"	INTEGER NOT NULL ,
	// 		PRIMARY KEY("id" AUTOINCREMENT)
	// 	);
	// `)
	statement, _ := db.Prepare(`
		CREATE TABLE "customer" (
			"id" 			INTEGER NOT NULL UNIQUE,
			"first_name"	TEXT NOT NULL,
			"last_name"		TEXT NOT NULL,
			"password" 		TEXT NOT NULL,
			"email"			TEXT NOT NULL,
			PRIMARY KEY("id" AUTOINCREMENT)
		);
	`)
	statement.Exec()

}

package book

import "database/sql"

type Book struct {
	DB *sql.DB
}

func (book *Book) DeleteItem(id int) {
	statement, _ := book.DB.Prepare(`
		delete from book where id = ?
	`)
	statement.Exec(id)
}

func (book *Book) GetItem(id int) *BookItem {
	allItems := book.GetAll()
	for _, item := range allItems {
		if item.ID == id {
			return &item
		}
	}
	return nil
}

func (book *Book) GetAll() []BookItem {
	rows, _ := book.DB.Query(`
		select * from book
	`)
	items := []BookItem{}
	newItem := BookItem{}
	for rows.Next() {
		rows.Scan(&newItem.ID, &newItem.Title, &newItem.ISBN, &newItem.Author, &newItem.Publisher)
		items = append(items, newItem)
	}
	return items
}

func (book *Book) AddItem(item BookItem) {
	statement, _ := book.DB.Prepare(`
		INSERT INTO book (
			title,
			isbn,
			author,
			publisher
		)
		VALUES (?, ?, ?, ?)
	`)
	statement.Exec(item.Title, item.ISBN, item.Author, item.Publisher)
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

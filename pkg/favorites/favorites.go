package favorites

import (
	"database/sql"
)

type Favorites struct {
	DB *sql.DB
}

func (favorites Favorites) GetAll() []FavoritesItem {
	rows, _ := favorites.DB.Query(`
		SELECT * from favorites
	`)
	items := []FavoritesItem{}
	newItem := FavoritesItem{}
	for rows.Next() {
		rows.Scan(&newItem.ID, &newItem.BookID, &newItem.CustomerID)
		items = append(items, newItem)
	}
	return items
}

func (favorites Favorites) AddItem(item FavoritesItem) {
	statement, _ := favorites.DB.Prepare(`
		INSERT INTO favorites (
			book_id,
			customer_id
		)
		VALUES (?, ?)
	`)
	statement.Exec(item.BookID, item.CustomerID)
}

func CreateTable(db *sql.DB) *Favorites {
	statement, _ := db.Prepare(`
		CREATE TABLE IF NOT EXISTS "favorites" (
			"id" INTEGER NOT NULL UNIQUE,
			"book_id"	INTEGER NOT NULL ,
			"customer_id"	INTEGER NOT NULL ,
			PRIMARY KEY("id" AUTOINCREMENT)
		);
	`)
	statement.Exec()
	return &Favorites{DB: db}
}

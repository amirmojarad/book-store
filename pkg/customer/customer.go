package customer

import (
	"database/sql"
	"strings"
)

type Customer struct {
	DB *sql.DB
}

// func (customer *Customer) DeleteItem(email string) *CustomerItem {
// 	allItems := customer.GetAll()

// }

func (customer *Customer) GetItem(email string) *CustomerItem {
	allItems := customer.GetAll()
	for _, item := range allItems {
		if strings.Compare(item.Email, email) == 0 {
			return &item
		}
	}
	return nil
}

func (customer *Customer) GetAll() []CustomerItem {
	items := []CustomerItem{}
	rows, _ := customer.DB.Query(`
		SELECT * FROM customer
	`)
	newItem := CustomerItem{}
	for rows.Next() {
		rows.Scan(&newItem.ID, &newItem.FirstName, &newItem.LastName, &newItem.Password, &newItem.Email)
		items = append(items, newItem)
	}
	return items
}

func (customer *Customer) Add(item CustomerItem) {
	statement, _ := customer.DB.Prepare(`
		INSERT INTO customer (
			first_name,
			last_name,
			password,
			email
		)
		VALUES (?, ?, ?, ?)
	`)
	statement.Exec(item.FirstName, item.LastName, item.Password, item.Email)
}

func CreateTable(db *sql.DB) *Customer {
	statement, _ := db.Prepare(`
		CREATE TABLE IF NOT EXISTS "customer" (
			"id" 			INTEGER NOT NULL UNIQUE,
			"first_name"	TEXT NOT NULL,
			"last_name"		TEXT NOT NULL,
			"password" 		TEXT NOT NULL,
			"email"			TEXT NOT NULL UNIQUE,
			PRIMARY KEY("id" AUTOINCREMENT)
		);	
	`)
	statement.Exec()
	return &Customer{
		DB: db,
	}
}

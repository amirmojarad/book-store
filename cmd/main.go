package main

import (
	"book-store/pkg/customer"
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, _ := sql.Open("sqlite3", "./db/bookstore.db")
	customerInstance := customer.CreateTable(db)
	fmt.Println(customerInstance.GetAll())
	findedCustomer := customerInstance.GetItem("amirmoajrad@gmail.com")
	fmt.Println(findedCustomer)
	customerInstance.DeleteItem("amirmoajrad@gmail.com")
	fmt.Println(customerInstance.GetAll())
}

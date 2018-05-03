package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

type Book struct {
	isbn   string
	title  string
	author string
	price  float32
}

var db *sql.DB

func main() {
	db, err := sql.Open("postgres", "user=charles password=nan dbname=db sslmode=disable")
	check(err)

	rows, err := db.Query("SELECT * FROM books")
	if err != nil {
		check(err)
	}
	defer rows.Close()

	books := []*Book{}
	for rows.Next() {
		bk := new(Book)
		err = rows.Scan(&bk.isbn, &bk.title, &bk.author, &bk.price)
		check(err)
		books = append(books, bk)
	}

	for _, bk := range books {
		fmt.Printf("%s, %s, %s, £%.2f\n", bk.isbn, bk.title, bk.author, bk.price)
	}
}

func createTable() {
	_, err := db.Query(`
		CREATE TABLE books (
		isbn    char(14) NOT NULL,
		title   varchar(255) NOT NULL,
		author  varchar(255) NOT NULL,
		price   decimal(5,2) NOT NULL )
	`)
	check(err)

	_, err = db.Exec(`
		INSERT INTO books (isbn, title, author, price) VALUES
		('978-1503261969', 'Emma', 'Jayne Austen', 9.44),
		('978-1505255607', 'The Time Machine', 'H. G. Wells', 5.99),
		('978-1503379640', 'The Prince', 'Niccolò Machiavelli', 6.99)
	`)
	check(err)

	_, err = db.Exec(`
		ALTER TABLE books ADD PRIMARY KEY(isbn)
	`)
	check(err)
}

func check(err error) {
	if err != nil {
		fmt.Println(err)
	}
}

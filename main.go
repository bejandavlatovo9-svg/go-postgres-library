package main

import (
    "database/sql"
    "fmt"
    "log"

    _ "github.com/lib/pq"
)

func showBooks(db *sql.DB) {
	rows, err := db.Query(`
		SELECT id, title, author, price, page
		FROM booksss
		ORDER BY id
	`)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	fmt.Println("\n--- The list of books ---")

	for rows.Next() {
		var id, price, page int
		var title, author string

		err := rows.Scan(
			&id,
			&title,
			&author,
			&price,
			&page,
		)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf(
			"ID: %d | Name: %s | Author: %s | Price: %d | Page: %d\n",
			id, title, author, price, page,
		)
	}

	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}
}

func main() {
	connStr := "host=localhost port=5433 user=postgres password=BEZH0909 dbname=postgres sslmode=disable"

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}

	fmt.Println("We have connected to PostgreSQL!")

	var newID int

	
	err = db.QueryRow(`
		INSERT INTO booksss (title, author, price, page)
		VALUES ($1, $2, $3, $4)
		RETURNING id
	`,
		"Go for beginners",
		"Bezhan Davlatov",
		250,
		300,
	).Scan(&newID)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("The book is added! ID:", newID)

	
	result, err := db.Exec(`
		UPDATE booksss
		SET price = $1
		WHERE id = $2
	`, 400, newID)

	if err != nil {
		log.Fatal(err)
	}

	count, err := result.RowsAffected()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("The book price was updated. Rows:", count)

	// SELECT: Display all books
	showBooks(db)

	// DELETE: Remove the new book
	result, err = db.Exec(`
		DELETE FROM booksss
		WHERE id = $1
	`, newID)

	if err != nil {
		log.Fatal(err)
	}

	count, err = result.RowsAffected()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("The book was removed. Rows:", count)

	// SELECT: Display the remaining books
	showBooks(db)
}
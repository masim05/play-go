package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

var query = `Select * FROM users ORDER BY id;`

func main() {
	db, err := sql.Open("postgres", "postgres://postgres:@localhost/coupon_service?sslmode=disable")
	if err != nil {
		panic(fmt.Sprintf("sql.Open error: %v", err))
	}

	defer db.Close()

	//age := 21
	rows, err := db.Query(query)

	defer rows.Close()

	if err != nil {
		panic(fmt.Sprintf("db.Query error: %v", err))
	}

	for rows.Next() {
		var (
			id   string
			name string
		)
		if err := rows.Scan(&id, &name); err != nil {
			panic(fmt.Sprintf("rows.Scan error: %v", err))
		}
		fmt.Printf("id: %s, name: %s\n", id, name)
	}
	if err := rows.Err(); err != nil {
		panic(fmt.Sprintf("rows.Err error: %v", err))
	}

	fmt.Println("Finish.\n")
}

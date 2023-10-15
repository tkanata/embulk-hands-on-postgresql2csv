package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

func main() {
	connStr := "host=localhost user=postgres password=password dbname=embulk-handson sslmode=disable"

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	teas := []struct {
		name  string
		color string
	}{
		{"Green Tea", "Green"},
		{"Black Tea", "Black"},
		{"White Tea", "White"},
		{"Oolong Tea", "Brown"},
		{"Herbal Tea", "Red"},
	}

	for _, tea := range teas {
		_, err := db.Exec("INSERT INTO teas (name, color) VALUES ($1, $2)", tea.name, tea.color)
		if err != nil {
			log.Fatal(err)
		}
	}

	fmt.Println("5件のデータを登録しました!")
}

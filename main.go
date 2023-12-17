package main

import (
	"database/sql"
	"encoding/csv"
	"fmt"
	"os"

	_ "github.com/lib/pq"
)

func main() {
	// データベースに接続
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable TimeZone=%s",
		os.Getenv("PSQL_ENDPOINT"), os.Getenv("PSQL_PORT"), os.Getenv("PSQL_USER"), os.Getenv("PSQL_PASSWORD"), os.Getenv("PSQL_DBNAME"), os.Getenv("PSQL_TZ"))
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	// テーブルから全てのレコードを削除
	fmt.Println("Deleting all existing records from the 'teas' table...")
	_, err = db.Exec("DELETE FROM teas")
	if err != nil {
		panic(err)  // エラーがあればここでパニックを起こす
	}
	fmt.Println("All existing records have been deleted.")


  // 既存のIDを取得してセットに保存
	existingIDs := make(map[string]bool)
	rows, err := db.Query("SELECT id FROM teas")
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	for rows.Next() {
		var id string
		if err := rows.Scan(&id); err != nil {
			panic(err)
		}
		existingIDs[id] = true
	}
	if err := rows.Err(); err != nil {
		panic(err)
	}

	// CSV ファイルを開く
	f, err := os.Open("teas.csv")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	// CSV 内容を読む
	r := csv.NewReader(f)
	records, err := r.ReadAll()
	if err != nil {
		panic(err)
	}

  // ヘッダー行をスキップするため、1から始める
	for _, record := range records[1:] {
		id := record[0]
		name := record[1]
		color := record[2]

		_, err := db.Exec("INSERT INTO teas (id, name, color) VALUES ($1, $2, $3)", id, name, color)
		if err != nil {
			panic(err)
		}
	}

	
	fmt.Println("Data imported successfully!")
}


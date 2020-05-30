package main

import (
	"fmt"
	"github.com/jmoiron/sqlx"
)

func main() {
	dns := "root:password@tcp(127.0.0.1:3306)/game_db"

	db, err := sqlx.Connect("mysql", dns)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	count := 0
	err = db.Get(&count, "SELECT COUNT(*) FROM account_data")
	if err != nil {
		panic(err)
	}
	fmt.Println(count)
}

package main

import "github.com/henrylee2cn/faygo/ext/db/sqlx"

func createTable() error {
	return sqlx.CallbackByName("activation_code", func(d sqlx.DbOrTx) error {
		_, err := d.Exec("CREATE TABLE code_category_data_1 LIKE code_category_data")
		if err != nil {
			return err
		}
		return nil
	})
}

func main() {
	err := createTable()
	if err != nil {
		panic(err)
	}
}

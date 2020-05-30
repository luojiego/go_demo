package main

import (
	//"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"math/rand"
	"strconv"
)

//往表 test1 中写入 1000 条数据

var (
	dns = "root:password@tcp(192.168.196.50:3306)/test"
	db  *sqlx.DB
)

func init() {
	var err error
	db, err = sqlx.Connect("mysql", dns)
	if err != nil {
		panic(err)
	}
	//defer db.Close()
}

func main() {
	for i := 1000; i < 11000; i++ {
		db.Exec("INSERT INTO test1(`user_id`,`name`,`level`,`vip_level`,`trophy`) VALUES (?,?,?,?,?)",
			i, "user"+strconv.FormatInt(int64(i), 10), rand.Intn(128), rand.Intn(128), rand.Intn(65535))
	}
}

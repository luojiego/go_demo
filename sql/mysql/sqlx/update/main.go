package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"sync"
	"time"
)

//并发的更新一个字段，和 mongo 对比性能

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
	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(100)
}

func query() {
	age := 0
	//err := db.Get(&age, "SELECT `trophy` FROM `test1` WHERE `name` = ?", "user1980")
	//err := db.Get(&age, "SELECT `trophy` FROM `test1` WHERE `name` = ?", "user10025")
	err := db.Get(&age, "SELECT `trophy` FROM `test1` WHERE `user_id` = ?", 10025)
	if err != nil {
		panic(err)
	}
	fmt.Println(age)
}

func incr(id int, wg *sync.WaitGroup) {
	defer wg.Done()
	/*_, err := db.Exec("UPDATE `test1` SET `trophy` = `trophy` + 1 WHERE `name` = ?",
	"user" + strconv.FormatInt(int64(id), 10))*/
	_, err := db.Exec("UPDATE `test1` SET `trophy` = `trophy` + 1 WHERE `user_id` = ?", id)
	//_, err := db.Exec("UPDATE `test1` SET `trophy` = `trophy` + 1 WHERE `name` = ?", "user1980")
	if err != nil {
		panic(err)
	}

}

func main() {
	query()
	now := time.Now()
	wg := sync.WaitGroup{}
	count := 1000
	for i := 0; i < count; i++ {
		wg.Add(1)
		go incr(10000+i, &wg)
	}
	wg.Wait()
	fmt.Printf("update count:%d use %0.5fs\n", count, time.Since(now).Seconds())
	query()
}

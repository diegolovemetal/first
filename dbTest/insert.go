package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type Person struct {
	UserId int `db:"user_id"`
	UserName string `db:"username"`
	Sex string `db:"sex"`
	Email string `db:"email"`
}

type Place struct {
	Country string `db:"country"`
	City string `db:"city"`
	TelCode string `db:"telcode"`
}

var DB *sqlx.DB
func init() {
	database, err := sqlx.Open("mysql", "root:root@tcp(127.0.0.1:3306)/test")
	if err != nil {
		fmt.Println("exec failed, err:", err)
		return
	}
	DB = database
}

func main() {
	r, err := DB.Exec("insert into person(username, sex, email)values (?, ?, ?)","stu001", "male", "test@test.com")
	if err != nil {
		fmt.Println("exec err :", err)
		return
	}
	id, err := r.LastInsertId()
	if err != nil {
		fmt.Println("exec err:", err)
		return
	}

	fmt.Println("insert success:", id)
}
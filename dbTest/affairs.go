package main

import (
"fmt"

_ "github.com/go-sql-driver/mysql"
"github.com/jmoiron/sqlx"
)

type Person struct {
UserId   int    `db:"user_id"`
Username string `db:"username"`
Sex      string `db:"sex"`
Email    string `db:"email"`
}

type Place struct {
Country string `db:"country"`
City    string `db:"city"`
TelCode int    `db:"telcode"`
}

var Db *sqlx.DB

func init() {
database, err := sqlx.Open("mysql", "root:root@tcp(127.0.0.1:3306)/test")
if err != nil {
fmt.Println("open mysql failed,", err)
return
}
Db = database
}

func main()  {
	conn, err := Db.Begin()
	if err != nil {
		fmt.Println("begin failed,err:", err)
		return
	}

	r, err := conn.Exec("insert into person(username, sex, email) values (?,?,?)","stu001","male","123@qq.com")
	if err != nil {
		fmt.Println("exec failed, err :",err)
		conn.Rollback()
		return
	}
	id, err := r.LastInsertId()
	if err != nil {
		fmt.Println("exec failed,err:", err)
		conn.Rollback()
		return
	}
	fmt.Println("insert success,id:", id)

	r, err = conn.Exec("insert into person(username, sex, email) values (?,?,?)","stu001","male","123@qq.com")
	if err != nil {
		fmt.Println("exec failed, err :",err)
		conn.Rollback()
		return
	}
	id, err = r.LastInsertId()
	if err != nil {
		fmt.Println("exec failed,err:", err)
		conn.Rollback()
		return
	}
	fmt.Println("insert success,id:", id)

	conn.Commit()

}
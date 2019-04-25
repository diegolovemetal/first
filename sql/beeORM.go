package main

import (
	"fmt"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)
// Model Struct
type User struct {
	Id int
	Name string `orm:"size(100)"`
}
func init() {
	// 设置默认数据库
	orm.RegisterDataBase("default", "mysql", "root:root@tcp(127.0.0.1:3306)/db_test?charset=utf8", 30)

	//注册定义的Model
	orm.RegisterModel(new(User))
	//RegisterModel 也可以同时注册多个 model
	//orm.RegisterModel(new(User), new(Profile), new(Post))

	//创建table
	orm.RunSyncdb("default", false, true)
}

func main() {
	o := orm.NewOrm()

	user := User{Name: "diego"}

	//插入表
	id, err := o.Insert(&user)
	fmt.Printf("ID: %d, Err: %v\n",id ,err)

	//更新表
	user.Name ="diegoUpdate"
	num, err := o.Update(&user)
	fmt.Printf("num: %d, Err: %v\n", num, err)

	//读取
	u := User{Id: user.Id}
	err = o.Read(&u)
	fmt.Printf("ERR: %v\n", err)
	// 删除表
	num, err = o.Delete(&u)
	fmt.Printf("num: %d, ERR: %v\n", num, err)
}

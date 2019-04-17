package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"time"
)
var db *gorm.DB

func init() {
	var err error
	db, err = gorm.Open("mysql", "root:root@tcp(127.0.0.1:3306)/test")	//<user>:<password>/<database>?charset=utf8&parseTime=True&loc=Local
	if err != nil {
		panic(err)
	}
	fmt.Println("connect mysql success")
}

type Like struct {
	ID        int    `gorm:"primary_key"`
	Ip        string `gorm:"type:varchar(20);not null;index:ip_idx"`
	Ua        string `gorm:"type:varchar(256);not null;"`
	Title     string `gorm:"type:varchar(128);not null;index:title_idx"`
	Hash      uint64 `gorm:"unique_index:hash_idx;"`
	CreatedAt time.Time
}


func main()  {
	if !db.HasTable(&Like{}) {
		if err := db.Set("gorm:table_options", "ENGINE=InnoDB DEFAULT CHARSET=utf8").CreateTable(&Like{}).Error; err != nil {
			panic(err)
		}
	}
	like := &Like{
	   Ip:        "1234567",
	   Ua:        "diego",
	   Title:     "I love U",
	   CreatedAt: time.Now(),
	}

	if err := db.Create(like).Error; err != nil {
	   return
	}

	if err := db.Where(&Like{Ua: "diego"}).Delete(Like{}).Error; err != nil {
		return
	}
	var count int
	err := db.Model(&Like{}).Where(&Like{Ua: "diego"}).Count(&count).Error
	if err != nil {
		return
	}
	//db.Model(&user).Update("name", "hello")
	//db.Model(&user).Updates(User{Name: "hello", Age: 18})
	//db.Model(&user).Updates(User{Name: "", Age: 0, Actived: false}) // nothing update
}

//事务
//func CreateAnimals(db *gorm.DB) err {
//    tx := db.Begin()
//    if err := tx.Create(&Animal{Name: "Giraffe"}).Error; err != nil {
//        tx.Rollback()
//        return err
//    }
//    if err := tx.Create(&Animal{Name: "Lion"}).Error; err != nil {
//        tx.Rollback()
//        return err
//    }
//    tx.Commit()
//    return nil
//}

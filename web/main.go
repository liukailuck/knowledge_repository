package main

import (
	"../dbhelp"
	"fmt"
	"time"
)

type User struct {
	Id      int64
	Name    string
	Age     int32
	DelFlag bool
	Created time.Time `xorm:"created"`
}

func main() {

	db := dbhelp.NewSingleDbEngine()

	users := []User{}
	if err := db.Find(&users); err != nil {
		//TODO   Log
		fmt.Println(err)
		return
	}
	fmt.Println(users)
}

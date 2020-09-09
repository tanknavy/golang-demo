package main

import (
	"fmt"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

type User struct{
	Name string `json:"name"`
}

func main(){
	fmt.Println("Go MySQL Tutorial")

	//open mysql connection
	db,err := sql.Open("mysql","hive:q1w2e3r4@tcp(spark3:3306)/test")

	if err != nil{
		panic(err.Error())
	}
	fmt.Println("successfully connected to mysql")
	
	defer db.Close() //最后关闭

	//mysql insert
	insert, err := db.Query("INSERT INTO tmp(name) VALUES('Carl Cheng')")
	if err != nil{
		panic(err.Error())
	}
	fmt.Println("successfully inserted into mysql tmp table")

	defer insert.Close() //插入并关闭

	//mysql select and print results
	results,err := db.Query("SELECT name from tmp")
	if err != nil{
		panic(err.Error())
	}
	for results.Next(){ //一直next
		var user User
		
		err = results.Scan(&user.Name) //拿到结果放到user变量地址
		if err != nil{
			panic(err.Error())
		}
		fmt.Println(user.Name)
	}


	
	
}
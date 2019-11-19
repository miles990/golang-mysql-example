package main

import(	
	db "github.com/golang-mysql-exampledb/mysql"	
)

import (
	"fmt"
	// "encoding/json"
)

func main()  {	

	db.ResetTable()	
	fmt.Println("insert ...")
	account1, _ := db.CreateAccount("Apple")
	fmt.Println(account1)
	account2, _ := db.CreateAccount("Banana")
	fmt.Println(account2)
	fmt.Println("")

	fmt.Println("query ...")
	accounts, _ := db.GetAccounts()
	for _, v := range accounts {
		fmt.Println(v)
	}	
	fmt.Println("")
}
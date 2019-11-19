package db

import (	
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"	
)

import (
	"fmt"
)

var dsn string
var connector *sqlx.DB

func init() {
	var err error
	dsn = fmt.Sprintf("%s:%s@tcp(%s)/%s", "dev", "dev", "127.0.0.1:3306", "gamedata")
	connector, err = sqlx.Open("mysql", dsn)
	if err != nil {
		fmt.Println("[DB] mysql connect error:", err.Error())
		return
	}
	err = connector.Ping()
	//client.Set()
	if err != nil {
		fmt.Println("[DB] mysql connect error:", err.Error())
		panic(err.Error())
	}
	fmt.Println("[DB] mysql init success... ")
}

func createTable() (err error) {
	var schema = `CREATE TABLE accounts (
	    uid text,
	    name text,
	    createtimestamp bigint);`
	 
	// execute a query on the server
	_, err = connector.Exec(schema)
	if err != nil {
		fmt.Println("[DB] mysql create table[accounts] error:", err.Error())
		return err
	}
	// fmt.Println(result)
	return nil
}

func dropTable() (err error) {
	var schema = `DROP TABLE accounts;`
	 
	// execute a query on the server
	_, err = connector.Exec(schema)
	if err != nil {
		fmt.Println("[DB] mysql drop table[accounts] error:", err.Error())
		return err
	}
	// fmt.Println(result)
	return nil
}

func ensureConnection() (err error) {
	err = connector.Ping()
	if err != nil {
		fmt.Println("[DB] mysql ping failed, try reconnect. ", err.Error())
		connector, err = sqlx.Open("mysql", dsn)
		if err != nil {
			fmt.Println("[DB] mysql reconnect failed. ", err.Error())
			return err
		}
	}
	return nil
}
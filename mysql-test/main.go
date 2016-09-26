package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"gopkg.in/ini.v1"
)

func main() {

	cfg, err := ini.Load("mysql.conf")
	if err != nil {
		fmt.Println("load ini fail", err)
	}

	host := cfg.Section("DEFAULT").Key("host").String()
	port := cfg.Section("DEFAULT").Key("port").String()
	user := cfg.Section("DEFAULT").Key("user").String()
	password := cfg.Section("DEFAULT").Key("password").String()
	dbname := cfg.Section("DEFAULT").Key("dbname").String()

	dsn := user + ":" + password + "@tcp(" + host + ":" + port + ")/" + dbname
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		panic(err.Error()) // Just for example purpose. You should use proper error handling instead of panic
	} else {
		fmt.Println("open OK")
	}
	defer db.Close()

	// Open doesn't open a connection. Validate DSN data:
	err = db.Ping()
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	} else {
		fmt.Println("ping OK")
	}

	stmtOut, err := db.Prepare("SELECT id FROM t_user")
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}
	defer stmtOut.Close()

	var x int // we "scan" the result in here

	// Query the square-number of 13
	err = stmtOut.QueryRow().Scan(&x)
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}
	fmt.Printf("SELECT 1 FROM DUAL: %d\n", x)

}

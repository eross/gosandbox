package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/go-sql-driver/mysql"
)

var db *sql.DB

func mainping() {
	cfg := mysql.NewConfig()
	// Capture connection properties.
	//db, err := sql.Open("mysql", "user:password@/dbname")
	//cfg := mysql.NewConfig(){
	//	User:   os.Getenv("DBUSER"),
	//	Passwd: os.Getenv("DBPASS"),
	//	Net:    "tcp",
	//	Addr:   "0.0.0.0:3306",
	//	DBName: "recordings",
	//}
	// Get a database handle.
	cfg.User = os.Getenv("DBUSER")
	cfg.Passwd = os.Getenv("DBPASS")
	cfg.Net = "tcp"
	cfg.Addr = "0.0.0.0:3306"
	cfg.DBName = "recordings"
	var err error
	db, err = sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		log.Fatal(err)
	}

	pingErr := db.Ping()
	if pingErr != nil {
		log.Fatal(pingErr)
	}
	fmt.Println("Connected!")
}

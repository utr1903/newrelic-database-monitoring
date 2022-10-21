package main

import (
	"database/sql"
	"fmt"
	"os"
	"strconv"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

var DB_DRIVER string = "mysql"
var NUM_TABLES int = 5

func main() {

	// Config
	ipAddress := os.Getenv("MYSQL_IP")
	username := "myuser"
	password := os.Getenv("PASSWORD")
	dbName := "mydb"

	db := createDatabase(username, password, ipAddress, dbName)
	defer db.Close()

	for i := 1; i < NUM_TABLES; i++ {
		tableName := "mytable" + strconv.Itoa(i)
		createTable(db, tableName)
	}

	for i := 1; i < NUM_TABLES; i++ {
		tableName := "mytable" + strconv.Itoa(i)

		if i == 0 { // For table 1

			go func() {
				for {
					// Insert 2 value every loop
					for i := 1; i < 2; i++ {
						insert(db, tableName)
					}

					// List 5 times every loop
					for i := 1; i < 5; i++ {
						list(db, tableName)
					}
				}
			}()
		} else if i == 1 { // For table 2

			go func() {
				for {
					// Insert 3 value every loop
					for i := 1; i < 3; i++ {
						insert(db, tableName)
					}

					// List 4 times every loop
					for i := 1; i < 4; i++ {
						list(db, tableName)
					}
				}
			}()
		} else { // For the rest of the tables

			go func() {
				for {
					// Insert 1 value every loop
					insert(db, tableName)

					// List 2 times every loop
					for i := 1; i < 2; i++ {
						list(db, tableName)
					}
				}
			}()
		}
	}

	for {
		sleep()
	}
}

func createDatabase(
	username string,
	password string,
	ipAddress string,
	dbName string,
) *sql.DB {

	conn := username + ":" + password + "@tcp(" + ipAddress + ":3306)/"
	fmt.Println("Connecting to MySQL [" + conn + "]...")
	db, err := sql.Open(DB_DRIVER, conn)
	if err != nil {
		panic(err)
	}
	fmt.Println(" -> Connected to MySQL.")

	fmt.Println("Creating DB...")
	_, err = db.Exec("CREATE DATABASE IF NOT EXISTS " + dbName)
	if err != nil {
		panic(err)
	}
	fmt.Println(" -> Created the DB.")

	fmt.Println("Using DB...")
	db, err = sql.Open(DB_DRIVER, conn+dbName)
	if err != nil {
		panic(err)
	}
	fmt.Println(" -> DB on use.")

	return db
}

func createTable(
	db *sql.DB,
	tableName string,
) {
	fmt.Println("Creating table [" + tableName + "]...")
	_, err := db.Exec("CREATE TABLE IF NOT EXISTS " + tableName + "( id INT UNSIGNED AUTO_INCREMENT PRIMARY KEY, data VARCHAR(255) )")
	if err != nil {
		panic(err)
	}
	fmt.Println(" -> Table [" + tableName + "] is created.")
}

func insert(
	db *sql.DB,
	tableName string,
) {
	fmt.Println("Inserting value to table [" + tableName + "]...")
	insert, err := db.Query("INSERT INTO " + tableName + "(data) VALUES ( 'TEST' )")
	if err != nil {
		panic(err.Error())
	}
	defer insert.Close()
	fmt.Println(" -> Value is inserted to table [" + tableName + "].")

	sleep()
}

func list(
	db *sql.DB,
	tableName string,
) {
	fmt.Println("Listing values from table [" + tableName + "]...")
	insert, err := db.Query("SELECT * FROM " + tableName + " LIMIT 5")
	if err != nil {
		panic(err.Error())
	}
	defer insert.Close()
	fmt.Println(" -> Values are listed from table [" + tableName + "].")

	sleep()
}

func sleep() {
	time.Sleep(2 * time.Second)
}

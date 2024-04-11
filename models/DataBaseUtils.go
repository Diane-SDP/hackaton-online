package models

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func OpenDb() {
	DB, _ = sql.Open("mysql", "Diane:J5T_/pg/G##u9~g@tcp(localhost:3306)/Hackaton")
}
func CreateDB() {
	_, err := DB.Exec(`
	CREATE TABLE IF NOT EXISTS Colis(
		Id INTEGER PRIMARY KEY AUTO_INCREMENT,
		Uid TEXT,
		IdShop INTEGER,
		StartAdress TEXT,
		FinalAdress TEXT,
		StartCity TEXT,
		ClientCity TEXT,
		Step INTEGER
	)
	`)
	if err != nil {
		panic(err)
	}
	_, err = DB.Exec(`
	CREATE TABLE IF NOT EXISTS Users(
		Id INTEGER PRIMARY KEY AUTO_INCREMENT,
		Passwd TEXT,
		Uid TEXT,
		Name TEXT,
		Admin INTEGER
	)
	`)
	if err != nil {
		panic(err)
	}
}

package db

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

const InitQuery = `CREATE TABLE IF NOT EXISTS users (
	userId INTEGER PRIMARY KEY AUTOINCREMENT,
	username TEXT NOT NULL,
	email TEXT NOT NULL,
	password TEXT NOT NULL );`



func Create() {
	database, err := sql.Open("sqlite3", "../../internal/database/db.db")
	if err != nil {
		log.Fatal(err.Error())
	}
	defer database.Close()

	_, err = database.Exec(InitQuery)
	if err != nil {
		log.Fatal(err, "Problem during UserDatabase creation")
	}
}

func VerifyUsername(username string, db *sql.DB) bool {
	var userId int
	err := db.QueryRow("SELECT userId FROM users WHERE username = ?", username).Scan(&userId)
	if err == nil {
		return true
	}
	return false
}

func VerifyMail(email string, db *sql.DB) bool {
	var userId int
	err := db.QueryRow("SELECT userId FROM users WHERE email = ?", email).Scan(&userId)
	if err == nil {
		return true
	}
	return false
}

func VerifyAccount(username string, password string, db *sql.DB) bool {
	var userId int
	err := db.QueryRow("SELECT userId FROM users WHERE username = ? AND password = ?", username, password).Scan(&userId)
	if err == nil {
		return false
	}
	return true
}

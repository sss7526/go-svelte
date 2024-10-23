package storage

import (
	"database/sql"
	"log"
	"errors"
	_ "github.com/mattn/go-sqlite3"
)

var db *sql.DB

func InitSQLiteDB() (*sql.DB, error) {
	database, err := sql.Open("sqlite3", "./user_database.db")
	if err != nil {
		return nil, err
	}

	createTableQuery := `
		CREATE TABLE IF NOT EXISTS user(
		username TEXT PRIMARY KEY,
		password TEXT NOT NULL,
		role TEXT NOT NULL
	)`

	statement, err := database.Prepare(createTableQuery)
	if err != nil{
		return nil, err
	}
	statement.Exec()

	db = database
	log.Println("SQLite database initialized.")
	return database, nil
}

func InsertUser(username, hashedPassword, role string) error {
	query := `
		INSERT INTO user (username, password, role)
		VALUES (?, ?, ?)	
	`
	
	stmt, err := db.Prepare(query)
	if err != nil {
		return err
	}
	_, err = stmt.Exec(username, hashedPassword, role)
	if err != nil {
		return err
	}
	return nil
}

func FindUserByUsername(username string) (User, error) {
	var user User
	query := `
		SELECT username, password, role
		FROM user
		WHERE username = ?
	`

	row := db.QueryRow(query, username)

	err := row.Scan(&user.Username, &user.Password, &user.Role)
	if err != nil {
		return user, errors.New("User not found")
	}
	return user, nil
}
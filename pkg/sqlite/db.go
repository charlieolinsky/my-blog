package sqlite

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

func InitDataBase(dbPath string) (*sql.DB, error) {
	//Validate Open Arguments
	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		return nil, err
	}

	//Ensure secure db connection using dummy query 
	err = db.Ping()
	if err != nil {
		return nil, err
	} 

	//Enable Foreign Key Support
	_, err = db.Exec("PRAGMA foreign_keys = ON")
	if err != nil {
		return nil, err
	}

	//Create Necessary Tables
	err = createTables(db)
	if err != nil {
		return nil, err
	}

	return db, nil
}

func createTables(db *sql.DB) error {

	//Create users table
	_, err := db.Exec(
	`CREATE TABLE users (
	user_id INTEGER PRIMARY KEY,
  	role TEXT NOT NULL,
  	email TEXT UNIQUE NOT NULL,
  	password TEXT NOT NULL,
  	first_name TEXT NOT NULL,
  	last_name TEXT NOT NULL,
  	profilePictureUrl TEXT,
  	created_at TEXT NOT NULL,
  	updated_at TEXT);`)
	
	if err != nil {
		return err
	}

	return nil
}

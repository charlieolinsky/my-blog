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
	`CREATE TABLE IF NOT EXISTS users (
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

	//Create projects table
	_, err = db.Exec(
	`CREATE TABLE IF NOT EXISTS projects (
  	project_id INTEGER PRIMARY KEY,
  	user_id INTEGER NOT NULL,
  	title TEXT NOT NULL,
  	body TEXT NOT NULL,
  	created_at TEXT NOT NULL,
  	updated_at TEXT,
  	FOREIGN KEY (user_id) REFERENCES users (user_id));`)
	if err != nil {
		return err
	}

	//Create posts table
	_, err = db.Exec(`
	CREATE TABLE IF NOT EXISTS posts (
  	post_id INTEGER PRIMARY KEY,
  	user_id INTEGER NOT NULL,
  	project_id INTEGER NOT NULL,
  	title TEXT NOT NULL,
  	body TEXT NOT NULL,
  	likes INTEGER DEFAULT 0,
  	created_at TEXT NOT NULL,
  	updated_at TEXT,
  	FOREIGN KEY (user_id) REFERENCES users (user_id),
  	FOREIGN KEY (project_id) REFERENCES projects (project_id));`)
	if err != nil {
		return err
	}

	//Create comments table
	_, err = db.Exec(`
	CREATE TABLE IF NOT EXISTS comments (
  	comment_id INTEGER PRIMARY KEY,
  	user_id INTEGER NOT NULL,
  	post_id INTEGER NOT NULL,
  	body TEXT NOT NULL,
  	likes INTEGER DEFAULT 0,
  	created_at TEXT NOT NULL,
  	updated_at TEXT,
  	FOREIGN KEY (user_id) REFERENCES users (user_id),
  	FOREIGN KEY (post_id) REFERENCES posts (post_id));`)
	if err != nil {
		return err
	}


	return nil
}

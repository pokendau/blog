package api

import (
	"database/sql"
	"errors"

	_ "github.com/mattn/go-sqlite3"
)

func DbSetup() (*sql.DB, error) {
	db, err := sql.Open("sqlite3", "../blog.db")

	if err != nil {
		return nil, errors.New("Error opening the database: " + err.Error())
	}

	_, err = db.Exec(
		`
    DROP TABLE IF EXISTS categories;
    CREATE TABLE categories(categoryid VARCHAR(36) PRIMARY KEY, name VARCHAR(20) UNIQUE NOT NULL);
    `,
	)

	if err != nil {
		return nil, errors.New("Error with categories: " + err.Error())
	}

	_, err = db.Exec(
		`
    DROP TABLE IF EXISTS articles;
    CREATE TABLE articles(UUID VARCHAR(36) PRIMARY KEY, title VARCHAR(50) NOT NULL, content TEXT, createdOn DATE,
    updatedOn TIMESTAMP DEFAULT CURRENT_TIMESTAMP, categoryid VARCHAR(36),
    FOREIGN KEY(categoryid) REFERENCES categories(categoryid)
);

    `,
	)

	if err != nil {
		return nil, errors.New("Error with articles: " + err.Error())
	}

	return db, nil
}

package services

import (
	"database/sql"
	"errors"

	"blog/api/models"
)

type Service interface {
	insert(db *sql.DB, model models.Model) error
	getOne(db *sql.DB, id string)
	getMore(db *sql.DB, filter ...string) ([]models.Model, error)
}

type ArticleService struct{}

type CategoryService struct{}

func (serv ArticleService) insert(db *sql.DB, article models.Article) error {
	_, err := db.Exec(
		"INSERT INTO articles (UUID, title, content, createdOn, updatedOn, category) VALUES (?, ?, ?, ?, ? ?);",
		article.Uuid,
		article.Title,
		article.Content,
		article.CreatedOn,
		article.UdpatedOn,
		article.Category.Uuid,
	)
	if err != nil {
		return errors.New("Error in article service: inserting - " + err.Error())
	}

	return nil
}

func (serv ArticleService) getMore(db *sql.DB, filter ...string) ([]models.Article, error) {
	var s []models.Article

	rows, err := db.Query("SELECT (UUID, title, content) FROM articles;")
	defer rows.Close()

	if err != nil {
		return nil, errors.New("Error in article service: retrieving - " + err.Error())
	}

	for rows.Next() {
		var art models.Article
		if err := rows.Scan(&art.Uuid, &art.Title, &art.Content); err != nil {
			return s, errors.New("Error in article service: reading rows - " + err.Error())
		}
		s = append(s, art)
	}

	if err = rows.Err(); err != nil {
		return s, errors.New("Error in article service: " + err.Error())
	}

	return s, nil
}

func (serv CategoryService) insert(db *sql.DB, category models.Category) error {
	_, err := db.Exec(
		"INSERT INTO category (categoryid, name) VALUES (?, ?);",
		category.Uuid,
		category.Name,
	)
	if err != nil {
		return errors.New("Error in category service: inserting - " + err.Error())
	}

	return nil
}

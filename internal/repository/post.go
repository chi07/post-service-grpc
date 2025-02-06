package repository

import (
	"errors"

	"post-service-grpc/internal/models"
	"post-service-grpc/pkg/db"
)

// CreateArticle inserts a new article into the database
func CreateArticle(title, content string) (uint, error) {
	article := models.Article{Title: title, Content: content}
	result := db.DB.Create(&article)
	return article.ID, result.Error
}

// EditArticle updates an existing article
func EditArticle(id uint, title, content string) error {
	var article models.Article
	result := db.DB.First(&article, id)
	if result.Error != nil {
		return errors.New("article not found")
	}

	article.Title = title
	article.Content = content
	return db.DB.Save(&article).Error
}

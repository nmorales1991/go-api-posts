package services

import (
	"context"
	"go-api-posts/pkg/models"
	"go-api-posts/pkg/repository"
	"go.mongodb.org/mongo-driver/bson"
)

type PostService struct {
	db *repository.DB
}

func NewPostService(db *repository.DB) *PostService {
	return &PostService{db: db}
}

func (s *PostService) GetPosts() ([]models.Post, error) {
	collection := s.db.GetCollection("api_posts", "posts")
	cursor, err := collection.Find(context.Background(), bson.D{})
	if err != nil {
		return nil, err
	}
	var posts []models.Post
	if err := cursor.All(context.Background(), &posts); err != nil {
		return nil, err
	}
	return posts, nil
}

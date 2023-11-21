package services

import (
	"context"
	"go-api-posts/pkg/models"
	"go-api-posts/pkg/repository"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type PostService struct {
	collection *mongo.Collection
}

func NewPostService(db *repository.DB) *PostService {
	return &PostService{collection: db.GetCollection("api_posts", "posts")}
}

func (s *PostService) GetPosts() ([]models.Post, error) {
	cursor, err := s.collection.Find(context.Background(), bson.D{})
	if err != nil {
		return nil, err
	}
	var posts []models.Post
	if err := cursor.All(context.Background(), &posts); err != nil {
		return nil, err
	}
	return posts, nil
}

func (s *PostService) CreatePost(post models.Post) error {
	_, err := s.collection.InsertOne(context.Background(), post)
	if err != nil {
		return err
	}
	return nil
}

func (s *PostService) GetPostById(id string) (models.Post, error) {
	var post models.Post
	objectID, _ := primitive.ObjectIDFromHex(id)
	err := s.collection.FindOne(context.Background(), bson.M{"_id": objectID}).Decode(&post)
	if err != nil {
		return post, err
	}
	return post, nil
}

func (s *PostService) UpdatePost(id string, post models.Post) (models.Post, error) {
	var newPost models.Post
	objectID, _ := primitive.ObjectIDFromHex(id)
	filter := bson.D{{Key: "_id", Value: objectID}}
	value := bson.D{
		{Key: "$set", Value: bson.D{
			{Key: "title", Value: post.Title},
			{Key: "description", Value: post.Description},
		}},
	}
	_, err := s.collection.UpdateOne(context.Background(), filter, value)

	if err != nil {
		return newPost, err
	}
	newPost, _ = s.GetPostById(id)
	return newPost, nil
}

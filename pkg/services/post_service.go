package services

import (
	"context"
	"errors"
	"fmt"
	"github.com/go-playground/validator/v10"
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
	if len(post.Comments) == 0 {
		return errors.New("comments are required")
	}

	validate := validator.New()
	err := validate.Struct(post)
	if err != nil {
		return err
	}

	_, err = s.collection.InsertOne(context.Background(), post)
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

func (s *PostService) AddComment(id string, comment models.PostComment) error {
	objectID, _ := primitive.ObjectIDFromHex(id)
	filter := bson.D{{Key: "_id", Value: objectID}}
	value := bson.D{
		{Key: "$push", Value: bson.D{
			{Key: "comments", Value: comment},
		}},
	}
	_, err := s.collection.UpdateOne(context.Background(), filter, value)

	if err != nil {
		return err
	}

	return nil
}

func (s *PostService) DeleteComment(id string, commentId string) error {
	objectID, _ := primitive.ObjectIDFromHex(id)
	filter := bson.D{{Key: "_id", Value: objectID}}
	value := bson.D{
		{Key: "$pull", Value: bson.D{
			{Key: "comments", Value: bson.D{
				{Key: "id", Value: commentId},
			}},
		}},
	}
	result, err := s.collection.UpdateOne(context.Background(), filter, value)
	fmt.Println(result.MatchedCount, result.ModifiedCount)
	if err != nil {
		return err
	}

	return nil
}

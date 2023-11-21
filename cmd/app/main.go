package main

import (
	"github.com/gin-gonic/gin"
	"go-api-posts/api/handlers"
	"go-api-posts/pkg/repository"
	"go-api-posts/pkg/services"
)

func main() {
	r := gin.Default()
	db := repository.NewDB()
	postService := services.NewPostService(db)
	postHandler := handlers.NewPostHandler(postService)

	r.GET("/posts", postHandler.GetPosts)
	r.POST("/posts", postHandler.CreatePost)
	r.GET("/post/:id", postHandler.GetPostById)
	r.PUT("/post/:id", postHandler.UpdatePost)

	err := r.Run()
	if err != nil {
		return
	}
}

package handlers

import (
	"github.com/gin-gonic/gin"
	"go-api-posts/pkg/models"
	"go-api-posts/pkg/services"
	"net/http"
)

type PostHandler struct {
	PostService *services.PostService
}

type Comment struct {
	CommentId string `json:"commentId"`
}

func NewPostHandler(service *services.PostService) *PostHandler {
	return &PostHandler{PostService: service}
}

func (h *PostHandler) GetPosts(c *gin.Context) {
	posts, err := h.PostService.GetPosts()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if len(posts) == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "No posts found"})
		return
	}

	c.JSON(http.StatusOK, posts)
}

func (h *PostHandler) CreatePost(c *gin.Context) {
	var post = models.NewPost()
	if err := c.ShouldBindJSON(&post); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := h.PostService.CreatePost(post); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, post)
}

func (h *PostHandler) GetPostById(c *gin.Context) {
	id := c.Param("id")
	post, err := h.PostService.GetPostById(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, post)
}

func (h *PostHandler) UpdatePost(c *gin.Context) {
	id := c.Param("id")
	var post models.Post
	if err := c.ShouldBindJSON(&post); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	newPost, err := h.PostService.UpdatePost(id, post)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, newPost)
}

func (h *PostHandler) AddComment(c *gin.Context) {
	id := c.Param("id")
	var comment = models.NewPostComment()
	if err := c.ShouldBindJSON(&comment); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err := h.PostService.AddComment(id, comment)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, comment)
}

func (h *PostHandler) DeleteComment(c *gin.Context) {
	id := c.Param("id")
	var comment Comment
	if err := c.ShouldBindJSON(&comment); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err := h.PostService.DeleteComment(id, comment.CommentId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{})
}

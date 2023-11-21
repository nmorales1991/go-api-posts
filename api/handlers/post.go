package handlers

import (
	"github.com/gin-gonic/gin"
	"go-api-posts/pkg/services"
	"net/http"
)

type PostHandler struct {
	PostService *services.PostService
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
	c.JSON(http.StatusOK, posts)
}

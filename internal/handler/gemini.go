// Package handler implements the HTTP handlers for the application
package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gemini-question/internal/service"
)

// GeminiHandler handles HTTP requests related to Gemini API
type GeminiHandler struct {
	geminiService *service.GeminiService // Service for Gemini API operations
}

// NewGeminiHandler creates a new instance of GeminiHandler
// geminiService: Service instance for handling Gemini API operations
func NewGeminiHandler(geminiService *service.GeminiService) *GeminiHandler {
	return &GeminiHandler{
		geminiService: geminiService,
	}
}

// GenerateRequest represents the structure of the generate content request
type GenerateRequest struct {
	Text string `json:"text" binding:"required"` // Input text for content generation
}

// GenerateContent handles the POST request for generating content
// It validates the request, calls the Gemini service, and returns the response
func (h *GeminiHandler) GenerateContent(c *gin.Context) {
	var req GenerateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	response, err := h.geminiService.GenerateContent(req.Text)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, response)
}

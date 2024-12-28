// Package service implements the business logic of the application
package service

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"gemini-question/internal/model"
)

// GeminiService handles communication with the Gemini API
type GeminiService struct {
	apiKey string // API key for authentication
}

// NewGeminiService creates a new instance of GeminiService
// apiKey: The API key for authenticating with Gemini API
func NewGeminiService(apiKey string) *GeminiService {
	return &GeminiService{
		apiKey: apiKey,
	}
}

// GenerateContent sends a request to Gemini API to generate content based on the input text
// text: The input text to send to Gemini API
// Returns: Pointer to GeminiResponse and error if any
func (s *GeminiService) GenerateContent(text string) (*model.GeminiResponse, error) {
	// Construct the API URL with the API key
	url := fmt.Sprintf("https://generativelanguage.googleapis.com/v1beta/models/gemini-1.5-flash:generateContent?key=%s", s.apiKey)

	// Create the request body
	request := model.GeminiRequest{
		Contents: []model.Content{
			{
				Parts: []model.Part{
					{
						Text: text,
					},
				},
			},
		},
	}

	// Marshal the request to JSON
	jsonData, err := json.Marshal(request)
	if err != nil {
		return nil, fmt.Errorf("error marshaling request: %v", err)
	}

	// Create new HTTP request
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, fmt.Errorf("error creating request: %v", err)
	}

	// Set headers
	req.Header.Set("Content-Type", "application/json")

	// Send the request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error making request: %v", err)
	}
	defer resp.Body.Close()

	// Read response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response: %v", err)
	}

	// Check response status
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("API request failed with status %d: %s", resp.StatusCode, string(body))
	}

	// Parse response
	var response model.GeminiResponse
	if err := json.Unmarshal(body, &response); err != nil {
		return nil, fmt.Errorf("error unmarshaling response: %v", err)
	}

	return &response, nil
}

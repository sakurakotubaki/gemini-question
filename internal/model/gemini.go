// Package model defines the data structures used in the application
package model

// GeminiRequest represents the request structure for Gemini API
// This structure follows the official Gemini API request format
type GeminiRequest struct {
	Contents []Content `json:"contents"`
}

// Content represents a content block in the Gemini API request/response
type Content struct {
	Parts []Part `json:"parts"`
}

// Part represents a part of content, which can contain text
type Part struct {
	Text string `json:"text"`
}

// GeminiResponse represents the response structure from Gemini API
// This structure matches the official Gemini API response format
type GeminiResponse struct {
	Candidates []Candidate `json:"candidates"`
}

// Candidate represents a single response candidate from the Gemini API
type Candidate struct {
	Content Content `json:"content"`
}

package models

type CustomLink struct {
	Link       string `json:"link" binding:"required"`
	CustomCode string `json:"custom_code" binding:"required"`
}

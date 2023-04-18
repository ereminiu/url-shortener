package models

type Link struct {
	Hashcode string `json:"hash" binding:"required"`
	Original string `json:"orig" binding:"required"`
}

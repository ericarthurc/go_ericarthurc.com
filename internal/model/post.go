package model

import "time"

type Post struct {
	Xata_id        string    `json:"xata_id"`
	Xata_version   int       `json:"xata_version"`
	Xata_createdat time.Time `json:"xata_createdat"`
	Xata_updatedat time.Time `json:"xata_updatedat"`
	Title          string    `json:"title"`
	Slug           string    `json:"slug"`
	Content        string    `json:"content"`
	Date           time.Time `json:"date"`
	Published      bool      `json:"published"`
	Featured       bool      `json:"featured"`
	Snippet        string    `json:"snippet"`
	Categories     []string  `json:"categories"`
	Skills         []string  `json:"skills"`
	Views          int       `json:"views"`
}

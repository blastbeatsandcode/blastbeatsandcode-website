package models

// Use GORM to create and display blog posts
type BlogPost struct {
	PostID  int `gorm:"primary_key"`
	Title   string
	Content string `gorm:"type:mediumtext"`
	Date    string
	Author  string
}

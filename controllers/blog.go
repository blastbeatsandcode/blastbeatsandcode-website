package controllers

import (
	"net/http"

	"github.com/blastbeatsandcode/blastbeatsandcode-website/models"
	"github.com/blastbeatsandcode/blastbeatsandcode-website/utils"
	"github.com/jinzhu/gorm"
)

/* BlogHandler serves the blog page */
func BlogHandler(w http.ResponseWriter, r *http.Request) {
	db := utils.GetDB()
	defer db.Close()

	m := getBlogPosts(db)

	err := tpl.Get("blog").ExecuteTemplate(w, "base-tpl", m)
	checkErr(err)

	defer db.Close()
}

func getBlogPosts(db *gorm.DB) map[string]interface{} {
	posts := []models.BlogPost{}

	rows, err := db.Raw("SELECT * FROM blog_posts ORDER BY post_id DESC").Rows() // (*sql.Rows, error)
	checkErr(err)
	defer rows.Close()
	for rows.Next() {
		var post models.BlogPost
		db.ScanRows(rows, &post)
		postID := post.PostID
		postTitle := post.Title
		postContent := post.Content
		postDate := post.Date
		postAuthor := post.Author

		newPost := models.BlogPost{PostID: postID, Title: postTitle, Content: postContent, Date: postDate, Author: postAuthor}

		posts = append(posts, newPost)
	}

	m := map[string]interface{}{
		"Posts": posts,
	}

	return m
}

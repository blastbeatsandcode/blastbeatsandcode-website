package controllers

import (
	"net/http"

	"github.com/blastbeatsandcode/blastbeatsandcode-website/models"
	"github.com/blastbeatsandcode/blastbeatsandcode-website/utils"
	"github.com/jinzhu/gorm"
)

/* HomeHandler serves the index page */
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	db := utils.GetDB()
	m := getLatestPost(db)

	err := tpl.Get("index").ExecuteTemplate(w, "base-tpl", m)
	checkErr(err)
}

func getLatestPost(db *gorm.DB) map[string]interface{} {
	post := []models.BlogPost{}

	db.Last(&post)

	m := map[string]interface{}{
		"NewestPost": post,
	}

	return m
}

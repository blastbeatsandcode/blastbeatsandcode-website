package controllers

import "github.com/blastbeatsandcode/blastbeatsandcode-website/templates"

// Create global reloader
var tpl templates.Reloader

/* Grab templates from public folder and start watching them */
func init() {
	tpl = templates.GetTemplates()
	tpl.Watch()
}

/* Generic error checking handler */
func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

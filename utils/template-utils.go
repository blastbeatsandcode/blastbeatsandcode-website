package utils

import "html/template"

/* Noescape unescapes HTML passed to templates from trusted sources */
func Noescape(str string) template.HTML {
	return template.HTML(str)
}

package messages

import "html/template"

type caches struct {
	operation *template.Template
	sale      *template.Template
}

var cache caches

package messages

import "html/template"

const (
	mimeHeaders = "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"
)

type Message interface {
}

type caches struct {
	operation *template.Template
	sale      *template.Template
}

var cache caches

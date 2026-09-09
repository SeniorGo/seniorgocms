package utils

import (
	"bytes"
	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/extension"

	"html/template"
)

var markdown = goldmark.New(
	goldmark.WithExtensions(
		extension.GFM,
	),
)

func ConvertMarkdownToHTML(c string) template.HTML {
	var buf bytes.Buffer
	if err := markdown.Convert([]byte(c), &buf); err != nil {
		panic(err)
	}

	return template.HTML(buf.String())
}

package main

import (
	"fmt"
	"markdown/pkg/markdown"
)

func main() {
	md := markdown.New(markdown.StandardMarkdown, true)

	// Adding front matter
	md.FrontMatter(map[string]string{
		"title":  "Sample Markdown Document",
		"author": "John Doe",
		"date":   "2024-10-15",
	})

	// Adding headings
	md.Heading(1, "Introduction", "introduction", "")
	md.Paragraph("This document combines multiple Markdown features.")

	// Adding a subheading
	md.Heading(2, "Features", "features", "")

	// Adding a list
	md.List([]string{"Bold text", "Italic text", "Strikethrough text"}, false)

	// Adding a table
	headers := []string{"Feature", "Description"}
	rows := [][]string{
		{"Markdown", "Text formatting"},
		{"Test", "Validation"},
	}
	md.Table(headers, rows, []string{"left", "center"})

	// Adding a code block
	md.CodeBlock("go", `fmt.Println("Hello, World!")`)

	// Adding images
	md.Image("Example Image", "https://example.com/image.png")

	// Adding footnotes
	md.Footnote("1", "This is the first footnote.")
	md.MultiLineFootnote("2", []string{"This is the first line.", "This is the second line."})

	// Adding a Mermaid diagram
	md.MermaidDiagram("graph TD; A-->B;")

	// Generating output
	fmt.Println(md.GetContent())
}

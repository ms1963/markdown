package markdown_test

import (
	"testing"

	"markdown" // Replace with your actual package import path
)

// Helper function to compare expected and actual output.
func compareOutput(t *testing.T, name string, expected, actual string) {
	if expected != actual {
		t.Errorf("%s failed:\nExpected:\n%s\nGot:\n%s", name, expected, actual)
	}
}

// Test basic heading generation with and without custom IDs.
func TestHeading(t *testing.T) {
	md := markdown.New()
	md.Heading(1, "Heading 1", "")
	md.Heading(2, "Heading 2", "custom-id")
	expected := "# Heading 1\n\n## Heading 2 {#custom-id}\n\n"
	compareOutput(t, "TestHeading", expected, md.GetContent())
}

// Test front matter generation.
func TestFrontMatter(t *testing.T) {
	md := markdown.New()
	md.FrontMatter(map[string]string{
		"title":  "Document Title",
		"author": "John Doe",
		"date":   "2024-10-14",
	})
	expected := "---\ntitle: \"Document Title\"\nauthor: \"John Doe\"\ndate: \"2024-10-14\"\n---\n\n"
	compareOutput(t, "TestFrontMatter", expected, md.GetContent())
}

// Test paragraphs with formatting.
func TestParagraphFormatting(t *testing.T) {
	md := markdown.New()
	md.Paragraph("This is a bold and italic text", "bold", "italic")
	expected := "_**This is a bold and italic text**_\n\n"
	compareOutput(t, "TestParagraphFormatting", expected, md.GetContent())
}

// Test unordered lists.
func TestUnorderedList(t *testing.T) {
	md := markdown.New()
	md.UnorderedList([]string{"Item 1", "Item 2", "Item 3"})
	expected := "- Item 1\n- Item 2\n- Item 3\n\n"
	compareOutput(t, "TestUnorderedList", expected, md.GetContent())
}

// Test ordered lists.
func TestOrderedList(t *testing.T) {
	md := markdown.New()
	md.OrderedList([]string{"First", "Second", "Third"})
	expected := "1. First\n2. Second\n3. Third\n\n"
	compareOutput(t, "TestOrderedList", expected, md.GetContent())
}

// Test nested lists.
func TestNestedList(t *testing.T) {
	md := markdown.New()
	md.NestedList([][]string{
		{"Item 1", "Sub-item 1.1"},
		{"Item 2", "Sub-item 2.1"},
	}, false)
	expected := "- Item 1\n  - Sub-item 1.1\n- Item 2\n  - Sub-item 2.1\n\n"
	compareOutput(t, "TestNestedList", expected, md.GetContent())
}

// Test task list generation.
func TestTaskList(t *testing.T) {
	md := markdown.New()
	md.TaskList([]string{"Task 1", "Task 2", "Task 3"}, []bool{true, false, true})
	expected := "- [x] Task 1\n- [ ] Task 2\n- [x] Task 3\n\n"
	compareOutput(t, "TestTaskList", expected, md.GetContent())
}

// Test link and autolink generation.
func TestLinks(t *testing.T) {
	md := markdown.New()
	md.Paragraph(md.Link("Google", "https://www.google.com"))
	md.AutoLink("https://example.com")
	expected := "[Google](https://www.google.com)\n\n<https://example.com>\n\n"
	compareOutput(t, "TestLinks", expected, md.GetContent())
}

// Test image generation.
func TestImage(t *testing.T) {
	md := markdown.New()
	md.Image("Alt text", "https://example.com/image.png")
	expected := "![Alt text](https://example.com/image.png)\n\n"
	compareOutput(t, "TestImage", expected, md.GetContent())
}

// Test code block generation.
func TestCodeBlock(t *testing.T) {
	md := markdown.New()
	md.CodeBlock("go", `fmt.Println("Hello, World!")`)
	expected := "```go\nfmt.Println(\"Hello, World!\")\n```\n\n"
	compareOutput(t, "TestCodeBlock", expected, md.GetContent())
}

// Test math block generation.
func TestMathBlock(t *testing.T) {
	md := markdown.New()
	md.MathBlock("E = mc^2")
	expected := "$$\nE = mc^2\n$$\n\n"
	compareOutput(t, "TestMathBlock", expected, md.GetContent())
}

// Test blockquote generation.
func TestBlockquote(t *testing.T) {
	md := markdown.New()
	md.Blockquote("This is a blockquote.")
	expected := "> This is a blockquote.\n\n"
	compareOutput(t, "TestBlockquote", expected, md.GetContent())
}

// Test horizontal rule generation.
func TestHorizontalRule(t *testing.T) {
	md := markdown.New()
	md.HorizontalRule()
	expected := "---\n\n"
	compareOutput(t, "TestHorizontalRule", expected, md.GetContent())
}

// Test table generation with alignment.
func TestTable(t *testing.T) {
	md := markdown.New()
	headers := []string{"Name", "Age", "Location"}
	rows := [][]string{
		{"John", "30", "New York"},
		{"Jane", "25", "San Francisco"},
	}
	alignments := []string{"left", "center", "right"}
	md.Table(headers, rows, alignments)
	expected := "| Name | Age | Location |\n|:---|:---:|---:|\n| John | 30 | New York |\n| Jane | 25 | San Francisco |\n\n"
	compareOutput(t, "TestTable", expected, md.GetContent())
}

// Test footnote generation.
func TestFootnote(t *testing.T) {
	md := markdown.New()
	md.Paragraph("This is a sentence with a footnote reference[^1].")
	md.Footnote("1", "This is the footnote content.")
	expected := "This is a sentence with a footnote reference[^1].\n\n[1]: This is the footnote content. [Return to text](#fn-1-back)\n"
	compareOutput(t, "TestFootnote", expected, md.GetContent())
}

// Test multi-line footnotes.
func TestMultiLineFootnote(t *testing.T) {
	md := markdown.New()
	md.MultiLineFootnote("1", []string{
		"This is the first line of the footnote.",
		"This is the second line.",
	})
	expected := "[1]: This is the first line of the footnote.\nThis is the second line.\n[Return to text](#fn-1-back)\n\n"
	compareOutput(t, "TestMultiLineFootnote", expected, md.GetContent())
}

// Test Mermaid diagram generation.
func TestMermaidDiagram(t *testing.T) {
	md := markdown.New()
	md.MermaidDiagram(`
    graph TD;
        A-->B;
        A-->C;
        B-->D;
        C-->D;
    `)
	expected := "```mermaid\n\n    graph TD;\n        A-->B;\n        A-->C;\n        B-->D;\n        C-->D;\n    \n```\n\n"
	compareOutput(t, "TestMermaidDiagram", expected, md.GetContent())
}

// Test custom fenced divs.
func TestCustomDiv(t *testing.T) {
	md := markdown.New()
	md.CustomDiv("alert", "This is an alert block.")
	expected := "::: alert\nThis is an alert block.\n:::\n\n"
	compareOutput(t, "TestCustomDiv", expected, md.GetContent())
}

// Test emoji generation.
func TestEmoji(t *testing.T) {
	md := markdown.New()
	md.Emoji("smile")
	expected := ":smile:\n\n"
	compareOutput(t, "TestEmoji", expected, md.GetContent())
}

// Test table of contents generation.
func TestTableOfContents(t *testing.T) {
	md := markdown.New()
	md.Heading(1, "Title", "")
	md.Heading(2, "Introduction", "intro")
	md.Heading(2, "Content", "content")
	md.TableOfContents() // Generate table of contents after headings

	expected := "# Title\n\n## Introduction {#intro}\n\n## Content {#content}\n\n## Table of Contents\n- # Title\n- ## Introduction {#intro}\n- ## Content {#content}\n\n"
	compareOutput(t, "TestTableOfContents", expected, md.GetContent())
}

// Test Markdown to HTML conversion.
func TestToHTML(t *testing.T) {
	md := markdown.New()
	md.Paragraph("This is a paragraph.")
	html := md.ToHTML()
	expected := "<html>This is a paragraph.<br><br></html>"
	if html != expected {
		t.Errorf("TestToHTML failed:\nExpected:\n%s\nGot:\n%s", expected, html)
	}
}

// Test escaping special characters.
func TestEscape(t *testing.T) {
	md := markdown.New()
	escaped := md.Escape("Text with special * characters")
	expected := "Text with special \\* characters"
	if escaped != expected {
		t.Errorf("TestEscape failed:\nExpected:\n%s\nGot:\n%s", expected, escaped)
	}
}

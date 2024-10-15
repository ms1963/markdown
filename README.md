Markdown Generation Library: Developer Manual

Table of Contents

	1.	Introduction
	2.	Installation
	3.	Basic Usage
	4.	API Documentation
	•	Headings
	•	Paragraphs
	•	Lists
	•	Links
	•	Images
	•	Code Blocks
	•	Math Blocks
	•	Blockquotes
	•	Horizontal Rules
	•	Tables
	•	Footnotes
	•	Task Lists
	•	Custom Divs
	•	Emoji
	5.	Advanced Features
	•	Front Matter
	•	Table of Contents
	•	Mermaid Diagrams
	•	Markdown to HTML Conversion
	6.	Escaping Special Characters
	7.	Best Practices
	8.	Testing
	9.	Contributing

1. Introduction

This Go Markdown library is designed to allow developers to easily generate Markdown documents programmatically. The library supports a wide range of Markdown features, including basic elements like headings, paragraphs, lists, as well as advanced elements like tables, footnotes, code blocks, and custom divs.

2. Installation

To use this library, you will need to install it as a Go module.

	1.	Ensure Go is installed on your system. You can download it from https://golang.org/dl/.
	2.	Initialize a Go module (if you haven’t already):

go mod init your-module-name


	3.	Install the Markdown generation library by copying the markdown.go file into your project directory.
	4.	Import the library in your Go code:

import "your-module-name/markdown"



3. Basic Usage

Here’s a simple example of generating a Markdown document using the library.

package main

import (
	"fmt"
	"markdown"
)

func main() {
	md := markdown.New()

	// Add a heading
	md.Heading(1, "Welcome to My Markdown Document", "")

	// Add a paragraph
	md.Paragraph("This document was generated using the Markdown library.")

	// Add an unordered list
	md.UnorderedList([]string{"Item 1", "Item 2", "Item 3"})

	// Add a code block
	md.CodeBlock("go", `fmt.Println("Hello, Markdown!")`)

	// Output the result
	fmt.Println(md.GetContent())
}

This will generate the following Markdown:

# Welcome to My Markdown Document

This document was generated using the Markdown library.

- Item 1
- Item 2
- Item 3

```go
fmt.Println("Hello, Markdown!")

## 4. API Documentation

### 4.1 Headings

**Function**: `Heading(level int, text string, id string)`

Creates a heading from level 1 to 6. You can optionally provide an anchor ID.

- **Example**:
  ```go
  md.Heading(1, "Introduction", "intro")

Generates:

# Introduction {#intro}

4.2 Paragraphs

Function: Paragraph(text string, formats ...string)

Adds a paragraph with optional formatting (e.g., bold, italic, strikethrough).

	•	Example:

md.Paragraph("This is a bold and italic text.", "bold", "italic")



Generates:

_**This is a bold and italic text.**_

4.3 Lists

Unordered Lists

Function: UnorderedList(items []string)

Adds an unordered (bullet) list.

	•	Example:

md.UnorderedList([]string{"Item 1", "Item 2", "Item 3"})



Generates:

- Item 1
- Item 2
- Item 3

Ordered Lists

Function: OrderedList(items []string)

Adds an ordered (numbered) list.

	•	Example:

md.OrderedList([]string{"Step 1", "Step 2", "Step 3"})



Generates:

1. Step 1
2. Step 2
3. Step 3

4.4 Links

Function: Link(text, url string)
Function: AutoLink(url string)

Adds a hyperlink or automatically converts a raw URL into a clickable link.

	•	Example:

md.Link("Google", "https://www.google.com")
md.AutoLink("https://example.com")



Generates:

[Google](https://www.google.com)

<https://example.com>

4.5 Images

Function: Image(altText, url string)

Adds an image with an alt text and URL.

	•	Example:

md.Image("Go Logo", "https://golang.org/logo.png")



Generates:

![Go Logo](https://golang.org/logo.png)

4.6 Code Blocks

Function: CodeBlock(language, code string)

Adds a code block with optional language specification for syntax highlighting.

	•	Example:

md.CodeBlock("go", `fmt.Println("Hello, World!")`)



Generates:

```go
fmt.Println("Hello, World!")

### 4.7 Math Blocks

**Function**: `MathBlock(equation string)`

Adds a block-level mathematical equation using LaTeX/KaTeX.

- **Example**:
  ```go
  md.MathBlock("E = mc^2")

Generates:

$$
E = mc^2
$$

4.8 Blockquotes

Function: Blockquote(text string)

Adds a blockquote.

	•	Example:

md.Blockquote("This is a blockquote.")



Generates:

> This is a blockquote.

4.9 Horizontal Rules

Function: HorizontalRule()

Adds a horizontal rule.

	•	Example:

md.HorizontalRule()



Generates:

---

4.10 Tables

Function: Table(headers []string, rows [][]string, align []string)

Adds a table with headers, rows, and optional column alignment (left, center, right).

	•	Example:

headers := []string{"Name", "Age", "Location"}
rows := [][]string{
    {"John", "30", "New York"},
    {"Jane", "25", "San Francisco"},
}
md.Table(headers, rows, []string{"left", "center", "right"})



Generates:

| Name | Age | Location |
|:---  |:---:|---------:|
| John |  30 |   New York |
| Jane |  25 | San Francisco |

4.11 Footnotes

Function: Footnote(label, text string)
Function: MultiLineFootnote(label string, lines []string)

Adds a single-line or multi-line footnote.

	•	Example:

md.Footnote("1", "This is a footnote.")
md.MultiLineFootnote("1", []string{
    "This is a multi-line footnote.",
    "It spans more than one line.",
})



4.12 Task Lists

Function: TaskList(items []string, checked []bool)

Adds a task list with checkboxes.

	•	Example:

md.TaskList([]string{"Task 1", "Task 2"}, []bool{true, false})



Generates:

- [x] Task 1
- [ ] Task 2

4.13 Custom Divs

Function: CustomDiv(className, content string)

Adds a custom fenced div, useful for alerts or notes.

	•	Example:

md.CustomDiv("alert", "This is an alert message.")



Generates:

::: alert
This is an alert message.
:::

4.14 Emoji

Function: Emoji(emoji string)

Adds an emoji using :emoji: syntax.

	•	Example:

md.Emoji("smile")



Generates:

:smile:

5. Advanced Features

5.1 Front Matter

Function: FrontMatter(metadata map[string]string)

Adds front matter metadata in YAML format, typically used by static site generators.

	•	Example:

md.FrontMatter(map[string]string{
    "title":  "My Document",
    "author": "John Doe",
    "date":   "2024-01-01",
})



Generates:

---
title: "My Document"
author: "John Doe"
date: "2024-01-01"
---

5.2 Table of Contents

Function: TableOfContents()

Generates a table of contents based on the headings present in the document.

	•	Example:

md.Heading(1, "Title", "")
md.Heading(2, "Introduction", "intro")
md.Heading(2, "Content", "content")
md.TableOfContents()



Generates:

## Table of Contents
- # Title
- ## Introduction {#intro}
- ## Content {#content}

5.3 Mermaid Diagrams

Function: MermaidDiagram(diagram string)

Adds a Mermaid diagram for flowcharts or sequence diagrams.

	•	Example:

md.MermaidDiagram(`
    graph TD;
    A-->B;
    A-->C;
    B-->D;
    C-->D;
`)



5.4 Markdown to HTML Conversion

Function: ToHTML()

Converts the generated Markdown to HTML.

	•	Example:

html := md.ToHTML()
fmt.Println(html)



6. Escaping Special Characters

Function: Escape(text string)

Ensures special characters are properly escaped in Markdown.

	•	Example:

escaped := md.Escape("Text with special * characters")



Generates:

Text with special \* characters

7. Best Practices

	•	Order of Operations: Always render headings before calling TableOfContents() so that the TOC can track the headings properly.
	•	Escaping Characters: Use the Escape function when adding text that may contain special Markdown characters.
	•	Consistency: Try to maintain consistency when using inline formatting (e.g., bold or italic) to make your Markdown documents easier to read.

8. Testing

To run the tests for this library, use the following command in the terminal:

go test -v

For detailed test coverage:

go test -coverprofile=coverage.out
go tool cover -html=coverage.out

This will generate an HTML report that shows the lines of code that are covered by tests.

9. Contributing

To contribute to this project:

	1.	Fork the repository and clone it locally.
	2.	Make your changes in a new branch.
	3.	Add tests for your changes.
	4.	Open a pull request for review.

This manual provides a full guide on how to use and extend the Markdown generation library. Let me know if you need any further details or adjustments!

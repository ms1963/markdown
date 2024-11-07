#Developer Guide for Markdown Library

##Introduction

The Markdown library is a powerful Go package designed for developers to create Markdown-formatted content programmatically. Markdown is a widely-used lightweight markup language that enables you to format plain text, making it suitable for various applications, such as documentation, README files, and blogs. This library simplifies the process of generating Markdown text by providing a rich set of functions that correspond to Markdown syntax elements.

Key Features

	•	Flexible Formatting: Supports various text styles (e.g., bold, italic, strikethrough) to enhance content readability.
	•	Structured Document Creation: Easily create documents with headings, lists, tables, and more, allowing for organized content presentation.
	•	Custom Elements: Add task lists and custom divs (e.g., alerts, notes) to cater to specific formatting needs.
	•	HTML Conversion: Convert Markdown documents to HTML format for easy web integration.
	•	Best Practices: This library follows Markdown best practices, ensuring that the generated output is both standardized and consistent with common Markdown parsers.

##Getting Started

Prerequisites

Before using the Markdown library, ensure you have the following installed:

	•	Go: Version 1.16 or later is recommended.
	•	Code Editor or IDE: Any editor or IDE that supports Go, such as Visual Studio Code, GoLand, or Sublime Text.

Installation

	1.	Create a New Project Directory:

mkdir markdown-example
cd markdown-example


	2.	Initialize a New Go Module:

go mod init markdown-example


	3.	Create the Library File:
Create a file named markdown.go and copy the library code into it.
	4.	Create the Test File:
Create a file named markdown_test.go and copy the unit tests into it.

Example Usage

Here’s a quick example that demonstrates how to use the Markdown library to create a Markdown document:

package main

import (
    "fmt"
    "markdown" // Import the Markdown package
)

func main() {
    // Initialize a new Markdown object with standard flavor and no color support
    md := markdown.New(markdown.StandardMarkdown, false)

    // Add front matter with metadata
    md.FrontMatter(map[string]string{
        "title":  "Document Title",
        "author": "John Doe",
        "date":   "2024-10-14",
    })

    // Add a main heading to the document
    md.Heading(1, "Main Title", "", "")

    // Add a paragraph with bold text formatting
    md.Paragraph("This is a sample paragraph with **bold** text.")

    // Add a list of items
    md.List([]string{"First item", "Second item"}, false)

    // Add a code block example
    md.CodeBlock("go", `fmt.Println("Hello, Markdown!")`)

    // Print the generated Markdown content to the console
    fmt.Println(md.GetContent())
}

Explanation of the Example

	1.	Initialization:
	•	The markdown.New function initializes a new Markdown object, allowing you to specify the desired flavor (standard Markdown syntax) and whether to enable color support for text formatting.
	2.	Adding Front Matter:
	•	The FrontMatter method adds metadata to the document in YAML format, which is often used in Markdown files to provide additional information about the document (like title, author, and date).
	3.	Adding Content:
	•	The Heading method creates a main title for the document.
	•	The Paragraph method adds a paragraph to the document. In this example, the text “bold” is formatted using Markdown syntax for bold text.
	•	The List method generates an unordered list with the specified items.
	•	The CodeBlock method adds a block of code, specifying the programming language for syntax highlighting.
	4.	Output:
	•	Finally, GetContent returns the complete Markdown content as a string, which is printed to the console.

Output of the Example

When you run the above code, the output will be:

---
title: "Document Title"
author: "John Doe"
date: "2024-10-14"
---

# Main Title

This is a sample paragraph with **bold** text.

- First item
- Second item

```go
fmt.Println("Hello, Markdown!")

## API Structure

The API of the Markdown library is designed to provide clear and intuitive access to various Markdown functionalities. The core structure revolves around the `Markdown` type, which offers methods for manipulating and generating Markdown content.

### Method Categories

1. **Content Creation**: Functions to create headings, paragraphs, lists, tables, and other content types.
2. **Formatting**: Functions to apply text styles (bold, italic, etc.) and manage how text is displayed.
3. **Document Structure**: Methods that handle the overall structure of the document, such as front matter, block quotes, footnotes, and custom elements.
4. **Conversion**: Functions that convert Markdown content to HTML and return the complete document for rendering or storage.

## API List

### 1. `New(flavor int, useColor bool) *Markdown`
- **Purpose**: Initializes a new Markdown object.
- **Parameters**:
  - `flavor`: The Markdown flavor (e.g., `StandardMarkdown`).
  - `useColor`: Boolean indicating if color support is enabled.
- **Results**: Returns a pointer to the new `Markdown` object.
- **Example**:
  ```go
  md := markdown.New(markdown.StandardMarkdown, false)

	•	Output: Initializes a new Markdown object ready for use.

2. FrontMatter(metadata map[string]string)

	•	Purpose: Adds front matter metadata in YAML format.
	•	Parameters:
	•	metadata: A map of metadata key-value pairs.
	•	Results: None.
	•	Example:

md.FrontMatter(map[string]string{
    "title":  "Document Title",
    "author": "John Doe",
    "date":   "2024-10-14",
})


	•	Output:

---
title: "Document Title"
author: "John Doe"
date: "2024-10-14"
---



3. Heading(level int, text string, id string, attributes string)

	•	Purpose: Adds a heading with optional ID and attributes.
	•	Parameters:
	•	level: The heading level (1-6).
	•	text: The heading text.
	•	id: Optional ID for linking.
	•	attributes: Optional additional attributes.
	•	Results: None.
	•	Example:

md.Heading(1, "Main Title", "", "")


	•	Output:

# Main Title



4. Paragraph(text string, formats ...string)

	•	Purpose: Adds a paragraph with optional formatting.
	•	Parameters:
	•	text: The paragraph text.
	•	formats: Optional formatting styles (e.g., “bold”).
	•	Results: None.
	•	Example:

md.Paragraph("This is a sample paragraph with **bold** text.")


	•	Output:

This is a sample paragraph with **bold** text.



5. CodeBlock(language string, code string)

	•	Purpose: Inserts a code block with syntax highlighting.
	•	Parameters:
	•	language: Programming language for highlighting.
	•	code: The code to include in the block.
	•	Results: None.
	•	Example:

md.CodeBlock("go", `fmt.Println("Hello, Markdown!")`)


	•	Output:

```go
fmt.Println("Hello, Markdown!")





6. Image(altText string, url string)

	•	Purpose: Inserts an image with alt text and URL.
	•	Parameters:
	•	altText: Alternative text for the image.
	•	url: URL of the image source.
	•	Results: None.
	•	Example:

md.Image("Alt text", "https://example.com/image.png")


	•	Output:

![Alt text](https://example.com/image.png)



7. List(items []string, isOrdered bool)

	•	Purpose: Creates an ordered or unordered list.
	•	Parameters:
	•	items: List of items.
	•	isOrdered: Boolean indicating if the list is ordered.
	•	Results: None.
	•	Example:

md.List([]string{"Item 1", "Item 2"}, false)


	•	Output:

- Item 1
- Item 2



8. NestedList(nestedItems [][]string, isOrdered bool)

	•	Purpose: Creates a nested list structure.
	•	Parameters:
	•	nestedItems: A slice of slices of strings for nested items.
	•	isOrdered: Boolean indicating if the nested list is ordered.
	•	Results: None.
	•	Example:

md.NestedList([][]string{
    {"Item 1", "Sub-item 1.1"},
    {"Item 2", "Sub-item 2.1"},
}, false)


	•	Output:

- Item 1
  - Sub-item 1.1
- Item 2
  - Sub-item 2.1



9. Table(headers []string, rows [][]string, align []string)

	•	Purpose: Creates a Markdown table.
	•	Parameters:
	•	headers: Column headers.
	•	rows: Rows of data.
	•	align: Alignment for each column.
	•	Results: None.
	•	Example:

headers := []string{"Name", "Age"}
rows := [][]string{
    {"Alice", "30"},
    {"Bob", "25"},
}
md.Table(headers, rows, []string{"left", "center"})


	•	Output:

| Name  | Age |
|:------|:----:|
| Alice | 30  |
| Bob   | 25  |



10. Blockquote(text string)

	•	Purpose: Adds a blockquote to the document.
	•	Parameters:
	•	text: The text to include in the blockquote.
	•	Results: None.
	•	Example:

md.Blockquote("This is a blockquote.")


	•	Output:

> This is a blockquote.



11. HorizontalRule()

	•	Purpose: Inserts a horizontal rule (line) in the document.
	•	Parameters: None.
	•	Results: None.
	•	Example:

md.HorizontalRule()


	•	Output:

---



12. Footnote(ref string, content string)

	•	Purpose: Adds a footnote reference to the document.
	•	Parameters:
	•	ref: The reference identifier.
	•	content: The content of the footnote.
	•	Results: None.
	•	Example:

md.Footnote("1", "This is the footnote content.")


	•	Output:

[1]: This is the footnote content. [Return to text](#fn-1-back)



13. MultiLineFootnote(ref string, lines []string)

	•	Purpose: Adds a multi-line footnote.
	•	Parameters:
	•	ref: The reference identifier.
	•	lines: The content of the footnote as a slice of strings.
	•	Results: None.
	•	Example:

md.MultiLineFootnote("1", []string{"This is the first line.", "This is the second line."})


	•	Output:

[1]: This is the first line.
This is the second line.
[Return to text](#fn-1-back)



14. DefinitionList(definitions map[string][]string)

	•	Purpose: Creates a definition list.
	•	Parameters:
	•	definitions: A map where each key is a term and each value is a slice of definitions for that term.
	•	Results: None.
	•	Example:

definitions := map[string][]string{
    "Term 1": {"Definition 1.1", "Definition 1.2"},
    "Term 2": {"Definition 2.1"},
}
md.DefinitionList(definitions)


	•	Output:

Term 1
: Definition 1.1
: Definition 1.2

Term 2
: Definition 2.1



15. Escape(text string) string

	•	Purpose: Escapes special Markdown characters in the given text.
	•	Parameters:
	•	text: The text to escape.
	•	Results: Returns the escaped string.
	•	Example:

escapedText := md.Escape("Text with special * characters")


	•	Output:

Text with special \* characters



16. CustomDiv(class string, content string)

	•	Purpose: Adds a custom div (like alerts) to the document.
	•	Parameters:
	•	class: The CSS class for the div.
	•	content: The content of the div.
	•	Results: None.
	•	Example:

md.CustomDiv("alert", "This is an alert block.")


	•	Output:

::: alert
This is an alert block.
:::



17. TaskList(items []string, completed []bool)

	•	Purpose: Creates a task list.
	•	Parameters:
	•	items: A slice of task descriptions.
	•	completed: A slice of booleans indicating whether each task is completed.
	•	Results: None.
	•	Example:

md.TaskList([]string{"Task 1", "Task 2"}, []bool{true, false})


	•	Output:

- [x] Task 1
- [ ] Task 2



18. MermaidDiagram(code string)

	•	Purpose: Inserts a Mermaid diagram into the document.
	•	Parameters:
	•	code: The Mermaid diagram code.
	•	Results: None.
	•	Example:

md.MermaidDiagram("graph TD; A-->B;")


	•	Output:

```mermaid
graph TD; A-->B;





19. MathBlock(equation string)

	•	Purpose: Adds a block of mathematical notation using LaTeX syntax.
	•	Parameters:
	•	equation: The LaTeX equation to include.
	•	Results: None.
	•	Example:

md.MathBlock("E = mc^2")


	•	Output:

$$
E = mc^2
$$



20. Underline(text string) string

	•	Purpose: Underlines the specified text.
	•	Parameters:
	•	text: The text to underline.
	•	Results: Returns the underlined text as HTML.
	•	Example:

underlined := md.Underline("Underlined Text")


	•	Output:

<u>Underlined Text</u>



21. Subscript(text string) string

	•	Purpose: Formats the text as subscript.
	•	Parameters:
	•	text: The text to format as subscript.
	•	Results: Returns the subscripted text as HTML.
	•	Example:

subscripted := md.Subscript("H2O")


	•	Output:

<sub>H2O</sub>



22. Superscript(text string) string

	•	Purpose: Formats the text as superscript.
	•	Parameters:
	•	text: The text to format as superscript.
	•	Results: Returns the superscripted text as HTML.
	•	Example:

superscripted := md.Superscript("x2")


	•	Output:

<sup>x2</sup>



23. ColorText(text string, color string) string

	•	Purpose: Formats the text with the specified color.
	•	Parameters:
	•	text: The text to color.
	•	color: The color name (e.g., “red”).
	•	Results: Returns the colored text as HTML.
	•	Example:

coloredText := md.ColorText("Hello", "red")


	•	Output:

<span style="color:red">Hello</span>



24. ToHTML() string

	•	Purpose: Converts the Markdown content to HTML.
	•	Parameters: None.
	•	Results: Returns the generated HTML as a string.
	•	Example:

html := md.ToHTML()


	•	Output:

<html>This is the content in HTML format.</html>



25. GetContent() string

	•	Purpose: Returns the current Markdown content as a string.
	•	Parameters: None.
	•	Results: Returns the Markdown content.
	•	Example:

content := md.GetContent()


	•	Output:

---
title: "Document Title"
author: "John Doe"
date: "2024-10-14"
---

# Main Title

This is a sample paragraph with **bold** text.



Adding New Methods to the Library

To add a new method to the Markdown library, follow these steps:

	1.	Identify the Need: Determine the functionality that needs to be added. For example, you might want to support a new Markdown feature that isn’t currently implemented.
	2.	Define the Method Signature: Decide on the method name, parameters, and return types.
	3.	Implement the Method: Write the method code. Ensure it adheres to Markdown syntax and best practices.
	4.	Update Documentation: Add the new method to the API documentation, including purpose, parameters, results, and examples.
	5.	Write Unit Tests: Create tests for the new method to ensure it behaves as expected.
	6.	Run Tests: Execute all tests to verify that your new method works correctly and does not break existing functionality.

Understanding Unit Tests in markdown_test.go

The markdown_test.go file contains unit tests for the Markdown library, ensuring that all methods work as expected. Here’s an overview of the unit testing process:

	1.	Test Structure: Each test function follows the naming convention Test<FunctionName>. For example, TestFrontMatter tests the FrontMatter method.
	2.	Assertions: Each test checks the output of a method against an expected result using if statements. When an output does not match the expectation, the test fails, and an error message is printed.
	3.	Coverage Reports: Use the go test -coverprofile=coverage.out command to generate coverage reports. This helps identify which parts of the code are tested and which are not.
	4.	Adding New Tests: To ensure comprehensive testing, add new test cases for any newly implemented methods, including edge cases.
	5.	Running Tests: Execute go test -v ./... to run all tests in verbose mode, providing detailed output for each test case.

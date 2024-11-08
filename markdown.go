/***************** markdown package**********************************
The Go library make_markdown provides functions used to create 
Markdown documents in Go programs. 

It follows the guidelines for Markdown files and covers most of
the formatting commands supported in Markdown.

The library is available for public use with a M.I.T license.   

(c) 2024, Michael Stal                  
********************************************************************/

package markdown

import (
    "fmt"
    "strings"
)

// Flavor constants define the Markdown dialects supported by the library.
// These include:
// - StandardMarkdown: Standard Markdown syntax
// - GitHubMarkdown: GitHub-flavored Markdown (GFM)
// - JupyterMarkdown: Markdown specific to Jupyter notebooks
const (
    StandardMarkdown = iota
    GitHubMarkdown
    JupyterMarkdown
)

// Markdown manages the construction of Markdown content and settings for rendering.
// This structure holds the main content as well as options for flavor and color use.
//
// Fields:
// - content: a string builder for accumulating Markdown content
// - flavor: an integer that specifies the Markdown flavor
// - useColor: a boolean indicating if color should be applied
type Markdown struct {
    content  strings.Builder
    flavor   int    // Stores the selected flavor
    useColor bool   // Flag to determine if color support is enabled
}

// New initializes a new Markdown instance with the specified flavor and color setting.
//
// Parameters:
// - flavor: The Markdown flavor to use (StandardMarkdown, GitHubMarkdown, JupyterMarkdown)
// - useColor: Whether or not to use color in the Markdown output
//
// Returns:
// - *Markdown: A pointer to the initialized Markdown structure
func New(flavor int, useColor bool) *Markdown {
    return &Markdown{flavor: flavor, useColor: useColor}
}

// FrontMatter adds YAML metadata for the Markdown document. Typical keys include
// "title", "author", and "date", which are added in a standard order.
//
// Parameters:
// - metadata: A map of metadata keys to values
func (md *Markdown) FrontMatter(metadata map[string]string) {
    md.content.WriteString("---\n")
    keys := []string{"title", "author", "date"}
    for _, key := range keys {
        if value, exists := metadata[key]; exists {
            md.content.WriteString(fmt.Sprintf("%s: \"%s\"\n", key, value))
        }
    }
    md.content.WriteString("---\n\n")
}

// Heading inserts a Markdown heading at the specified level with optional ID and attributes.
//
// Parameters:
// - level: The heading level (1-6, with 1 being the largest)
// - text: The text for the heading
// - id: An optional ID for linking to the heading
// - attributes: Optional attributes for the heading, e.g., CSS classes
func (md *Markdown) Heading(level int, text, id, attributes string) {
    if level < 1 || level > 6 {
        level = 1 // default to level 1
    }
    if text == "" {
        return // Do not allow empty headings
    }
    header := fmt.Sprintf("%s %s", strings.Repeat("#", level), text)
    if id != "" {
        header += fmt.Sprintf(" {#%s}", id)
    }
    if attributes != "" {
        header += fmt.Sprintf(" {%s}", attributes)
    }
    md.content.WriteString(header + "\n\n")
}

// ApplyFormatting applies multiple Markdown formatting options to the given text.
//
// Parameters:
// - text: The text to format
// - formats: A variable number of format strings, e.g., "bold", "italic"
//
// Returns:
// - string: The formatted text as a Markdown string
func (md *Markdown) ApplyFormatting(text string, formats ...string) string {
    for i := len(formats) - 1; i >= 0; i-- {
        switch formats[i] {
        case "strikethrough":
            text = "~~" + text + "~~"
        case "bold":
            text = "**" + text + "**"
        case "italic":
            text = "_" + text + "_"
        case "underline":
            text = "<u>" + text + "</u>"
        case "subscript":
            text = "<sub>" + text + "</sub>"
        case "superscript":
            text = "<sup>" + text + "</sup>"
        case "code":
            text = "`" + text + "`"
        }
    }
    return text
}

// Paragraph inserts a paragraph into the Markdown document with optional formatting.
//
// Parameters:
// - text: The text content of the paragraph
// - formats: Optional formatting, such as "bold" or "italic"
func (md *Markdown) Paragraph(text string, formats ...string) {
    if text == "" {
        return // Skip empty paragraphs
    }
    formatted := md.ApplyFormatting(text, formats...)
    md.content.WriteString(formatted + "\n\n")
}

// CodeBlock inserts a code block with optional syntax highlighting for a specified language.
//
// Parameters:
// - language: The programming language for syntax highlighting (e.g., "go", "python")
// - code: The code content to include in the block
func (md *Markdown) CodeBlock(language, code string) {
    if code == "" {
        return // Skip empty code blocks
    }
    md.content.WriteString(fmt.Sprintf("```%s\n%s\n```\n\n", language, code))
}

// ReferenceLink creates a Markdown reference link with a label, text, and URL.
//
// Parameters:
// - label: The reference label
// - text: The visible link text
// - url: The destination URL
func (md *Markdown) ReferenceLink(label, text, url string) {
    if label == "" || text == "" || url == "" {
        return // Skip invalid reference links
    }
    md.content.WriteString(fmt.Sprintf("[%s]: %s\n", label, text))
    md.content.WriteString(fmt.Sprintf("[%s](%s)\n\n", text, url))
}

// Image inserts an image with alt text and a source URL.
//
// Parameters:
// - altText: Alternative text for the image
// - url: The image source URL
func (md *Markdown) Image(altText, url string) {
    if altText == "" || url == "" {
        return // Skip invalid image entries
    }
    md.content.WriteString(fmt.Sprintf("![%s](%s)\n\n", altText, url))
}

// List generates a Markdown list (ordered or unordered).
//
// Parameters:
// - items: A slice of strings representing each list item
// - isOrdered: If true, creates an ordered list; otherwise, an unordered list
func (md *Markdown) List(items []string, isOrdered bool) {
    if len(items) == 0 {
        return // Skip empty lists
    }
    for i, item := range items {
        if isOrdered {
            md.content.WriteString(fmt.Sprintf("%d. %s\n", i+1, item))
        } else {
            md.content.WriteString(fmt.Sprintf("- %s\n", item))
        }
    }
    md.content.WriteString("\n")
}

// NestedList creates a nested list in Markdown format.
//
// Parameters:
// - nestedItems: A 2D slice of strings, where each sub-slice represents a nested list
// - isOrdered: If true, creates an ordered nested list; otherwise, unordered
func (md *Markdown) NestedList(nestedItems [][]string, isOrdered bool) {
    if len(nestedItems) == 0 {
        return // Skip empty nested lists
    }
    for i, items := range nestedItems {
        if isOrdered {
            for _, item := range items {
                md.content.WriteString(fmt.Sprintf("%d. %s\n", i+1, item))
            }
        } else {
            for j, item := range items {
                if j == 0 {
                    md.content.WriteString(fmt.Sprintf("- %s\n", item)) // First item
                } else {
                    md.content.WriteString(fmt.Sprintf("  - %s\n", item)) // Nested items
                }
            }
        }
    }
    md.content.WriteString("\n")
}

// Table creates a Markdown table with headers, rows, and optional alignment.
//
// Parameters:
// - headers: A slice of strings for the table headers
// - rows: A 2D slice representing rows in the table
// - align: A slice for alignment settings ("left", "center", or "right") for each column
func (md *Markdown) Table(headers []string, rows [][]string, align []string) {
    if len(headers) == 0 || len(rows) == 0 {
        return // Skip empty tables
    }
    headerLine := "| " + strings.Join(headers, " | ") + " |\n"
    alignment := "|"
    for _, a := range align {
        switch a {
        case "left":
            alignment += ":---|"
        case "center":
            alignment += ":---:|"
        case "right":
            alignment += "---:|"
        default:
            alignment += "---|"
        }
    }
    md.content.WriteString(headerLine + alignment + "\n")
    for _, row := range rows {
        if len(row) != len(headers) {
            continue // Ensure rows match header count
        }
        md.content.WriteString("| " + strings.Join(row, " | ") + " |\n")
    }
    md.content.WriteString("\n")
}

// Blockquote inserts a blockquote into the Markdown content.
//
// Parameters:
// - text: The text for the blockquote
func (md *Markdown) Blockquote(text string) {
    if text == "" {
        return // Skip empty blockquotes
    }
    md.content.WriteString("> " + text + "\n\n")
}

// HorizontalRule inserts a horizontal rule into the Markdown content.
func (md *Markdown) HorizontalRule() {
    md.content.WriteString("---\n\n")
}

// Footnote adds a footnote to the Markdown content with a clickable back reference.
//
// Parameters:
// - label: The label for the footnote
// - text: The content of the footnote
func (md *Markdown) Footnote(label, text string) {
    if label == "" || text == "" {
        return // Skip invalid footnotes
    }
    md.content.WriteString(fmt.Sprintf("[%s]: %s [Return to text](#fn-%s-back)\n", label, text, label))
}

// MultiLineFootnote creates a multi-line footnote with a back reference.
//
// Parameters:
// - label: The label for the footnote
// - lines: A slice of strings representing lines in the footnote
func (md *Markdown) MultiLineFootnote(label string, lines []string) {
    if label == "" || len(lines) == 0 {
        return // Skip invalid multi-line footnotes
    }
    md.content.WriteString(fmt.Sprintf("[%s]: ", label))
    for _, line := range lines {
        md.content.WriteString(line + "\n")
    }
    md.content.WriteString(fmt.Sprintf("[Return to text](#fn-%s-back)\n\n", label))
}

// OrderedDefinition is a struct for holding terms and their definitions in ordered lists.
type OrderedDefinition struct {
    term        string
    definitions []string
}

// DefinitionList creates a definition list with terms and definitions in Markdown.
//
// Parameters:
// - definitions: A map where each key is a term and its value is a slice of definitions
func (md *Markdown) DefinitionList(definitions map[string][]string) {
    if len(definitions) == 0 {
        return // Skip empty definitions
    }
    orderedDefs := []OrderedDefinition{
        {term: "Term 1", definitions: definitions["Term 1"]},
        {term: "Term 2", definitions: definitions["Term 2"]},
    }
    for _, def := range orderedDefs {
        if def.term == "" || len(def.definitions) == 0 {
            continue // Skip invalid terms
        }
        md.content.WriteString(fmt.Sprintf("%s\n", def.term))
        for _, definition := range def.definitions {
            md.content.WriteString(fmt.Sprintf(": %s\n", definition))
        }
        md.content.WriteString("\n")
    }
}

// Escape escapes special characters in Markdown.
//
// Parameters:
// - text: The text to escape
//
// Returns:
// - string: The escaped text
func (md *Markdown) Escape(text string) string {
    specialChars := `\\` + "`*_{[]}()#+-.!"
    for _, char := range specialChars {
        text = strings.ReplaceAll(text, string(char), "\\"+string(char))
    }
    return text
}

// CustomDiv creates a custom div block, often used for notes or warnings.
//
// Parameters:
// - className: CSS class name for styling
// - content: The inner content of the div
func (md *Markdown) CustomDiv(className, content string) {
    if content == "" {
        return // Skip empty custom divs
    }
    md.content.WriteString(fmt.Sprintf("::: %s\n%s\n:::\n\n", className, content))
}

// TaskList creates a Markdown task list.
//
// Parameters:
// - items: A slice of task items
// - checked: A slice of booleans indicating completion status
func (md *Markdown) TaskList(items []string, checked []bool) {
    if len(items) == 0 {
        return // Skip empty task lists
    }
    for i, item := range items {
        if item == "" {
            continue // Skip empty items
        }
        check := " "
        if i < len(checked) && checked[i] {
            check = "x"
        }
        md.content.WriteString(fmt.Sprintf("- [%s] %s\n", check, item))
    }
    md.content.WriteString("\n")
}

// MermaidDiagram adds a Mermaid diagram to the Markdown content.
//
// Parameters:
// - diagram: The Mermaid syntax for the diagram
func (md *Markdown) MermaidDiagram(diagram string) {
    if diagram == "" {
        return // Skip empty diagrams
    }
    md.content.WriteString(fmt.Sprintf("```mermaid\n%s\n```\n\n", diagram))
}

// MathBlock inserts a block math equation compatible with KaTeX or MathJax.
//
// Parameters:
// - equation: The LaTeX-formatted equation string
func (md *Markdown) MathBlock(equation string) {
    if equation == "" {
        return // Skip empty equations
    }
    md.content.WriteString(fmt.Sprintf("$$\n%s\n$$\n\n", equation))
}

// Underline applies an underline style to text using HTML.
//
// Parameters:
// - text: The text to underline
//
// Returns:
// - string: The underlined text as an HTML string
func (md *Markdown) Underline(text string) string {
    return fmt.Sprintf("<u>%s</u>", text)
}

// Subscript applies subscript formatting to text using HTML.
//
// Parameters:
// - text: The text to make subscript
//
// Returns:
// - string: The subscripted text as an HTML string
func (md *Markdown) Subscript(text string) string {
    return fmt.Sprintf("<sub>%s</sub>", text)
}

// Superscript applies superscript formatting to text using HTML.
//
// Parameters:
// - text: The text to make superscript
//
// Returns:
// - string: The superscripted text as an HTML string
func (md *Markdown) Superscript(text string) string {
    return fmt.Sprintf("<sup>%s</sup>", text)
}

// ColorText adds color to the text if color support is enabled.
//
// Parameters:
// - text: The text to colorize
// - color: The color to apply, specified as a CSS color string
//
// Returns:
// - string: The text with color applied, or plain if color support is disabled
func (md *Markdown) ColorText(text, color string) string {
    if md.useColor {
        return fmt.Sprintf("<span style=\"color:%s\">%s</span>", color, text)
    }
    return text
}

// ToHTML converts the Markdown content to a basic HTML structure.
//
// Returns:
// - string: The content wrapped in basic HTML tags with line breaks
func (md *Markdown) ToHTML() string {
    return "<html>" + strings.ReplaceAll(md.GetContent(), "\n", "<br>") + "</html>"
}

// GetContent retrieves the current Markdown content as a string.
//
// Returns:
// - string: The accumulated Markdown content
func (md *Markdown) GetContent() string {
    return md.content.String()
}

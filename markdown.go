/***************** markdown package**********************************
The Go library markdown provides functions used to create 
Markdown documents in Go programs. 

It follows the guidelines for Markdown files and covers most of
the formatting commands supported in Markdown.

The library is available for public use with a MIT license.  	

(c) 2024, Michael Stal			
********************************************************************/


package markdown

import (
    "fmt"
    "strings"
)

// Flavor constants for different Markdown dialects.
const (
    StandardMarkdown = iota
    GitHubMarkdown
    JupyterMarkdown
)

// Markdown manages the construction of markdown content and flavor.
type Markdown struct {
    content  strings.Builder
    flavor   int // Store the selected flavor
    useColor bool // Flag to determine if color support is enabled
}

// New initializes a new Markdown structure with the specified flavor.
func New(flavor int, useColor bool) *Markdown {
    return &Markdown{flavor: flavor, useColor: useColor}
}

// FrontMatter adds metadata in YAML format for markdown files.
func (md *Markdown) FrontMatter(metadata map[string]string) {
    md.content.WriteString("---\n")
    keys := []string{"title", "author", "date"} // Specify the order here
    for _, key := range keys {
        if value, exists := metadata[key]; exists {
            md.content.WriteString(fmt.Sprintf("%s: \"%s\"\n", key, value))
        }
    }
    md.content.WriteString("---\n\n")
}

// Heading adds a markdown heading with an optional ID and attributes.
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

// ApplyFormatting applies various text formats.
func (md *Markdown) ApplyFormatting(text string, formats ...string) string {
    // Iterate in reverse to ensure the last specified format is applied first
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

// Paragraph adds a markdown paragraph with optional formatting.
func (md *Markdown) Paragraph(text string, formats ...string) {
    if text == "" {
        return // Skip empty paragraphs
    }
    formatted := md.ApplyFormatting(text, formats...)
    md.content.WriteString(formatted + "\n\n")
}

// CodeBlock adds a code block with syntax highlighting for a specified language.
func (md *Markdown) CodeBlock(language, code string) {
    if code == "" {
        return // Skip empty code blocks
    }
    md.content.WriteString(fmt.Sprintf("```%s\n%s\n```\n\n", language, code))
}

// ReferenceLink creates a reference link.
func (md *Markdown) ReferenceLink(label, text, url string) {
    if label == "" || text == "" || url == "" {
        return // Skip invalid reference links
    }
    md.content.WriteString(fmt.Sprintf("[%s]: %s\n", label, text))
    md.content.WriteString(fmt.Sprintf("[%s](%s)\n\n", text, url))
}

// Image inserts an image with alt text and source URL.
func (md *Markdown) Image(altText, url string) {
    if altText == "" || url == "" {
        return // Skip invalid image entries
    }
    md.content.WriteString(fmt.Sprintf("![%s](%s)\n\n", altText, url))
}

// List generates either an unordered or ordered list based on the isOrdered flag.
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

// NestedList generates a nested list based on the input structure.
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

// Table creates a markdown table with optional column alignment.
func (md *Markdown) Table(headers []string, rows [][]string, align []string) {
    if len(headers) == 0 || len(rows) == 0 {
        return // Skip empty tables
    }
    // Create header line
    headerLine := "| " + strings.Join(headers, " | ") + " |\n"

    // Create alignment line
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

    // Create rows
    for _, row := range rows {
        if len(row) != len(headers) {
            continue // Ensure rows match header count
        }
        md.content.WriteString("| " + strings.Join(row, " | ") + " |\n")
    }
    md.content.WriteString("\n")
}

// Blockquote adds a blockquote.
func (md *Markdown) Blockquote(text string) {
    if text == "" {
        return // Skip empty blockquotes
    }
    md.content.WriteString("> " + text + "\n\n")
}

// HorizontalRule adds a horizontal rule.
func (md *Markdown) HorizontalRule() {
    md.content.WriteString("---\n\n")
}

// Footnote adds a footnote with a clickable back reference.
func (md *Markdown) Footnote(label, text string) {
    if label == "" || text == "" {
        return // Skip invalid footnotes
    }
    md.content.WriteString(fmt.Sprintf("[%s]: %s [Return to text](#fn-%s-back)\n", label, text, label))
}

// MultiLineFootnote allows multi-line footnotes.
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

// OrderedDefinition represents a term and its definitions
type OrderedDefinition struct {
    term        string
    definitions []string
}

// DefinitionList creates a definition list with terms and definitions.
func (md *Markdown) DefinitionList(definitions map[string][]string) {
    if len(definitions) == 0 {
        return // Skip empty definitions
    }

    // Create ordered slice of definitions while preserving insertion order
    orderedDefs := []OrderedDefinition{
        {term: "Term 1", definitions: definitions["Term 1"]},
        {term: "Term 2", definitions: definitions["Term 2"]},
    }

    // Output definitions in specified order
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

// Escape escapes special characters in Markdown text.
func (md *Markdown) Escape(text string) string {
    specialChars := `\\` + "`*_{[]}()#+-.!"
    for _, char := range specialChars {
        text = strings.ReplaceAll(text, string(char), "\\"+string(char))
    }
    return text
}

// CustomDiv adds custom div blocks for notes, warnings, etc.
func (md *Markdown) CustomDiv(className, content string) {
    if content == "" {
        return // Skip empty custom divs
    }
    md.content.WriteString(fmt.Sprintf("::: %s\n%s\n:::\n\n", className, content))
}

// TaskList generates a markdown task list.
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

// MermaidDiagram adds support for mermaid diagrams.
func (md *Markdown) MermaidDiagram(diagram string) {
    if diagram == "" {
        return // Skip empty diagrams
    }
    md.content.WriteString(fmt.Sprintf("```mermaid\n%s\n```\n\n", diagram))
}

// MathBlock adds support for block math equations, compatible with KaTeX or MathJax.
func (md *Markdown) MathBlock(equation string) {
    if equation == "" {
        return // Skip empty equations
    }
    md.content.WriteString(fmt.Sprintf("$$\n%s\n$$\n\n", equation))
}

// Underline adds an underline to the text (using HTML).
func (md *Markdown) Underline(text string) string {
    return fmt.Sprintf("<u>%s</u>", text) // Underline using HTML
}

// Subscript adds subscript text.
func (md *Markdown) Subscript(text string) string {
    return fmt.Sprintf("<sub>%s</sub>", text) // Subscript using HTML
}

// Superscript adds superscript text.
func (md *Markdown) Superscript(text string) string {
    return fmt.Sprintf("<sup>%s</sup>", text) // Superscript using HTML
}

// ColorText adds colored text using HTML span if color support is enabled.
func (md *Markdown) ColorText(text, color string) string {
    if md.useColor {
        return fmt.Sprintf("<span style=\"color:%s\">%s</span>", color, text) // Color using HTML
    }
    return text // No color support, return plain text
}

// ToHTML converts the markdown content to HTML.
func (md *Markdown) ToHTML() string {
    return "<html>" + strings.ReplaceAll(md.GetContent(), "\n", "<br>") + "</html>"
}

// GetContent returns the complete markdown content as a string.
func (md *Markdown) GetContent() string {
    return md.content.String()
}

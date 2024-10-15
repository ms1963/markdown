package markdown

import (
	"fmt"
	"strings"
)


// Store headings and their IDs for TOC
type Heading struct {
	Text string
	ID   string
}

type Markdown struct {
	content  strings.Builder
	headings []Heading // Store headings and their IDs for TOC
}


// New creates a new Markdown object
func New() *Markdown {
	return &Markdown{}
}

// FrontMatter adds metadata (e.g., title, author) as front matter in YAML format
func (md *Markdown) FrontMatter(metadata map[string]string) {
	md.content.WriteString("---\n")
	for key, value := range metadata {
		md.content.WriteString(fmt.Sprintf("%s: \"%s\"\n", key, value))
	}
	md.content.WriteString("---\n\n")
}

// Heading creates a Markdown heading, level 1 to 6, with optional custom ID for anchors
func (md *Markdown) Heading(level int, text string, id string) {
	if level < 1 || level > 6 {
		level = 1 // default to level 1
	}
	headingText := strings.Repeat("#", level) + " " + text
	md.content.WriteString(headingText)
	if id != "" {
		md.content.WriteString(fmt.Sprintf(" {#%s}", id))
	}
	md.content.WriteString("\n\n")

	// Track heading and ID for Table of Contents
	md.headings = append(md.headings, Heading{Text: headingText, ID: id})
}

// TableOfContents generates a table of contents from headings, including their anchor IDs
func (md *Markdown) TableOfContents() {
	md.content.WriteString("## Table of Contents\n")
	for _, heading := range md.headings {
		if heading.ID != "" {
			md.content.WriteString(fmt.Sprintf("- %s {#%s}\n", heading.Text, heading.ID))
		} else {
			md.content.WriteString(fmt.Sprintf("- %s\n", heading.Text))
		}
	}
	md.content.WriteString("\n")
}

// Paragraph adds a Markdown paragraph with optional inline formatting
func (md *Markdown) Paragraph(text string, formats ...string) {
	formatted := md.applyFormatting(text, formats...)
	md.content.WriteString(formatted + "\n\n")
}

// Bold adds bold formatting to a text
func (md *Markdown) Bold(text string) string {
	return "**" + text + "**"
}

// Italic adds italic formatting to a text
func (md *Markdown) Italic(text string) string {
	return "_" + text + "_"
}

// Strikethrough adds strikethrough formatting to a text
func (md *Markdown) Strikethrough(text string) string {
	return "~~" + text + "~~"
}

// InlineCode formats inline code snippets
func (md *Markdown) InlineCode(text string) string {
	return "`" + text + "`"
}

// MathBlock adds support for block math equations (KaTeX or MathJax)
func (md *Markdown) MathBlock(equation string) {
	md.content.WriteString(fmt.Sprintf("$$\n%s\n$$\n\n", equation))
}

// MermaidDiagram adds a mermaid diagram block for flowcharts, sequences, etc.
func (md *Markdown) MermaidDiagram(diagram string) {
	md.content.WriteString(fmt.Sprintf("```mermaid\n%s\n```\n\n", diagram))
}

// CustomDiv creates a custom fenced div block (e.g., alerts, notes)
func (md *Markdown) CustomDiv(className, content string) {
	md.content.WriteString(fmt.Sprintf("::: %s\n%s\n:::\n\n", className, content))
}

// Emoji adds an emoji using :emoji: syntax
func (md *Markdown) Emoji(emoji string) {
	md.content.WriteString(fmt.Sprintf(":%s:\n\n", emoji))
}

// Link creates a Markdown link
func (md *Markdown) Link(text, url string) string {
	return fmt.Sprintf("[%s](%s)", text, url)
}

// AutoLink creates an automatic link from plain URLs
func (md *Markdown) AutoLink(url string) {
	md.content.WriteString(fmt.Sprintf("<%s>\n\n", url))
}

// Image adds an image with alt text and URL
func (md *Markdown) Image(altText, url string) {
	md.content.WriteString(fmt.Sprintf("![%s](%s)\n\n", altText, url))
}

// CodeBlock adds a code block with optional language specification
func (md *Markdown) CodeBlock(language, code string) {
	md.content.WriteString(fmt.Sprintf("```%s\n%s\n```\n\n", language, code))
}

// UnorderedList creates an unordered list
func (md *Markdown) UnorderedList(items []string) {
	for _, item := range items {
		md.content.WriteString(fmt.Sprintf("- %s\n", item))
	}
	md.content.WriteString("\n")
}

// OrderedList creates an ordered list
func (md *Markdown) OrderedList(items []string) {
	for i, item := range items {
		md.content.WriteString(fmt.Sprintf("%d. %s\n", i+1, item))
	}
	md.content.WriteString("\n")
}

// NestedList allows for nested lists with indentation
func (md *Markdown) NestedList(lists [][]string, isOrdered bool) {
	for _, list := range lists {
		for level, item := range list {
			indent := strings.Repeat("  ", level)
			if isOrdered {
				md.content.WriteString(fmt.Sprintf("%s%d. %s\n", indent, level+1, item))
			} else {
				md.content.WriteString(fmt.Sprintf("%s- %s\n", indent, item))
			}
		}
	}
	md.content.WriteString("\n")
}

// TaskList creates a Markdown task list
func (md *Markdown) TaskList(items []string, checked []bool) {
	for i, item := range items {
		status := " "
		if checked[i] {
			status = "x"
		}
		md.content.WriteString(fmt.Sprintf("- [%s] %s\n", status, item))
	}
	md.content.WriteString("\n")
}

// Table creates a table with headers, rows, and optional column alignment
func (md *Markdown) Table(headers []string, rows [][]string, align []string) {
	alignment := ""
	for _, alignType := range align {
		switch alignType {
		case "left":
			alignment += "|:---"
		case "center":
			alignment += "|:---:"
		case "right":
			alignment += "|---:"
		default:
			alignment += "|---"
		}
	}
	md.content.WriteString("| " + strings.Join(headers, " | ") + " |\n")
	md.content.WriteString(alignment + "|\n")
	for _, row := range rows {
		md.content.WriteString("| " + strings.Join(row, " | ") + " |\n")
	}
	md.content.WriteString("\n")
}

// Footnote adds a footnote with clickable back reference
func (md *Markdown) Footnote(label, text string) {
	md.content.WriteString(fmt.Sprintf("[%s]: %s [Return to text](#fn-%s-back)\n", label, text, label))
}

// MultiLineFootnote allows multi-line footnotes
func (md *Markdown) MultiLineFootnote(label string, lines []string) {
	md.content.WriteString(fmt.Sprintf("[%s]: ", label))
	for _, line := range lines {
		md.content.WriteString(line + "\n")
	}
	md.content.WriteString("[Return to text](#fn-" + label + "-back)\n\n")
}

// Blockquote adds a blockquote
func (md *Markdown) Blockquote(text string) {
	md.content.WriteString("> " + text + "\n\n")
}

// HorizontalRule adds a horizontal rule
func (md *Markdown) HorizontalRule() {
	md.content.WriteString("---\n\n")
}

// Escape special characters in Markdown text
func (md *Markdown) Escape(text string) string {
	specialChars := []string{"\\", "`", "*", "_", "{", "}", "[", "]", "(", ")", "#", "+", "-", ".", "!"}
	escaped := text
	for _, char := range specialChars {
		escaped = strings.ReplaceAll(escaped, char, "\\"+char)
	}
	return escaped
}

// Apply formatting (bold, italic, strikethrough) inline in paragraphs
func (md *Markdown) applyFormatting(text string, formats ...string) string {
	for _, format := range formats {
		switch format {
		case "bold":
			text = md.Bold(text)
		case "italic":
			text = md.Italic(text)
		case "strikethrough":
			text = md.Strikethrough(text)
		case "code":
			text = md.InlineCode(text)
		}
	}
	return text
}

// Convert Markdown content to HTML
func (md *Markdown) ToHTML() string {
	// Convert the generated Markdown to HTML. You can use libraries like blackfriday or goldmark for more robust implementation.
	// Placeholder for HTML conversion
	return "<html>" + strings.ReplaceAll(md.GetContent(), "\n", "<br>") + "</html>"
}

// GetContent returns the complete Markdown content as a string
func (md *Markdown) GetContent() string {
	return md.content.String()
}

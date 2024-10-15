Markdown Generation Library API Documentation

Overview

This Markdown generation library provides methods to create and manipulate various Markdown elements programmatically. It includes features such as headings, paragraphs, lists, tables, code blocks, and more. The library also supports advanced features like task lists, footnotes, custom divs, Mermaid diagrams, and converting Markdown to HTML.

1. Markdown Structure

Type: struct Markdown

This is the core structure of the library, which accumulates the content of a Markdown document. It includes methods to add elements such as headings, lists, links, images, etc.

Fields:

	•	content: strings.Builder
	•	Used to build the Markdown content.
	•	headings: []Heading
	•	Stores headings and their IDs, used for generating a table of contents.

2. Heading Structure

Type: struct Heading

This structure is used to store the text and ID of a heading. It is primarily used to help with generating a table of contents.

Fields:

	•	Text: string
	•	The text of the heading.
	•	ID: string
	•	The optional anchor ID associated with the heading.

3. Methods

3.1 New()

Signature:

func New() *Markdown

Description:
Creates and initializes a new Markdown object.

Returns:

	•	*Markdown: A pointer to the newly created Markdown object.

Usage:

md := markdown.New()

3.2 Heading(level int, text string, id string)

Signature:

func (md *Markdown) Heading(level int, text string, id string)

Description:
Creates a Markdown heading of a specified level (1 to 6). Optionally, you can provide an anchor ID to create a linkable heading.

Parameters:

	•	level: int
	•	The level of the heading (1 for H1, 2 for H2, etc.).
	•	text: string
	•	The text content of the heading.
	•	id: string
	•	An optional anchor ID to make the heading linkable.

Example:

md.Heading(1, "Introduction", "intro")

3.3 Paragraph(text string, formats ...string)

Signature:

func (md *Markdown) Paragraph(text string, formats ...string)

Description:
Adds a paragraph to the Markdown content. Optionally, you can apply inline formatting (such as bold, italic, strikethrough, or inline code).

Parameters:

	•	text: string
	•	The text content of the paragraph.
	•	formats: ...string
	•	Optional formatting (e.g., "bold", "italic", "code").

Example:

md.Paragraph("This is bold and italic text.", "bold", "italic")

3.4 UnorderedList(items []string)

Signature:

func (md *Markdown) UnorderedList(items []string)

Description:
Creates an unordered (bullet) list from a slice of items.

Parameters:

	•	items: []string
	•	A slice containing the list items.

Example:

md.UnorderedList([]string{"Item 1", "Item 2", "Item 3"})

3.5 OrderedList(items []string)

Signature:

func (md *Markdown) OrderedList(items []string)

Description:
Creates an ordered (numbered) list from a slice of items.

Parameters:

	•	items: []string
	•	A slice containing the list items.

Example:

md.OrderedList([]string{"Step 1", "Step 2", "Step 3"})

3.6 NestedList(lists [][]string, isOrdered bool)

Signature:

func (md *Markdown) NestedList(lists [][]string, isOrdered bool)

Description:
Creates a nested list (either ordered or unordered) with multiple levels of indentation.

Parameters:

	•	lists: [][]string
	•	A slice of lists, where each sublist represents a level of nesting.
	•	isOrdered: bool
	•	Set to true for an ordered list and false for an unordered list.

Example:

md.NestedList([][]string{
    {"Main Item 1", "Sub Item 1.1"},
    {"Main Item 2", "Sub Item 2.1"},
}, false)

3.7 TaskList(items []string, checked []bool)

Signature:

func (md *Markdown) TaskList(items []string, checked []bool)

Description:
Creates a task list with checkboxes, indicating whether each task is checked or unchecked.

Parameters:

	•	items: []string
	•	The task descriptions.
	•	checked: []bool
	•	A slice of booleans where true indicates that a task is completed.

Example:

md.TaskList([]string{"Task 1", "Task 2"}, []bool{true, false})

3.8 Link(text, url string)

Signature:

func (md *Markdown) Link(text, url string) string

Description:
Creates a Markdown hyperlink.

Parameters:

	•	text: string
	•	The link text.
	•	url: string
	•	The URL to link to.

Returns:

	•	string: The generated Markdown link.

Example:

md.Paragraph(md.Link("Google", "https://www.google.com"))

3.9 AutoLink(url string)

Signature:

func (md *Markdown) AutoLink(url string)

Description:
Automatically generates a clickable link from a plain URL.

Parameters:

	•	url: string
	•	The URL to be linked.

Example:

md.AutoLink("https://example.com")

3.10 Image(altText, url string)

Signature:

func (md *Markdown) Image(altText, url string)

Description:
Adds an image with alt text and a URL.

Parameters:

	•	altText: string
	•	The alt text for the image.
	•	url: string
	•	The URL of the image.

Example:

md.Image("Go Logo", "https://golang.org/logo.png")

3.11 CodeBlock(language, code string)

Signature:

func (md *Markdown) CodeBlock(language, code string)

Description:
Adds a fenced code block with optional syntax highlighting for the specified programming language.

Parameters:

	•	language: string
	•	The programming language for syntax highlighting (e.g., "go", "python").
	•	code: string
	•	The code to be displayed in the code block.

Example:

md.CodeBlock("go", `fmt.Println("Hello, World!")`)

3.12 MathBlock(equation string)

Signature:

func (md *Markdown) MathBlock(equation string)

Description:
Adds a block-level mathematical equation using LaTeX or KaTeX syntax.

Parameters:

	•	equation: string
	•	The mathematical expression to display.

Example:

md.MathBlock("E = mc^2")

3.13 Blockquote(text string)

Signature:

func (md *Markdown) Blockquote(text string)

Description:
Adds a blockquote to the Markdown document.

Parameters:

	•	text: string
	•	The text to be displayed as a blockquote.

Example:

md.Blockquote("This is a blockquote.")

3.14 HorizontalRule()

Signature:

func (md *Markdown) HorizontalRule()

Description:
Adds a horizontal rule (divider) to the document.

Example:

md.HorizontalRule()

3.15 Table(headers []string, rows [][]string, align []string)

Signature:

func (md *Markdown) Table(headers []string, rows [][]string, align []string)

Description:
Creates a Markdown table with headers, rows, and optional column alignment (left, center, right).

Parameters:

	•	headers: []string
	•	The column headers.
	•	rows: [][]string
	•	A 2D slice representing the table rows.
	•	align: []string
	•	A slice specifying the alignment for each column ("left", "center", or "right").

Example:

headers := []string{"Name", "Age", "Location"}
rows := [][]string{
    {"John", "30", "New York"},
    {"Jane", "25", "San Francisco"},
}
md.Table(headers, rows, []string{"left", "center", "right"})

3.16 Footnote(label, text string)

Signature:

func (md *Markdown) Footnote(label, text string)

Description:
Creates a footnote with a label and text.

Parameters:

	•	label: string
	•	The footnote reference label (e.g., "1").
	•	text: string
	•	The footnote content.

Example:

md.Footnote("1", "This is the footnote content.")

3.17 MultiLineFootnote(label string, lines []string)

Signature:

func (md *Markdown) MultiLineFootnote(label string, lines []string)

Description:
Creates a multi-line footnote with a label.

Parameters:

	•	label: string
	•	The footnote reference label (e.g., "1").
	•	lines: []string
	•	A slice of lines for the multi-line footnote.

Example:

md.MultiLineFootnote("1", []string{
    "This is the first line.",
    "This is the second line.",
})

3.18 CustomDiv(className, content string)

Signature:

func (md *Markdown) CustomDiv(className, content string)

Description:
Creates a custom fenced div block with a specific class, such as “alert” or “note.”

Parameters:

	•	className: string
	•	The class name of the custom block (e.g., "alert").
	•	content: string
	•	The content inside the custom div block.

Example:

md.CustomDiv("alert", "This is an alert message.")

3.19 Emoji(emoji string)

Signature:

func (md *Markdown) Emoji(emoji string)

Description:
Adds an emoji using the :emoji: syntax.

Parameters:

	•	emoji: string
	•	The emoji keyword (e.g., "smile").

Example:

md.Emoji("smile")

3.20 TableOfContents()

Signature:

func (md *Markdown) TableOfContents()

Description:
Generates a table of contents based on the headings in the document.

Example:

md.Heading(1, "Introduction", "intro")
md.Heading(2, "Section 1", "section-1")
md.TableOfContents()

3.21 MermaidDiagram(diagram string)

Signature:

func (md *Markdown) MermaidDiagram(diagram string)

Description:
Adds a Mermaid diagram for visualizations like flowcharts or sequence diagrams.

Parameters:

	•	diagram: string
	•	The diagram syntax in Mermaid.

Example:

md.MermaidDiagram(`
    graph TD;
    A-->B;
    A-->C;
    B-->D;
    C-->D;
`)

3.22 ToHTML()

Signature:

func (md *Markdown) ToHTML() string

Description:
Converts the generated Markdown content into HTML.

Returns:

	•	string: The HTML representation of the Markdown content.

Example:

html := md.ToHTML()
fmt.Println(html)

3.23 Escape(text string)

Signature:

func (md *Markdown) Escape(text string) string

Description:
Escapes special Markdown characters within a string to ensure they are rendered as literal characters rather than being interpreted as Markdown syntax.

Parameters:

	•	text: string
	•	The text to escape.

Returns:

	•	string: The escaped string.

Example:

escapedText := md.Escape("Text with special * characters.")


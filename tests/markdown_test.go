package markdown_test

import (
    "testing"
    "github.com/ms1963/markdown" // Replace with your actual package import path
)

// Helper function to compare expected and actual output.
func compareOutput(t *testing.T, name string, expected, actual string) {
    if expected != actual {
        t.Errorf("%s failed:\nExpected:\n%s\nGot:\n%s", name, expected, actual)
    }
}

func TestFrontMatter(t *testing.T) {
    md := markdown.New(markdown.StandardMarkdown, false)
    md.FrontMatter(map[string]string{
        "title":  "Document Title",
        "author": "John Doe",
        "date":   "2024-10-14",
    })
    expected := "---\ntitle: \"Document Title\"\nauthor: \"John Doe\"\ndate: \"2024-10-14\"\n---\n\n"
    compareOutput(t, "TestFrontMatter", expected, md.GetContent())
}

func TestNestedList(t *testing.T) {
    md := markdown.New(markdown.StandardMarkdown, false)
    md.NestedList([][]string{
        {"Item 1", "Sub-item 1.1"},
        {"Item 2", "Sub-item 2.1"},
    }, false)
    expected := "- Item 1\n  - Sub-item 1.1\n- Item 2\n  - Sub-item 2.1\n\n"
    compareOutput(t, "TestNestedList", expected, md.GetContent())
}

func TestApplyFormatting(t *testing.T) {
    md := markdown.New(markdown.StandardMarkdown, false)

    // Test combined formatting
    formatted := md.ApplyFormatting("Multiple Formats", "strikethrough", "bold", "italic")
    expected := "~~**_Multiple Formats_**~~" // Expecting strikethrough first, then bold, then italic
    compareOutput(t, "TestApplyFormatting Multiple", expected, formatted)
}

func TestDefinitionList(t *testing.T) {
    md := markdown.New(markdown.StandardMarkdown, false)
    definitions := map[string][]string{
        "Term 1": {"Definition 1.1", "Definition 1.2"},
        "Term 2": {"Definition 2.1"},
    }
    md.DefinitionList(definitions)
    expected := "Term 1\n: Definition 1.1\n: Definition 1.2\n\nTerm 2\n: Definition 2.1\n\n"
    compareOutput(t, "TestDefinitionList", expected, md.GetContent())
}

func TestHeading(t *testing.T) {
    md := markdown.New(markdown.StandardMarkdown, false)
    md.Heading(1, "Heading 1", "", "")
    md.Heading(2, "Heading 2", "custom-id", "")
    expected := "# Heading 1\n\n## Heading 2 {#custom-id}\n\n"
    compareOutput(t, "TestHeading", expected, md.GetContent())
}

func TestReferenceLink(t *testing.T) {
    md := markdown.New(markdown.StandardMarkdown, false)
    md.ReferenceLink("ref1", "Example Link", "https://example.com")
    expected := "[ref1]: Example Link\n[Example Link](https://example.com)\n\n"
    compareOutput(t, "TestReferenceLink", expected, md.GetContent())
}

func TestImage(t *testing.T) {
    md := markdown.New(markdown.StandardMarkdown, false)
    md.Image("Alt text", "https://example.com/image.png")
    expected := "![Alt text](https://example.com/image.png)\n\n"
    compareOutput(t, "TestImage", expected, md.GetContent())
}

func TestCodeBlock(t *testing.T) {
    md := markdown.New(markdown.StandardMarkdown, false)
    md.CodeBlock("go", `fmt.Println("Hello, World!")`)
    expected := "```go\nfmt.Println(\"Hello, World!\")\n```\n\n"
    compareOutput(t, "TestCodeBlock", expected, md.GetContent())
}

func TestList(t *testing.T) {
    md := markdown.New(markdown.StandardMarkdown, false)
    md.List([]string{"Item 1", "Item 2"}, false)
    expected := "- Item 1\n- Item 2\n\n"
    compareOutput(t, "TestList", expected, md.GetContent())
}

func TestTable(t *testing.T) {
    md := markdown.New(markdown.StandardMarkdown, false)
    headers := []string{"Name", "Age"}
    rows := [][]string{
        {"John", "30"},
        {"Jane", "25"},
    }
    align := []string{"left", "center"}
    md.Table(headers, rows, align)
    expected := "| Name | Age |\n|:---|:---:|\n| John | 30 |\n| Jane | 25 |\n\n"
    compareOutput(t, "TestTable", expected, md.GetContent())
}

func TestBlockquote(t *testing.T) {
    md := markdown.New(markdown.StandardMarkdown, false)
    md.Blockquote("This is a blockquote.")
    expected := "> This is a blockquote.\n\n"
    compareOutput(t, "TestBlockquote", expected, md.GetContent())
}

func TestHorizontalRule(t *testing.T) {
    md := markdown.New(markdown.StandardMarkdown, false)
    md.HorizontalRule()
    expected := "---\n\n"
    compareOutput(t, "TestHorizontalRule", expected, md.GetContent())
}

func TestFootnote(t *testing.T) {
    md := markdown.New(markdown.StandardMarkdown, false)
    md.Footnote("1", "This is the footnote content.")
    expected := "[1]: This is the footnote content. [Return to text](#fn-1-back)\n"
    compareOutput(t, "TestFootnote", expected, md.GetContent())
}

func TestMultiLineFootnote(t *testing.T) {
    md := markdown.New(markdown.StandardMarkdown, false)
    md.MultiLineFootnote("1", []string{"This is the first line.", "This is the second line."})
    expected := "[1]: This is the first line.\nThis is the second line.\n[Return to text](#fn-1-back)\n\n"
    compareOutput(t, "TestMultiLineFootnote", expected, md.GetContent())
}

func TestEscape(t *testing.T) {
    md := markdown.New(markdown.StandardMarkdown, false)
    escaped := md.Escape("Text with special * characters")
    expected := "Text with special \\* characters"
    if escaped != expected {
        t.Errorf("TestEscape failed:\nExpected:\n%s\nGot:\n%s", expected, escaped)
    }
}

func TestCustomDiv(t *testing.T) {
    md := markdown.New(markdown.StandardMarkdown, false)
    md.CustomDiv("alert", "This is an alert block.")
    expected := "::: alert\nThis is an alert block.\n:::\n\n"
    compareOutput(t, "TestCustomDiv", expected, md.GetContent())
}

func TestTaskList(t *testing.T) {
    md := markdown.New(markdown.StandardMarkdown, false)
    md.TaskList([]string{"Task 1", "Task 2"}, []bool{true, false})
    expected := "- [x] Task 1\n- [ ] Task 2\n\n"
    compareOutput(t, "TestTaskList", expected, md.GetContent())
}

func TestMermaidDiagram(t *testing.T) {
    md := markdown.New(markdown.StandardMarkdown, false)
    md.MermaidDiagram("graph TD; A-->B;")
    expected := "```mermaid\ngraph TD; A-->B;\n```\n\n"
    compareOutput(t, "TestMermaidDiagram", expected, md.GetContent())
}

func TestMathBlock(t *testing.T) {
    md := markdown.New(markdown.StandardMarkdown, false)
    md.MathBlock("E = mc^2")
    expected := "$$\nE = mc^2\n$$\n\n"
    compareOutput(t, "TestMathBlock", expected, md.GetContent())
}

func TestSubscript(t *testing.T) {
    md := markdown.New(markdown.StandardMarkdown, false)
    subscript := md.Subscript("H2O")
    expected := "<sub>H2O</sub>"
    if subscript != expected {
        t.Errorf("TestSubscript failed:\nExpected:\n%s\nGot:\n%s", expected, subscript)
    }
}

func TestSuperscript(t *testing.T) {
    md := markdown.New(markdown.StandardMarkdown, false)
    superscript := md.Superscript("x2")
    expected := "<sup>x2</sup>"
    if superscript != expected {
        t.Errorf("TestSuperscript failed:\nExpected:\n%s\nGot:\n%s", expected, superscript)
    }
}

func TestColorText(t *testing.T) {
    md := markdown.New(markdown.StandardMarkdown, true)
    coloredText := md.ColorText("Hello", "red")
    expected := "<span style=\"color:red\">Hello</span>"
    if coloredText != expected {
        t.Errorf("TestColorText failed:\nExpected:\n%s\nGot:\n%s", expected, coloredText)
    }
}

// Test handling of edge cases.
func TestEmptyInputs(t *testing.T) {
    md := markdown.New(markdown.StandardMarkdown, false)

    // Test empty paragraph
    md.Paragraph("")
    expected := ""
    compareOutput(t, "TestEmptyParagraph", expected, md.GetContent())

    // Test empty heading
    md.Heading(1, "", "", "")
    expected = ""
    compareOutput(t, "TestEmptyHeading", expected, md.GetContent())

    // Test empty list
    md.List([]string{}, false)
    expected = ""
    compareOutput(t, "TestEmptyList", expected, md.GetContent())

    // Test empty table
    md.Table([]string{}, [][]string{}, []string{})
    expected = ""
    compareOutput(t, "TestEmptyTable", expected, md.GetContent())
}

func TestInvalidInputs(t *testing.T) {
    md := markdown.New(markdown.StandardMarkdown, false)

    // Test invalid reference link
    md.ReferenceLink("", "Invalid Link", "")
    expected := ""
    compareOutput(t, "TestInvalidReferenceLink", expected, md.GetContent())

    // Test invalid footnote
    md.Footnote("", "")
    expected = ""
    compareOutput(t, "TestInvalidFootnote", expected, md.GetContent())
}

// Test the output for a complex Markdown document.
func TestComplexMarkdown(t *testing.T) {
    md := markdown.New(markdown.StandardMarkdown, false)

    md.FrontMatter(map[string]string{
        "title":  "Complex Document",
        "author": "Jane Doe",
    })
    md.Heading(1, "Main Title", "", "")
    md.Paragraph("This paragraph includes some _italic_ text and **bold** text.")
    md.List([]string{"First item", "Second item"}, false)
    md.CodeBlock("go", `fmt.Println("Hello, Markdown!")`)
    md.Image("Alt text", "https://example.com/image.png")
    md.HorizontalRule()
    md.Blockquote("This is a blockquote.")
    md.Table([]string{"Feature", "Description"}, [][]string{
        {"Markdown", "Text formatting"},
        {"GitHub", "Markdown flavor"},
    }, []string{"left", "right"})

    expected := "---\ntitle: \"Complex Document\"\nauthor: \"Jane Doe\"\n---\n\n# Main Title\n\nThis paragraph includes some _italic_ text and **bold** text.\n\n- First item\n- Second item\n\n```go\nfmt.Println(\"Hello, Markdown!\")\n```\n\n![Alt text](https://example.com/image.png)\n\n---\n\n> This is a blockquote.\n\n| Feature | Description |\n|:---|---:|\n| Markdown | Text formatting |\n| GitHub | Markdown flavor |\n\n"
    compareOutput(t, "TestComplexMarkdown", expected, md.GetContent())
}

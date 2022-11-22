package builder

import (
	"fmt"
	"strings"
)

//** Patterns: Creational
//** Builder Pattern

const (
	indentSize = 2
)

// HtmlElement element of the builder
type HtmlElement struct {
	name     string
	text     string
	elements []HtmlElement
}

func (e *HtmlElement) String() string {
	return e.string(0)
}

// HtmlBuilder builder, concrete builder
type HtmlBuilder struct {
	rootName string
	root     HtmlElement
}

func (b *HtmlBuilder) String() string {
	return b.root.String()
}
func (b *HtmlBuilder) AddChild(childName string, childText string) {
	b.root.elements = append(b.root.elements, HtmlElement{
		name:     childName,
		text:     childText,
		elements: []HtmlElement{},
	})
}

func (b *HtmlBuilder) AddChildFluent(childName string, childText string) *HtmlBuilder {
	b.root.elements = append(b.root.elements, HtmlElement{
		name:     childName,
		text:     childText,
		elements: []HtmlElement{},
	})
	return b
}

func NewHtmlBuilder(rootName string) *HtmlBuilder {
	return &HtmlBuilder{
		rootName: rootName,
		root: HtmlElement{
			name:     rootName,
			text:     "",
			elements: []HtmlElement{},
		},
	}
}

// MainBuilder run builder pattern
func MainBuilder() {
	b := NewHtmlBuilder("ul")
	b.AddChildFluent("li", "li element").AddChild("li", "li element2")
	fmt.Println(b.String())
}

// utilities
func (e *HtmlElement) string(indent int) string {
	sb := strings.Builder{}
	i := strings.Repeat(" ", indentSize*indent)
	sb.WriteString(fmt.Sprintf("%s<%s>\n",
		i, e.name))
	if len(e.text) > 0 {
		sb.WriteString(strings.Repeat(" ",
			indentSize*(indent+1)))
		sb.WriteString(e.text)
		sb.WriteString("\n")
	}

	for _, el := range e.elements {
		sb.WriteString(el.string(indent + 1))
	}
	sb.WriteString(fmt.Sprintf("%s</%s>\n",
		i, e.name))
	return sb.String()
}

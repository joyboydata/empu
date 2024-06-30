package querystruct

import (
	"bytes"
	"fmt"
	"text/template"
)

type Construct struct {
	Template  string
	Functions map[string]any
}

type Renderer interface {
	Render() (string, error)
}

type QueryModel = any

type ConstructRenderer struct {
	Construct  Construct
	QueryModel any
}

func (c ConstructRenderer) Render() (string, error) {
	return renderTemplate(c.Construct.Template, c.Construct.Functions, c.QueryModel)
}

type ReferenceRenderer interface {
	Render() (string, error)
}

type Reference struct {
	Name      string
	Reference ReferenceRenderer
}

// RenderConstruct wrapper on top of generateTemplate for better organization
func RenderConstruct(constructRenderer ConstructRenderer, references ...Reference) (string, error) {
	if len(references) == 0 {
		return constructRenderer.Render()
	}

	for _, reference := range references {
		// TODO: implement reference render logic
		_ = reference
	}
	// TODO: implement root render logic
	return "", nil
}

// renderTemplate is a function helper to create templatize SQL using GO template. The function return the rendered template as text with trimmed new line and spaces. Return any error from executing the template.
func renderTemplate(tmpl string, funcMap template.FuncMap, input interface{}) (string, error) {
	t := template.Must(template.New("").Funcs(funcMap).Parse(tmpl))
	out := bytes.Buffer{}
	err := t.Execute(&out, input)

	if err != nil {
		return "", fmt.Errorf("fail to render SQL for specified template: %w", err)
	}

	return CleanQuery(out.String()), nil
}

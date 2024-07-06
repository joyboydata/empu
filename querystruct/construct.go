package querystruct

import (
	"bytes"
	"fmt"
	"text/template"
)

// Construct defines a SQL template with associated functions for rendering.
type Construct struct {
	Template  string
	Functions map[string]any
}

// Renderer interface defines a method to render a SQL template.
type Renderer interface {
	Render() (string, error)
}

// QueryModel is an alias for any type, representing the model used in the template rendering.
type QueryModel = any

// ConstructRenderer renders a Construct with a given QueryModel.
type ConstructRenderer struct {
	// The Construct containing the template and functions
	Construct Construct
	// The model data to be used in the template
	QueryModel any
}

// Render generates the rendered SQL string from the Construct and QueryModel.
func (c ConstructRenderer) Render() (string, error) {
	return renderTemplate(c.Construct.Template, c.Construct.Functions, c.QueryModel)
}

// ReferenceRenderer interface defines a method to render a reference template.
type ReferenceRenderer interface {
	Render() (string, error)
}

// Reference holds the name and the reference renderer for a sub-template.
type Reference struct {
	Name      string
	Reference ReferenceRenderer
}

// RenderConstruct renders the main construct with optional references.
func RenderConstruct(constructRenderer ConstructRenderer, references ...Reference) (string, error) {
	if len(references) == 0 {
		return constructRenderer.Render()
	}

	t, err := template.New("").
		Funcs(constructRenderer.Construct.Functions).
		Parse(constructRenderer.Construct.Template)
	if err != nil {
		return "", err
	}

	for _, reference := range references {
		refName := reference.Name
		partial, err := reference.Reference.Render()
		if err != nil {
			return "", fmt.Errorf("fail to render SQL for %s template reference: %w", refName, err)
		}

		_, err = t.New(refName).Parse(partial)
		if err != nil {
			return "", fmt.Errorf("fail to render SQL for %s template reference: %w", refName, err)
		}
	}

	out := bytes.Buffer{}
	err = t.Execute(&out, constructRenderer.QueryModel)
	if err != nil {
		return "", fmt.Errorf("fail to render SQL for root template: %w", err)
	}

	return CleanQuery(out.String()), nil
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

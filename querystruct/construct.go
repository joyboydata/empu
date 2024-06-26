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

// RenderConstruct wrapper on top of generateTemplate for better organization
func RenderConstruct(construct Construct, queryModel interface{}) (string, error) {
	return renderTemplate(construct.Template, construct.Functions, queryModel)
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

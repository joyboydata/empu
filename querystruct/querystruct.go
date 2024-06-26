package querystruct

import (
	"bytes"
	"fmt"
	"text/template"
)

type QueryStructure struct {
	QueryTmpl string
	TmplFns   map[string]any
}

// RenderQueryStructure wrapper on top of generateTemplate for better organization
func RenderQueryStructure(queryStructure QueryStructure, queryModel interface{}) (string, error) {
	return renderTemplate(queryStructure.QueryTmpl, queryStructure.TmplFns, queryModel)
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

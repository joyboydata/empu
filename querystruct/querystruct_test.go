package querystruct

import (
	"testing"
	"text/template"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestRenderTemplate(t *testing.T) {
	tests := []struct {
		name      string
		tmpl      string
		funcMap   template.FuncMap
		input     interface{}
		expected  string
		expectErr bool
	}{
		{
			name:    "should render template when there is no error",
			tmpl:    "SELECT * FROM `TABLE` WHERE name=\"{{.Name}}\"",
			funcMap: template.FuncMap{},
			input: struct {
				Name string
			}{
				Name: "World",
			},
			expected:  "SELECT * FROM `TABLE` WHERE name=\"World\"",
			expectErr: false,
		},
		{
			name:    "should return error when fail to render template",
			tmpl:    "SELECT * FROM `TABLE` WHERE name=\"{{.Name}}\"",
			funcMap: template.FuncMap{},
			input: struct {
				FirstName string
			}{
				FirstName: "World",
			},
			expected:  "",
			expectErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := renderTemplate(tt.tmpl, tt.funcMap, tt.input)
			if tt.expectErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
				assert.Equal(t, tt.expected, result)
			}
		})
	}
}

func TestRenderQueryStructure(t *testing.T) {
	type queryModel struct {
		ID int
	}

	tests := []struct {
		name           string
		queryStructure QueryStructure
		queryModel     interface{}
		expected       string
		expectErr      bool
	}{
		{
			name: "should render query from QueryStructure when there is no error",
			queryStructure: QueryStructure{
				QueryTmpl: "SELECT * FROM users WHERE id = {{.ID}}",
				TmplFns:   template.FuncMap{},
			},
			queryModel: queryModel{ID: 1},
			expected:   "SELECT * FROM users WHERE id = 1",
			expectErr:  false,
		},
		{
			name: "should return error when fail to render QueryStructure",
			queryStructure: QueryStructure{
				QueryTmpl: "SELECT * FROM users WHERE name = {{.Name}}",
				TmplFns:   template.FuncMap{},
			},
			queryModel: queryModel{ID: 1},
			expected:   "",
			expectErr:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := RenderQueryStructure(tt.queryStructure, tt.queryModel)
			if tt.expectErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
				assert.Equal(t, tt.expected, result)
			}
		})
	}
}

package querystruct

import (
	"regexp"
)

var userDefineFunctions = regexp.MustCompile(`\{\{[^{}]*\|\s*([\w]+)|\{\{\s*([\w]+)\s*[^{}]*\}\}`)

func ExtractQueryFunctions(template string) ([]string, error) {
	functions := extractFunctions(template)
	return functions, nil
}

func extractFunctions(tmpl string) []string {
	reserve := map[string]struct{}{
		"if":       {},
		"else":     {},
		"end":      {},
		"range":    {},
		"with":     {},
		"define":   {},
		"block":    {},
		"template": {},
		"null":     {},

		// builtins
		"and":      {},
		"call":     {},
		"html":     {},
		"index":    {},
		"slice":    {},
		"js":       {},
		"len":      {},
		"not":      {},
		"or":       {},
		"print":    {},
		"printf":   {},
		"println":  {},
		"urlquery": {},

		"eq": {}, // ==
		"ge": {}, // >=
		"gt": {}, // >
		"le": {}, // <=
		"lt": {}, // <
		"ne": {}, // !=
	}

	matches := userDefineFunctions.FindAllStringSubmatch(tmpl, -1)

	// Extract function names from the matches
	functions := make(map[string]struct{})
	for _, match := range matches {
		if len(match) > 1 {
			if match[1] != "" {
				if _, found := reserve[match[1]]; !found {
					functions[match[1]] = struct{}{}
				}
			} else if match[2] != "" {
				if _, found := reserve[match[2]]; !found {
					functions[match[2]] = struct{}{}
				}
			}
		}
	}

	var fns []string
	for fName := range functions {
		fns = append(fns, fName)
	}

	return fns
}

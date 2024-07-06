/*
Package querystruct provides a SQL query templating toolkit designed to simplify the creation of data applications.
Empu leverages Go templates to facilitate the construction of dynamic SQL queries, making it easier to manage and
reuse SQL code in a structured and efficient manner.

# Overview

Empu aims to simplify the process of creating complex SQL queries by allowing developers to define templates
for their SQL code. These templates can include placeholders for dynamic values and can be customized with
functions to handle specific rendering logic. The package supports the following key features:

 1. SQL Templates: Define SQL queries with placeholders for dynamic values, making it easy to construct queries
    with varying parameters.

 2. Template Functions: Extend the capabilities of your SQL templates by defining custom functions to handle
    specific logic within the templates.

 3. Query Models: Use any Go data structure as a model to populate the placeholders in your SQL templates,
    allowing for flexible and dynamic query generation.

 4. Reference Templates: Manage complex SQL queries by breaking them down into smaller, reusable templates
    that can be referenced within the main template.

# Package Components

The querystruct package provides several key components to facilitate SQL query templating:

1. Construct: Represents a SQL template with associated functions for rendering.

2. Renderer Interface: Defines the Render method for rendering templates.

3. ConstructRenderer: Implements the Renderer interface, allowing the rendering of a Construct with a given QueryModel.

4. ReferenceRenderer Interface: Defines the Render method for reference templates.

5. Reference: Represents a named reference to a sub-template, enabling modular and reusable SQL query components.

# Usage

To use the querystruct package, define a Construct with your SQL template and any custom functions. Create a
ConstructRenderer with the Construct and a QueryModel. Then, use the RenderConstruct function to generate
the final SQL query, optionally including any Reference templates.

Example:

	package main

	import (
		"fmt"
		"querystruct"
	)

	func main() {
		construct := querystruct.Construct{
			Template: "SELECT * FROM users WHERE name = {{uppercase .Name}}",
			Functions: map[string]interface{}{
				"uppercase": strings.ToUpper,
			},
		}

		queryModel := struct {
			Name int
		}{
			Name: 1,
		}

		renderer := querystruct.ConstructRenderer{
			Construct:  construct,
			QueryModel: queryModel,
		}

		query, err := querystruct.RenderConstruct(renderer)
		if err != nil {
			fmt.Println("Error rendering query:", err)
			return
		}

		fmt.Println("Rendered query:", query)
	}
*/
package querystruct

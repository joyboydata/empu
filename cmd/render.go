package cmd

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/joyboydata/empu/querystruct"
	"github.com/spf13/cobra"
)

var renderCmd = &cobra.Command{
	Use:   "render [path-to-template] [json-data]",
	Short: "Render sql template using given data",
	Args:  cobra.ExactArgs(2),

	Run: func(cmd *cobra.Command, args []string) {
		pathToTemplate := args[0]
		jsonData := args[1]

		template, err := os.ReadFile(pathToTemplate)
		if err != nil {
			fmt.Printf("fail to read template: %v", err)
			os.Exit(1)
		}

		construct := querystruct.Construct{
			Template: string(template),
		}

		var queryModel map[string]any
		err = json.Unmarshal([]byte(jsonData), &queryModel)
		if err != nil {
			fmt.Printf("fail to read query model: %v", err)
			os.Exit(1)
		}

		query, err := querystruct.RenderConstruct(querystruct.ConstructRenderer{
			Construct:  construct,
			QueryModel: queryModel,
		})

		if err != nil {
			fmt.Printf("fail to render template: %v", err)
			os.Exit(1)
		}

		fmt.Print(query)

	},
}

func init() {
	rootCmd.AddCommand(renderCmd)
}

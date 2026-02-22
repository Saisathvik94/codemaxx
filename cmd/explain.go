/*
Copyright © 2026 Saisathvik94
*/
package cmd

import (
	"fmt"

	"github.com/charmbracelet/glamour"
	"github.com/spf13/cobra"

	"github.com/Saisathvik94/codemaxx/internal/ai"
	"github.com/Saisathvik94/codemaxx/internal/files"
	"github.com/Saisathvik94/codemaxx/internal/prompts"
)

func Output(content string) {
	out, _ := glamour.Render(content, "dark")
	fmt.Println(out)
}

var explainUserPrompt string

// explainCmd represents the ask command
var explainCmd = &cobra.Command{
	Use:   "explain [file] [question]",
	Short: "Get explaination of your code directly from the terminal",
	Long: `Sends your question to the selected AI model and returns a direct answer 
inside your terminal.

Supports multiple providers and models.

Examples:
  	codemaxx explain ./home.tsx
	codemaxx explain main.go --prompt "enhance code quality"`,
	RunE: func(cmd *cobra.Command, args []string) error {
		var fullPrompt string
		if len(args) == 0 && explainUserPrompt != "" {
			fullPrompt = fmt.Sprintf("\n %s \nUser Question:\n%s", prompts.ExplainSystemPrompt, explainUserPrompt)
		} else if len(args) == 1 {
			filepath := args[0]

			// allowed extensions
			allowedExts := []string{".ts", ".js", ".py", ".go", ".jsx", ".tsx", ".java", ".c", ".cpp"}

			// vaildates existense of file in directory
			if err := files.ValidateExists(filepath); err != nil {
				return err
			}

			// validates file extension
			if err := files.ValidateExtension(filepath, allowedExts); err != nil {
				return err
			}

			fileContent, err := files.ReadFile(filepath)
			if err != nil {
				return fmt.Errorf("failed to read file %w", err)
			}

			if explainUserPrompt == "" {
				fullPrompt = fmt.Sprintf("\n %s \nFile Content: \n%s", prompts.ExplainSystemPrompt, fileContent)

			} else {
				fullPrompt = fmt.Sprintf("\n %s \nFile Content: \n%s \nUser Instructions : \n %s", prompts.ExplainSystemPrompt, fileContent, explainUserPrompt)
			}
		} else {
			return fmt.Errorf("Provide a file or a --prompt question?")
		}

		resp, err := ai.Generate(cmd.Context(), ai.Request{
			Provider: "",
			Prompt:   fullPrompt,
		})

		if err != nil {
			return fmt.Errorf("Explanation failed: %w", err)
		}

		fmt.Print("\n=========== Code Explanation ===========\n")
		Output(resp.Content)

		return nil

	},
}

func init() {
	rootCmd.AddCommand(explainCmd)
	explainCmd.Flags().StringVarP(&explainUserPrompt, "prompt", "p", "", "Extra Instruction from user")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// askCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// explainCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

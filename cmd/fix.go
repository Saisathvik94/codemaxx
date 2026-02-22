/*
Copyright © 2026 Saisathvik94
*/
package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/Saisathvik94/codemaxx/internal/ai"
	"github.com/Saisathvik94/codemaxx/internal/diff"
	"github.com/Saisathvik94/codemaxx/internal/files"
	"github.com/Saisathvik94/codemaxx/internal/prompts"
	"github.com/spf13/cobra"
)

var userPrompt string

// fixCmd represents the fix command
var fixCmd = &cobra.Command{
	Use:   "fix [file]",
	Short: "Fix and improve a file using AI",
	Long: `Analyzes the specified file and applies AI-powered fixes or improvements.

	Codemaxx updates the file directly with clean code output. 
	Before finalizing, you can review the changes and choose to keep or revert them.

	Examples:
	codemaxx fix ./home.tsx
	codemaxx fix main.go --prompt "enhance code quality"`,
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) == 0 {
			return fmt.Errorf("file path is required")
		}
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

		originalFileContent, err := files.ReadFile(filepath)
		if err != nil {
			return fmt.Errorf("failed to read file %w", err)
		}

		var fullPrompt string

		if userPrompt == "" {
			fullPrompt = fmt.Sprintf("\n %s \nFile Content: \n%s", prompts.SystemPrompt, originalFileContent)

		} else {
			fullPrompt = fmt.Sprintf("\n %s \nFile Content: \n%s \nUser Instructions : \n %s", prompts.SystemPrompt, originalFileContent, userPrompt)
		}

		resp, err := ai.Generate(cmd.Context(), ai.Request{
			Provider: "",
			Prompt:   fullPrompt,
		})

		if err != nil {
			return fmt.Errorf("Code generation failed: %w", err)
		}

		var modifiedFileContent = resp.Content

		err = os.WriteFile(filepath, []byte(modifiedFileContent), 0644)
		if err != nil {
			return fmt.Errorf("failed to write AI changes: %w", err)
		}

		diff.ShowDiff(originalFileContent, modifiedFileContent)

		for {
			fmt.Print("\nKeep changes? (y/n): ")

			var input string
			fmt.Scanln(&input)

			input = strings.TrimSpace(strings.ToLower(input))

			switch input {
			case "y", "yes":
				fmt.Println("✔ Changes Applied")
				return nil

			case "n", "no":
				if err := os.WriteFile(filepath, []byte(originalFileContent), 0644); err != nil {
					return fmt.Errorf("failed to restore original file: %w", err)
				}
				fmt.Println("✔ Changes Reverted")
				return nil

			default:
				fmt.Println("Please enter 'y' or 'n'")
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(fixCmd)
	fixCmd.Flags().StringVarP(&userPrompt, "prompt", "p", "", "Extra Instruction from user")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// fixCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// fixCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"

	"github.com/Saisathvik94/codemaxx/internal/ai"
	"github.com/Saisathvik94/codemaxx/internal/prompts"
	"github.com/spf13/cobra"
)

// reviewCmd represents the review command
var reviewCmd = &cobra.Command{
	Use:   "review",
	Short: "Review staged git changes using AI",
	Long: `Review analyzes your currently staged git changes using an AI model.

	It runs:

	git diff --cached

	and sends the diff to the selected AI provider for analysis.

	The AI will:
	- Detect potential bugs
	- Suggest improvements
	- Identify code smells
	- Recommend refactoring
	- Highlight security or performance concerns

	Only staged changes are reviewed, so make sure you run:

	git add <files>

	before executing this command.

	Example:
	codemaxx review

	This command does NOT modify your files. It only prints an AI-generated review in your terminal.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		
		// gitcommand
		gitCmd := exec.Command("git", "diff", "--cached")

		var out bytes.Buffer
		gitCmd.Stdout = &out

		err:= gitCmd.Run() 

		if err != nil {
			fmt.Println("Failed to get staged changes.")
			fmt.Println("Make sure you are inside a git repository.")

			os.Exit(1)
		}

		diff:= out.String()

		if diff == "" {
			fmt.Println("No staged changes found.")
			return nil
		}

		if len(diff)>5000 {
			fmt.Println("Staged diff is too large to review at once.")
			fmt.Println("Please stage smaller changes and review them incrementally.")
			return nil
		}

		fmt.Println("Reviewing staged changes....")

		var prompt = fmt.Sprintf("\n %s \nDiffs: \n%s", prompts.ReviewPrompt, diff)

		// send staged changes to selected model
		resp, err := ai.Generate(cmd.Context(), ai.Request{
			Provider: "",
			Prompt:   prompt,
		})
		if err != nil {
			return fmt.Errorf("AI review failed: %w", err)
		}

		fmt.Print("\n=========== AI code review ===========\n")
		Output(resp.Content)

		return nil

	},
}

func init() {
	rootCmd.AddCommand(reviewCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// reviewCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// reviewCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

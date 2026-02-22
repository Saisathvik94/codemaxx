/*
Copyright © 2026 Saisathvik94

*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"
	_ "github.com/Saisathvik94/codemaxx/internal/models/perplexity"
	_ "github.com/Saisathvik94/codemaxx/internal/models/openai"
	_ "github.com/Saisathvik94/codemaxx/internal/models/gemini"
	_ "github.com/Saisathvik94/codemaxx/internal/models/anthropic"
)



// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "codemaxx",
	Short: "AI-powered terminal assistant for editing, fixing, and generating code",
	Long: `Codemaxx is a developer-focused CLI tool that brings multiple AI models 
directly into your terminal workflow.

It allows you to:
- Fix and refactor files
- Generate new code
- Improve existing codebases
- Work with multiple LLM providers (OpenAI, Anthropic, Gemini, etc.)
- Review and accept or revert changes safely

Codemaxx updates files directly with clean code output and lets you 
approve or discard changes before finalizing them.

Built with Go and Cobra for speed, simplicity, and extensibility.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.codemaxx.git.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}



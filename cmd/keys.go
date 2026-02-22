/*
Copyright © 2026 Saisathvik94
*/
package cmd

import (
	"fmt"

	"github.com/Saisathvik94/codemaxx/internal/models"
	"github.com/Saisathvik94/codemaxx/internal/ui"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/spf13/cobra"
)

// keysCmd represents the keys command
var keysCmd = &cobra.Command{
	Use:   "keys",
	Short: "Manage API keys for AI providers",
	Long: `Configure and manage API keys for supported AI providers.

	Codemaxx uses API keys to connect to different LLM services.
	Keys can be set through environment variables or configured interactively.

	Examples:
	codemaxx keys
	codemaxx keys openai
	codemaxx keys anthropic`,
	RunE: func(cmd *cobra.Command, args []string) error {
		providers := models.ListProviders()

		if len(providers) == 0 {
			return fmt.Errorf("no provider registered")
		}

		p := tea.NewProgram(ui.SetNewKey(providers))

		_, err := p.Run()

		return err
	},
}

func init() {
	rootCmd.AddCommand(keysCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// keysCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// keysCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

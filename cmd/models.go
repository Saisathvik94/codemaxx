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

// modelsCmd represents the models command
var modelsCmd = &cobra.Command{
	Use:   "models",
	Short: "List available AI models",
	Long: `Displays all supported AI providers and available models.

	Helps you choose which model to use for your commands.

	Examples:
	codemaxx models`,

	RunE: func(cmd *cobra.Command, args []string) error {
		providers := models.ListProviders()

		if len(providers) == 0 {
			return fmt.Errorf("no provider registered")
		}

		p := tea.NewProgram(ui.NewModelSelector(providers))

		_, err := p.Run()

		return err

	},
}

func init() {
	rootCmd.AddCommand(modelsCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// modelsCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// modelsCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

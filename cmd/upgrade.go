package cmd

import (
	"fmt"
	"os"

	"github.com/blang/semver"
	"github.com/rhysd/go-github-selfupdate/selfupdate"
	"github.com/spf13/cobra"
)

const repo = "Saisathvik94/codemaxx"

var force bool

var upgradeCmd = &cobra.Command{
	Use:   "upgrade",
	Short: "Upgrade CodeMaxx to the latest released version",
	Long: `Upgrade checks GitHub for the latest CodeMaxx release
	and automatically downloads and replaces the current
	binary if a newer version is available.

	Example:
	codemaxx upgrade`,
	Run: func(cmd *cobra.Command, args []string) {

		// block dev builds
		if Version == "" || Version == "dev" {
			fmt.Println("Cannot upgrade a development build.")
			fmt.Println("Please install CodeMaxx from a GitHub release.")
			return
		}

		current, err := semver.ParseTolerant(Version)
		if err != nil {
			fmt.Println("Invalid version format:", err)
			os.Exit(1)
		}

		fmt.Printf("Current version: %s\n", current)
		fmt.Println("Checking for updates...")

		latest, err := selfupdate.UpdateSelf(current, repo)
		if err != nil {
			fmt.Println("Update failed:", err)
			fmt.Println("Please check your internet connection.")
			os.Exit(1)
		}

		// If already latest
		if latest.Version.Equals(current) && !force {
			fmt.Println("You are already using the latest version.")
			return
		}

		fmt.Printf("Successfully updated to version: %s\n", latest.Version)
		fmt.Println("Please restart CodeMaxx.")
	},
}

func init() {
	rootCmd.AddCommand(upgradeCmd)
}

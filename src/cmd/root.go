package cmd

import (
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Version: "0.0.1",
	Use:     "nsreaper",
	Short:   "a cli for clean previews namespace",
	Long: `a cli for clean previews namespace after a ttl, by default 10 day:
namespace clean --dry-run true --ttl 10`,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		panic(err)
	}
}


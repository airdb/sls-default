package cmd

import (
	"fmt"
	"os"

	"github.com/airdb/scf-airdb/internal/version"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:     "cli",
	Short:   "Cli tool",
	Long:    "Cli Tool",
	Version: version.Version,
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
	},
}

func initRoot() {
	initDomainCmd()
	initShell()
	initDailyEnglish()
}

func Execute() {
	initRoot()

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

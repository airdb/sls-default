package cmd

import (
	"fmt"
	"log"

	"github.com/airdb/scf-airdb/model/vo"
	"github.com/airdb/scf-airdb/pkg/airsdk"
	"github.com/spf13/cobra"
)

// shellCmd represents the shell command
var shellCmd = &cobra.Command{
	Use:     "shell",
	Short:   "A brief description of your command",
	Example: createShellExample,
	Run: func(cmd *cobra.Command, args []string) {
		ListShell()
	},
}

// shellCmd represents the shell command
var listShellCmd = &cobra.Command{
	Use:     "list",
	Short:   "A brief description of your command",
	Example: createShellExample,
	Aliases: []string{"ls"},
	Run: func(cmd *cobra.Command, args []string) {
		ListShell()
	},
}

// shellCmd represents the shell command
var addShellCmd = &cobra.Command{
	Use:     "add",
	Short:   "A brief description of your command",
	Example: createShellExample,
	Run: func(cmd *cobra.Command, args []string) {
		createShell()
	},
}

func initShell() {
	rootCmd.AddCommand(shellCmd)

	shellCmd.AddCommand(listShellCmd)
	shellCmd.AddCommand(addShellCmd)

	listShellCmd.PersistentFlags().StringVarP(&createShellFlags.Command, "command", "c", "", "command")

	addShellCmd.PersistentFlags().StringVarP(&createShellFlags.Command, "command", "c", "", "command")
	addShellCmd.PersistentFlags().StringVarP(&createShellFlags.Shell, "shell", "s", "", "shell")
	addShellCmd.PersistentFlags().StringVarP(&createShellFlags.Comment, "comment", "", "", "comment")
	addShellCmd.PersistentFlags().StringVarP(&createShellFlags.Ref, "ref", "r", "", "ref")
	addShellCmd.PersistentFlags().StringVarP(&createShellFlags.Tags, "tag", "t", "", "tag")
}

var listShellFlags vo.ListLinuxShellReq

func ListShell() {
	resp, err := airsdk.ListLinuxShell(&listShellFlags)
	if err != nil {
		log.Fatalln(err)
	}

	for _, shell := range resp.Shells {
		fmt.Printf("%-16s | %s | %s\n", shell.Command, shell.Shell, shell.Comment)
	}
}

var createShellFlags vo.CreateLinuxShellReq

func createShell() {
	resp, err := airsdk.CreateLinuxShell(&createShellFlags)
	if err != nil {
		log.Fatalln(err)
	}

	log.Println(resp)
}

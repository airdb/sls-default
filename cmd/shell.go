/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
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

	shellCmd.AddCommand(addShellCmd)

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
		fmt.Println(*shell)
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

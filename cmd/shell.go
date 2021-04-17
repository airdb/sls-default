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
	"log"

	"github.com/airdb/scf-airdb/model/vo"
	"github.com/airdb/scf-airdb/pkg/airsdk"
	"github.com/spf13/cobra"
)

// shellCmd represents the shell command
var shellCmd = &cobra.Command{
	Use:   "shell",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		createShell()
	},
}

func initShell() {
	rootCmd.AddCommand(shellCmd)
}

func createShell() {
	resp, err := airsdk.CreateLinuxShell(&vo.CreateLinuxShellReq{
		Command: "mc",
		Comment: "",
		Shell:   "wget https://dl.min.io/client/mc/release/linux-amd64/archive/mc.RELEASE.2019-10-02T19-41-02Z",
		Ref:     "",
		Tags:    "minio",
	})
	if err != nil {
		log.Fatalln(err)
	}

	log.Println(resp)
}

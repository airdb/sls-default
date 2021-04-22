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
var dailyEnglishCmd = &cobra.Command{
	Use:     "english",
	Short:   "A brief description of your command",
	Example: createDailyEnglishExample,
	Aliases: []string{"en"},
	Run: func(cmd *cobra.Command, args []string) {
		ListDailyEnglish()
	},
}

// shellCmd represents the shell command
var addDailyEnglishCmd = &cobra.Command{
	Use:     "add",
	Short:   "A brief description of your command",
	Example: createDailyEnglishExample,
	Run: func(cmd *cobra.Command, args []string) {
		createDailyEnglish()
	},
}

func initDailyEnglish() {
	rootCmd.AddCommand(dailyEnglishCmd)

	dailyEnglishCmd.AddCommand(addDailyEnglishCmd)

	addDailyEnglishCmd.PersistentFlags().StringVarP(&createDailyEnglishFlags.Cn, "chinese", "c", "", "command")
	addDailyEnglishCmd.PersistentFlags().StringVarP(&createDailyEnglishFlags.En, "english", "e", "", "comment")
	addDailyEnglishCmd.PersistentFlags().StringVarP(&createDailyEnglishFlags.Category, "category", "c", "", "category")
	addDailyEnglishCmd.PersistentFlags().StringVarP(&createDailyEnglishFlags.Tags, "tag", "t", "", "tag")
}

var listDailyEnglishFlags vo.ListDailyEnglishReq

func ListDailyEnglish() {
	resp, err := airsdk.ListDailyEnglish(&listDailyEnglishFlags)
	if err != nil {
		log.Fatalln(err)
	}

	for _, shell := range resp.En {
		fmt.Println(*shell)
	}
}

var createDailyEnglishFlags vo.CreateDailyEnglishReq

func createDailyEnglish() {
	resp, err := airsdk.CreateDailyEnglish(&createDailyEnglishFlags)
	if err != nil {
		log.Fatalln(err)
	}

	log.Println(resp)
}

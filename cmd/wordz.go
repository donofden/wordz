/*
Copyright © 2019 NAME HERE <EMAIL ADDRESS>

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
	"os"

	"github.com/spf13/cobra"
	"github.com/donofden/wordz/pkg/search"
)

var wordCmd = &cobra.Command{
	Use:   "find",
	Short: "A brief description of your command",
	Long:  `A longer description of count command.`,
	Run: func(cmd *cobra.Command, args []string) {

		if len(args) > 0 {
			fmt.Println("Please wait while we search for the meaning...")
		} else {
			fmt.Println("Please Specify a word to find meaning.")
			os.Exit(1)
		}
		wordz.SearchWord()
		fmt.Println(args)
	},
}

func init() {
	rootCmd.AddCommand(wordCmd)
}

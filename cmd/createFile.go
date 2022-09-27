/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

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

	"os"

	"github.com/spf13/cobra"
)

// createFileCmd represents the createFile command
var createFileCmd = &cobra.Command{
	Use:   "createFile",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("createFile called")
		_, err := os.Create(args[0])
		if err != nil {
			log.Fatal(err)
			os.Exit(1)
		} else {
			fmt.Println("file created with name -> ", args[0])
		}
		// message := []byte("Hello, Gophers!")
		// error := ioutil.WriteFile("data.txt", message, 0644)
		// if error != nil {
		// 	log.Fatal(error)
		// }
	},
}

func init() {
	rootCmd.AddCommand(createFileCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// createFileCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// createFileCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

/*
Copyright © 2022 Parav Kaushal <paravkaushal.kv1@gmail.com>

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

	homedir "github.com/mitchellh/go-homedir"
	"github.com/paravkaushal/go-journal/pkg"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "go-journal",
	Short: "A personal journal log with user management",

	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Welcome to Journal Application")
		WelcomeFunction()
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.go-journal.yaml)")
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		// Search config in home directory with name ".go-journal" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigName(".go-journal")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}
// Welcome function notes user's intention
func WelcomeFunction() {
	fmt.Println("Please Login or Signup:")
	fmt.Println("1. LogIn\n2. SignUp")
	var input string
	fmt.Scanln(&input)

	if input == "1" {
		pkg.LogIn()
	} else if input == "2" {
		pkg.SignUp()
	} else {
		fmt.Println("Please enter the correct option")
		WelcomeFunction()
	}
}

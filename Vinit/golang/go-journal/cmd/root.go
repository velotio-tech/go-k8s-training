// Package cmd /*
package cmd

import (
	"bufio"
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"go-journal/constants"
	"go-journal/users"
	"os"
	"time"
)

//var JournalBuffer = make([]string, 50, 50)
var cfgFile string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "go-journal",
	Short: "A brief description of your application",
	Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	Run: func(cmd *cobra.Command, args []string) {
		makeUserDb()
		fmt.Println("Welcome to the Journal App")
		user := new(users.User)
		GetStarted(user)
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}

func init() {
	cobra.OnInitialize(initConfig)

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.go-journal.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := os.UserHomeDir()
		cobra.CheckErr(err)

		// Search config in home directory with name ".go-journal" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigType("yaml")
		viper.SetConfigName(".go-journal")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Fprintln(os.Stderr, "Using config file:", viper.ConfigFileUsed())
	}
}

func isError(err error) bool {
	if err != nil {
		fmt.Println(err.Error())
	}

	return err != nil
}

func makeUserDb(){
	var _, err = os.Stat(constants.DB_LOCATION+constants.USER_DB)
	if os.IsNotExist(err) {
		UsersDatabase, e := os.Create(constants.DB_LOCATION+constants.USER_DB)
		if isError(e) {
			return
		}
		defer UsersDatabase.Close()
	}
}

func GetStarted(user *users.User){
	var inp int
	for {
		if user.IsAuthenticated {
			fmt.Printf("User Loggedin as %s \n", user.Username)
		}
		fmt.Println("1. Signup \n 2. Login \n 3. List \n 4. Add Entry \n 5. exit")
		_, err := fmt.Scanf("%d", &inp)
		if err != nil {
			fmt.Println("Error taking the input, Try again!")
			continue
		}
		switch inp {
		case 1:
		//	signup process here
			fmt.Println("Enter the username")
			var username string
			fmt.Scanln(&username)
			fmt.Println("Enter the password")
			var pass string
			fmt.Scanln(&pass)
			user = createJournal(username, pass)
		case 2:
			fmt.Println("Enter the username")
			var username string
			fmt.Scanln(&username)
			fmt.Println("Enter the password")
			var pass string
			fmt.Scanln(&pass)
			user = AuthUser(username, pass)
		case 3:
			if !user.IsAuthenticated {
				fmt.Println("User Not Authenticated login to view the journal")
				continue
			} else{
				user.ReadJournal()
			}
		case 4:
			if !user.IsAuthenticated {
				fmt.Println("User Not Authenticated login to view the journal")
				continue
			} else {
				fmt.Println("Add the entry to the journal")
				var entry string
				reader := bufio.NewReader(os.Stdin)
				read, _, err := reader.ReadLine()
				if err != nil {
					fmt.Println(err)
				}
				entry = string(read)
				entry = time.Now().Format("02-01-06 15:01") + " " + entry
				//JournalBuffer = append(JournalBuffer, entry)
				user.WriteJournal(entry)
			}
		case 5:
			os.Exit(0)
		default:
			continue
		}
	}
}

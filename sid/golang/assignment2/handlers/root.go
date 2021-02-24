package handlers

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

func Root() func(*cobra.Command, []string) {
	return func(cmd *cobra.Command, args []string) {
		fmt.Print("Enter:\n1 - To login as an exiting user\n2 - To register as a new user\n")

		reader := bufio.NewReader(os.Stdin)
		option, _ := reader.ReadString('\n')
		option = strings.TrimSpace(option)

		if option == "1" {
			Login()(cmd, args)
		} else {
			Signup()(cmd, args)
		}
	}
}

package cmd

import (
	"bufio"
	"fmt"
	"insomnia/cmd/endpoint"
	"insomnia/cmd/workspace"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "insomnia",
	Short: "Insomnia - A CLI tool for monitoring uptime and endpoints",
	Long: `Insomnia is a lightweight CLI tool designed for uptime monitoring.
It helps you manage workspaces and monitor endpoints effortlessly.
Easily track, create, and manage server and service availability
with an intuitive command-line interface.`,
	Run: func(cmd *cobra.Command, args []string) {

	},
}

// Function to start the interactive shell
func StartShell() {

	fmt.Println("Welcome to Insomnia CLI! Type 'help' for a list of commands.")

	reader := bufio.NewReader(os.Stdin)
	// Get input command for the shell
	for {

		fmt.Print("\n\033[1;32minsomnia 💤 > \033[0m")

		input, err := reader.ReadString('\n')
		input = strings.TrimSpace(input)
		if err != nil {
			fmt.Printf("Error getting input: %v\n", err)
			continue
		}

		// Check for exit or quit
		if (input == "exit") || (input == "quit") {
			fmt.Println("Exiting Insomnia CLI")
			break
		}

		// Check for help
		if input == "help" {
			rootCmd.Help()
			continue
		}

		// Process user input
		args := strings.Fields(input)
		if len(args) > 0 {
			// check for invalidcommands
			_, _, err := rootCmd.Find(args)
			if err != nil {
				fmt.Printf("Invalid command : %v\n", err)
				continue
			}

			// Set the arguments for rootCmd to execute
			rootCmd.SetArgs(args)

			// handle error on execution
			err = rootCmd.Execute()
			if err != nil {
				fmt.Printf("Error : %v\n", err)
			}
		}
	}
}

func Execute() {

	if len(os.Args) > 1 {
		if err := rootCmd.Execute(); err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	} else {
		StartShell()
	}
}

func init() {
	rootCmd.AddCommand(workspace.WorkspaceCmd)
	rootCmd.AddCommand(endpoint.EndpointCmd)
}

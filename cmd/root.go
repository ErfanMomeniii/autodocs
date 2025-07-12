package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var (
	modelName   string
	apiKey      string
	projectPath string
)

// rootCmd represents the base command when called without any subcommands.
var rootCmd = &cobra.Command{
	Use:   "autodocs",
	Short: "AutoDocs is a very fast go docs generator by using AI",
	Long: `AutoDocs is an intelligent documentation generator for Go projects powered by AI.
It analyzes your codebase to generate clear, accurate, and readable documentation for functions, structs, and methods.

Designed for speed and simplicity, AutoDocs helps you:
  - Save time writing docs
  - Improve code understanding
  - Keep documentation in sync with your code

Ideal for both small libraries and large projects.`,
}

func init() {
	// init initializes the command line flags.
	rootCmd.PersistentFlags().StringVarP(&modelName, "model", "m", "gpt-4o", "model name")
	rootCmd.PersistentFlags().StringVarP(&apiKey, "apikey", "k", "", "api key")
	rootCmd.PersistentFlags().StringVarP(&projectPath, "project", "p", "./", "project path")

	rootCmd.AddCommand(runCmd)
}

// Execute runs the root command and exits the program if an error occurs.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
package cmd

import (
	"fmt"
	models "github.com/erfanmomeniii/autodocs/ai-models"
	"github.com/spf13/cobra"
	"os"
)

var (
	modelName   string
	apiKey      string
	projectPath string
)

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
	Run: func(cmd *cobra.Command, args []string) {
		model, err := models.NewAIModel(
			models.WithApiKey(apiKey),
			models.WithName(modelName))
		if err != nil {
			println(err.Error())
			return
		}
		if err := model.Generator.Generate(projectPath); err != nil {
			println(err.Error())
			return
		}
		return
	},
}

func init() {
	rootCmd.PersistentFlags().StringVarP(&modelName, "model", "m", "gpt-4o", "model name")
	rootCmd.PersistentFlags().StringVarP(&apiKey, "apikey", "k", "", "api key")
	rootCmd.PersistentFlags().StringVarP(&projectPath, "project", "p", "./", "project path")
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

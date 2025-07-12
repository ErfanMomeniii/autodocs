package cmd

import (
	models "github.com/erfanmomeniii/autodocs/ai-models"
	"github.com/spf13/cobra"
)

// runCmd represents the run command which initializes and runs the AI model generator.
var runCmd = &cobra.Command{
	Use: "run",
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
package cli_commands

import (
	"fmt"

	"example.com/todo-app/database"
	"example.com/todo-app/models"
	"github.com/AlecAivazis/survey/v2"
	"github.com/spf13/cobra"
)

var questions = []*survey.Question{
	{
		Name:     "Name",
		Prompt:   &survey.Input{Message: "Give your task a name:"},
		Validate: survey.Required,
	},
	{
		Name: "Description",
		Prompt: &survey.Input{
			Message: "Describe your task:",
			Default: "No description",
		},
		Validate: survey.Required,
	},
	{
		Name:   "IsDone",
		Prompt: &survey.Select{Message: "Is it done already ?", Options: []string{"Yes", "No"}},
	},
}

var createTodoCmd = &cobra.Command{
	Use:   "create:todo",
	Short: "Add a todo",

	Run: func(cmd *cobra.Command, args []string) {
		answers := struct {
			Name        string
			Description string
			IsDone      string
		}{}
		err := survey.Ask(questions, &answers)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		todo := models.Todo{
			Title:       answers.Name,
			Description: answers.Description,
			IsDone:      answers.IsDone == "Yes",
		}
		// marhsal print the struct in json format
		err = database.CreateTodo(&todo)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		fmt.Println("Todo created successfully")

	},
}

func init() {
	rootCmd.AddCommand(createTodoCmd)
}

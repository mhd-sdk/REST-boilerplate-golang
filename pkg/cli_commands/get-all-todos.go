package cli_commands

import (
	"fmt"

	"example.com/todo-app/database"
	"github.com/alexeyco/simpletable"
	"github.com/spf13/cobra"
)

var getAllCmd = &cobra.Command{
	Use:   "get:todos",
	Short: "Show all todosn in the database",

	Run: func(cmd *cobra.Command, args []string) {
		todos := database.GetTodos()
		table := simpletable.New()
		table.Header = &simpletable.Header{
			Cells: []*simpletable.Cell{
				{Align: simpletable.AlignCenter, Text: "#"},
				{Align: simpletable.AlignCenter, Text: "Title"},
				{Align: simpletable.AlignCenter, Text: "Description"},
				{Align: simpletable.AlignCenter, Text: "IsDone"},
			},
		}
		for _, row := range todos {
			r := []*simpletable.Cell{
				{Align: simpletable.AlignRight, Text: fmt.Sprintf("%d", row.ID)},
				{Text: row.Title},
				{Text: row.Description},
				{Align: simpletable.AlignCenter, Text: fmt.Sprintf("%t", row.IsDone)},
			}

			table.Body.Cells = append(table.Body.Cells, r)
		}
		table.SetStyle(simpletable.StyleDefault)
		fmt.Println(table.String())
	},
}

func init() {
	rootCmd.AddCommand(getAllCmd)
}

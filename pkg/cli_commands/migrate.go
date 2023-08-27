package cli_commands

import (
	"fmt"

	"example.com/todo-app/database"
	"github.com/spf13/cobra"
)

var migrateCmd = &cobra.Command{
	Use:   "migrate",
	Short: "Generate database",
	Long:  `Generate database schema from struct models`,
	Run: func(cmd *cobra.Command, args []string) {
		database.Migrate()
		fmt.Println("Database migration done.")
	},
}

func init() {
	rootCmd.AddCommand(migrateCmd)
}

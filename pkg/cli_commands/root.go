package cli_commands

import (
	"fmt"
	"os"
	"strings"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "todo-cli",
	Short: "todo-cli is a very simple cli for managing the todo-app backend",
	Long: `This cli allow you to manage migrations, manual inserts, and many more
Made by Mehdi Seddik.
Project made for the purpose of learning golang.`,
}

func Execute() {
	// change color of headings in help
	cobra.AddTemplateFunc("StyleHeading", color.New(color.BgMagenta, color.FgWhite).SprintFunc())
	//ANSI Shadow font
	banner := ` 
████████╗ ██████╗ ██████╗  ██████╗        ██████╗██╗     ██╗
╚══██╔══╝██╔═══██╗██╔══██╗██╔═══██╗      ██╔════╝██║     ██║
   ██║   ██║   ██║██║  ██║██║   ██║█████╗██║     ██║     ██║
   ██║   ██║   ██║██║  ██║██║   ██║╚════╝██║     ██║     ██║
   ██║   ╚██████╔╝██████╔╝╚██████╔╝      ╚██████╗███████╗██║
   ╚═╝    ╚═════╝ ╚═════╝  ╚═════╝        ╚═════╝╚══════╝╚═╝
`
	if len(os.Args) == 1 || os.Args[1] == "help" {
		fmt.Println(color.New(color.FgMagenta).Sprint(banner))
	}
	usageTemplate := rootCmd.UsageTemplate()
	usageTemplate = strings.NewReplacer(
		`Usage:`, `{{StyleHeading "Usage:"}}`,
		`Aliases:`, `{{StyleHeading "Aliases:"}}`,
		`Available Commands:`, `{{StyleHeading "Available Commands:"}}`,
		`Flags:`, `{{StyleHeading "Flags:"}}`,
	).Replace(usageTemplate)
	rootCmd.SetUsageTemplate(usageTemplate)
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "Whoops. There was an error while executing your CLI '%s'", err)
		os.Exit(1)
	}
}

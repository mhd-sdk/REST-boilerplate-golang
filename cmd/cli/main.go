package main

import (
	"example.com/todo-app/database"
	"example.com/todo-app/pkg/cli_commands"
	"github.com/spf13/viper"
)

func main() {
	viper.SetConfigFile("./../../.env")
	viper.ReadInConfig()
	database.Init()
	cli_commands.Execute()
}

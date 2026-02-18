package cmd

import (
	"flag"
	"strings"

	"github.com/knyazev-ro/vulcan-migrate/api"
)

func Handle(args []string) {

	command := args[1]

	if strings.Split(command, ":")[0] != "pertdb" {
		return
	}

	if command == "pertdb:help" {
		api.GetHelp()
		return
	}

	api.Init()

	flagSet := flag.NewFlagSet("args", flag.ContinueOnError)
	flagSet.Parse(args[2:])
	argsArr := flagSet.Args()

	switch command {
	case "pertdb:create-table":
		api.CreateMigration(argsArr)
	case "pertdb:alter-table":
		api.AlterMigration(argsArr)
	case "pertdb:run":
		api.Migrate(argsArr)
	case "pertdb:rollback":
		api.Rollback(argsArr)
	default:
		println("Error. Unknown command ", command)
		return
	}

}

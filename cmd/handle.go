package cmd

import (
	"flag"

	"github.com/knyazev-ro/perturabo/api"
)

func Handle(args []string) {

	api.Init()

	command := args[1]
	if command == "help" {
		// GetHelp()
		return
	}

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

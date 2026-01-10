package api

import "fmt"

func GetHelp() {
	fmt.Println(`
gerard-framework: perturabo

Usage:
  go run main.go <command> [options]

Available commands:

  pertdb:create-table <migration-name> <table-name>
      Create a new table migration.

  pertdb:alter-table <migration-name> <table-name>
      Alter an existing table with a new migration.

  pertdb:run
      Apply all pending migrations.

  pertdb:rollback
      Roll back the last applied migration.
`)
}

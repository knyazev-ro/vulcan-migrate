package api

import "perturabo/migrate"

func Rollback(args []string) {
	migrate.Down()
}

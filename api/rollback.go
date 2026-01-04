package api

import "github.com/knyazev-ro/perturabo/migrate"

func Rollback(args []string) {
	migrate.Down()
}

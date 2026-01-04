package api

import "github.com/knyazev-ro/perturabo/migrate"

func Migrate(args []string) {
	migrate.Up()
}

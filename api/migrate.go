package api

import "github.com/knyazev-ro/vulcan-migrate/migrate"

func Migrate(args []string) {
	migrate.Up()
}

package api

import "github.com/knyazev-ro/vulcan-migrate/migrate"

func Rollback(args []string) {
	migrate.Down()
}

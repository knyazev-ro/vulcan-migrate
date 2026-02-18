package main

import (
	"os"

	"github.com/knyazev-ro/vulcan-migrate/cmd"
	_ "github.com/knyazev-ro/vulcan-migrate/migrations"
)

func main() {
	cmd.Handle(os.Args)
}

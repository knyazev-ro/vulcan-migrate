package main

import (
	"os"

	"github.com/knyazev-ro/perturabo/cmd"
	_ "github.com/knyazev-ro/perturabo/migrations"
)

func main() {
	cmd.Handle(os.Args)
}

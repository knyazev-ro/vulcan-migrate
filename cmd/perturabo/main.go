package main

import (
	"os"

	"github.com/knyazev-ro/perturabo/cmd"
)

func main() {
	cmd.Handle(os.Args)
}

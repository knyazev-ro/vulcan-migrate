package migrate

import (
	"fmt"

	"github.com/knyazev-ro/perturabo/registry"
)

func Up() {
	arr, err := Get(registry.Action.Up)
	if err != nil {
		fmt.Println("Migration stopped. Error.", err.Error())
	}

	fmt.Println(arr)
}

func Down() {
	arr, err := Get(registry.Action.Down)
	if err != nil {
		fmt.Println("Migration stopped. Error.", err.Error())
	}

	fmt.Println(arr)
}

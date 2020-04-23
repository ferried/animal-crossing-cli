package main

import (
	"animal-crossing-cli/pkg/insect"
	"fmt"
)

func init() {
}

func main() {
	insect.Store()
	fmt.Println(*insect.Insects)
}

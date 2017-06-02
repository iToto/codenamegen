package main

import (
	"fmt"

	"github.com/iToto/codenamegen/data"
)

func main() {
	randomColour := data.RandomColour()
	randomAnimal := data.RandomAnimal()
	codename := randomColour.String + randomAnimal.String

	fmt.Printf("%s\n", codename)
}

package main

import (
	"fmt"
	"marvel-go/models"
	"marvel-go/services"
)

func main() {
	fmt.Print("Hello world!")
	fmt.Print(services.Add(1, 2))
	fmt.Print(models.CharacterDataWrapper{Status: "true"})
}

package main

import (
	"fmt"

	"github.com/badoux/checkmail"
)

func main() {

	err := checkmail.ValidateFormat("invalid")
	fmt.Println("writing from the main")
	fmt.Println(err)
}

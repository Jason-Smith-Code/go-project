package main

import (
	"fmt"

	"github.com/Jason-Smith-Code/go-project/models"
)

func main() {
	u := models.User{
		ID: 2,
		FirstName: "jason",
		LastName: "smith",
	}

	fmt.Println(u)

}
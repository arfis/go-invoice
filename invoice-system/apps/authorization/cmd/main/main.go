package main

import (
	"fmt"
	database "github.com/arfis/go-invoice/authorization/pkg/db"
)

// this is still nil
var (
	db *database.Database
)

func main() {
	// this automatically does an object an assigns an address
	fmt.Println("STARTING APP")
}

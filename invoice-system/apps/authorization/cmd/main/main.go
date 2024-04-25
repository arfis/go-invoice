package main

import (
	"fmt"
	database "github.com/arfis/go-invoice/authorization/pkg/db"
	"net/http"
	"runtime"
)

// this is still nil
var (
	db *database.Database
)

func main() {
	// this automatically does an object an assigns an address
	db2 := &database.Database{}
	db2.Test()
	//db.Test3()
	returnValue := db2.Test2()
	fmt.Printf("After setting non pointer value IS: %d and the returned %d \n", db2.Val, returnValue.Val)
	db2.Test3()

	//fmt.Printf("The value IS: %d", db.Val)
	http.HandleFunc("/", handler)
	//db.CreateConnection()
	//db.TrySelect()
	//db.Test()
	fmt.Println("Server starting on port 8080...")
	err := http.ListenAndServe(":8080", nil) // This will block and keep the program running
	if err != nil {
		fmt.Println("Error starting server:", err)
		return
	}

}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Println(runtime.NumCPU())
	w.Write([]byte("Hello World from docker and has CPU:" + string(runtime.NumCPU())))
}

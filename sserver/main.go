// Simple static file server
package main

import (
	"fmt"
	"net/http"
	"os"
)

const usage = `
Usage:
./sserver [enter] (now go to a web browser and type localhost:8080 [enter])

`

func main() {
	// get current working director path and store in dir variable
	dir, _ := os.Getwd()

	// Print defult directory to terminal after command run
	fmt.Println("Default server directory:", dir)

	// Print usage instruction
	fmt.Print(usage)

	// Start a static webserver
	http.ListenAndServe(":8080", http.FileServer(http.Dir(dir)))
}

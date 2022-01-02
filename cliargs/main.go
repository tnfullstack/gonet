// Accepting online agrs
package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"
)

const usage = `
Usage:
./cliargs -port=8080 -path=test (port could be any number)
	now go to a web browser and type localhost:8080 [enter]

./cliargs [enter] (default port will be 3000)
	now type in localhost:3000 [enter]

`

func main() {

	var dir string
	static := "/test"

	path := flag.String("path", "", "File server directory")
	port := flag.String("port", "3000", "File server listener port number")
	flag.Parse()

	if *path == "" {
		dir, _ = os.Getwd()
		dir += static
		fmt.Println(dir)
	} else {
		dir = *path
	}

	fmt.Print(usage)

	http.ListenAndServe(":"+*port, http.FileServer(http.Dir(dir)))
}

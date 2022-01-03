package main 

import (
	"log"
	"godms"
)

func main() {

	srv := server.newHttpServer(":8080")
	log.Fatal(srv.ListenAndServer())
}
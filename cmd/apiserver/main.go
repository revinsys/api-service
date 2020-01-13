package main

import (
	"{GIT_PATH}/internal/app/apiserver"
	"log"
)

func main() {
	server := apiserver.NewAPIServer()

	log.Println(server.Start())
}

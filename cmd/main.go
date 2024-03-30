package main

import (
	"log"

	"github.com/maiaaraujo5/controle-de-transacao/app/fx/server"
)

func main() {
	err := server.Start()
	if err != nil {
		log.Fatalf("error to run application: %v", err)
	}
}

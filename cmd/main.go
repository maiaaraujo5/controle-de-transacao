package main

import (
	"log"

	"github.com/maiaaraujo5/controle-de-transacao/app/fx/server"
	_ "github.com/maiaaraujo5/controle-de-transacao/docs"
)

// @title Controle de Transação API | Pismo
// @version 1.0
// @description Controle de Transação API
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8080
// @schemes http
func main() {
	err := server.Start()
	if err != nil {
		log.Fatalf("error to run application: %v", err)
	}
}

package main

import (
	"flag"
	"fmt"
	"log"
	"smalldoc124/service/cmd/sales-api/handlers"
	"smalldoc124/service/internal/database"

	"github.com/gin-gonic/gin"
)

func main() {
	var port, dbhost, dbschema, dbusername, dbpassword, disableTLS string
	var dbport int
	flag.StringVar(&port, "port", "3000", "port for open service")
	flag.StringVar(&dbhost, "dbhost", "localhost", "database host name")
	flag.IntVar(&dbport, "dbport", 5432, "database port")
	flag.StringVar(&dbschema, "dbschema", "smalldoc", "database schema name")
	flag.StringVar(&dbusername, "dbusername", "smalldoc", "database user name")
	flag.StringVar(&dbpassword, "dbpassword", "example", "database password")
	flag.StringVar(&disableTLS, "disableTLS", "Y", "database disableTLS[Y/n]")
	flag.Parse()
	var databaseTSL bool
	if disableTLS == "n" {
		databaseTSL = false
	} else {
		databaseTSL = true
	}
	dbConfig := database.Config{
		User:       dbusername,
		Password:   dbpassword,
		Host:       dbhost,
		Port:       dbport,
		Name:       dbschema,
		DisableTLS: databaseTSL,
	}

	db, err := database.Open(dbConfig)
	if err != nil {
		log.Fatal("connecting database fail", err)
	}
	engine := gin.Default()
	handlers.API(db, engine)
	log.Fatal(engine.Run(fmt.Sprintf(":%s", port)))
}

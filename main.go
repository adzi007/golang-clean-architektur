package main

import (
	mongodb "cobasatu/repository/mongo_repo"
	"cobasatu/transport/rest"
)

func main() {

	// Mongo Database
	mongodb.ConnectMongo()

	// Rest
	rest := rest.RestServer()
	rest.Static("assets", "./assets")
	rest.Run("localhost:8000")

}

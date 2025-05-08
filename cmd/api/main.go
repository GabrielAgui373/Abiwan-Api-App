package main

import (
	"log"

	"github.com/gabrielagui373/obiwanapp-api/internal/config"
	"github.com/gabrielagui373/obiwanapp-api/internal/routes"
	"github.com/gabrielagui373/obiwanapp-api/internal/utils"
)

func main() {
	//load configs
	dbConfig := config.LoadDBConfig()
	jwtConfig := config.LoadJWTConfig()

	//load db
	db, err := utils.InitDB(dbConfig)
	if err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}

	router := routes.SetupRouter(routes.RouterDependecies{DB: db, JWTConfig: jwtConfig})

	//init server
	log.Printf("Server running on port %s", "8080")
	if err := router.Run(":" + "8080"); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}

}

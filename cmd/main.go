package main

import (
	"fmt"
	"log"

	"github.com/SaidovZohid/auth-signin-signup-middleware/api"
	"github.com/SaidovZohid/auth-signin-signup-middleware/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	cfg := config.Load(".")
	psql := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		cfg.Postgres.Host,
		cfg.Postgres.Port,
		cfg.Postgres.User,
		cfg.Postgres.Password,
		cfg.Postgres.Database,
	)
	psqlConn, err := gorm.Open(postgres.Open(psql), &gorm.Config{})

	if err != nil {
		log.Fatalf("failed to connec to existing database: %s", err)
	}
	fmt.Println("Succesfully Connected!")

	apiServer := api.New(&api.RouteOptions{
		Cfg: psqlConn,
	})

	err = apiServer.Run(cfg.HttpPort)

	if err != nil {
		log.Fatalf("failed to run server: %s", err)
	}

	log.Print("Server Stopped!")
}

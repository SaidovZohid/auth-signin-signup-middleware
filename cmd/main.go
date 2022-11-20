package main

import (
	"fmt"
	"log"

	"github.com/SaidovZohid/auth-signin-signup-middleware/api"
	"github.com/SaidovZohid/auth-signin-signup-middleware/config"
	"github.com/SaidovZohid/auth-signin-signup-middleware/storage"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
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
	psqlConn, err := sqlx.Connect("postgres", psql)

	if err != nil {
		log.Fatalf("failed to connect to existing database: %s", err)
	}
	fmt.Println("Succesfully Connected!")

	strg := storage.NewStorageI(psqlConn)

	apiServer := api.New(&api.RouteOptions{
		Cfg: &cfg,
		Storage: strg,
	})

	err = apiServer.Run(cfg.HttpPort)

	if err != nil {
		log.Fatalf("failed to run server: %s", err)
	}

	log.Print("Server Stopped!")
}

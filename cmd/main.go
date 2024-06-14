package main

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"itlivecrm/config"
	"log"
)

func main() {
	cfg := config.Load()
	fmt.Println(cfg)
	psqlUrl := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		cfg.Postgres.Host, cfg.Postgres.Port, cfg.Postgres.User, cfg.Postgres.Password,
	)
	fmt.Println(psqlUrl)
	psqlConn, err := sqlx.Connect("postgres", psqlUrl)
	if err != nil {
		log.Fatalf("Failed to connect to database")
	}
	log.Println("Connection succesfull")
	fmt.Println("connection succesful")
	_ = psqlConn
}

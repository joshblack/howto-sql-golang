package main

import (
	"github.com/joshblack/howto-sql-golang/config"
	"github.com/joshblack/howto-sql-golang/database"
	"log"
)

func main() {
	cfg := &config.Config{}
	err := cfg.Setup()
	if err != nil {
		log.Fatalf("error parsing configuration: %s", err)
	}

	dbConnString := database.NewConnString(
		cfg.DatabaseUser,
		cfg.DatabasePassword,
		cfg.DatabaseAddr,
		cfg.DatabaseName,
	)

	log.Println(dbConnString)
}

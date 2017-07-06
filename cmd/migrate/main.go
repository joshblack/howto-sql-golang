package main

import (
	"github.com/joshblack/howto-sql-golang/config"
	"github.com/joshblack/howto-sql-golang/database"
	// "github.com/mattes/migrate"
	"log"
	"os"
	"path"
	// "path/filepath"
)

func main() {
	cfg := &config.Config{}
	err := cfg.Setup()
	if err != nil {
		log.Fatalf("error parsing configuration: %s", err)
	}

	dbConnString := database.NewDatabaseConnString(
		cfg.DatabaseUser,
		cfg.DatabasePassword,
		cfg.DatabaseAddr,
		cfg.DatabaseName,
	)

	log.Println(dbConnString)

	ex, err := os.Executable()
	if err != nil {
		log.Fatal(err)
	}
	exPath := path.Dir(ex)
	// migrationsPath := filepath.Rel("../../database/migrations", exPath)

	log.Println(exPath)

	argsWithoutProg := os.Args[1:]

	log.Println(argsWithoutProg)

	// m, err := migrate.New(
	// database.NewDatabaseConnString(
	// cfg.DatabaseUser,
	// cfg.DatabasePassword,
	// cfg.DatabaseAddr,
	// cfg.DatabaseName,
	// ),
	// )
	// if err != nil {
	// log.Fatal(err)
	// }

	// if err := m.Up(); err != nil {
	// log.Fatal(err)
	// }
}

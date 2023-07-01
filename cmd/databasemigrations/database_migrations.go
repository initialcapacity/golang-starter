package main

import (
	"flag"
	"log"
	"os"

	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/initialcapacity/golang-starter/pkg/migrationsupport"
)

func main() {
	back := flag.Bool("back", false, "migrate down")
	dev := flag.Bool("development", false, "migrate development database")
	test := flag.Bool("test", false, "migrate test database")
	flag.Parse()

	var databaseUrls []string
	if *dev {
		databaseUrls = append(databaseUrls,
			"postgres://starter:starter@localhost:5432/starter_development?sslmode=disable")
	}
	if *test {
		databaseUrls = append(databaseUrls,
			"postgres://starter:starter@localhost:5432/starter_test?sslmode=disable")
	}
	databaseURL := os.Getenv("DATABASE_URL")
	if databaseURL != "" && len(databaseUrls) == 0 {
		databaseUrls = append(databaseUrls, databaseURL)
	}
	log.Println(databaseUrls)
	schemaMigration := migrationsupport.NewSchemaMigration("file://./databases/starter", databaseUrls)

	if *back {
		schemaMigration.BackOne()
	} else {
		schemaMigration.Forward()
	}
}

package main

import (
	"flag"

	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/initialcapacity/golang-starter/pkg/migrationsupport"
)

func main() {
	back := *flag.Bool("back", false, "migrate down")
	flag.Parse()

	schemaMigration := migrationsupport.NewSchemaMigration(
		"file://./databases/starter",
		[]string{"postgres://starter:starter@localhost:5432/starter_test?sslmode=disable"},
	)

	if back {
		schemaMigration.BackOne()
	} else {
		schemaMigration.Forward()
	}
}

package migrationsupport

import (
	"log"

	"github.com/golang-migrate/migrate/v4"
)

type SchemaMigration struct {
	migrations []*migrate.Migrate
}

func NewSchemaMigration(source string, databaseUrls []string) SchemaMigration {
	migrations := make([]*migrate.Migrate, 0, len(databaseUrls))

	for _, url := range databaseUrls {
		migration, err := migrate.New(source, url)
		if err != nil {
			log.Fatalf("failed to connect to %s", err)
		}
		migration.Log = logger{}
		migrations = append(migrations, migration)
	}

	return SchemaMigration{migrations: migrations}
}

func (m SchemaMigration) Forward() {
	log.Println("migrating forward")

	for _, migration := range m.migrations {
		handleError(migration.Up())
	}
}

func (m SchemaMigration) BackOne() {
	log.Println("migrating back one migration")

	for _, migration := range m.migrations {
		handleError(migration.Steps(-1))
	}
}

func handleError(err error) {
	if err == migrate.ErrNoChange {
		log.Printf("no new migrations detected: %s\n", err)
	} else if err != nil {
		log.Fatalf("failed to migrate: %s", err)
	}
}

type logger struct{}

func (l logger) Verbose() bool { return true }

func (l logger) Printf(format string, v ...interface{}) {
	log.Printf(format, v...)
}

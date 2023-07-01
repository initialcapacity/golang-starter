package main

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	_ = os.Setenv("POSTGRESQL_URL", "postgres://postgres:postgres@localhost:5432/postgres?sslmode=disable")
	analyzer := newDataCollector()
	assert.NotNil(t, analyzer)
}

func TestMissingPostgres(t *testing.T) {
	_ = os.Setenv("POSTGRESQL_URL", "")
	defer func() {
		if err := recover(); err == nil {
			t.Fail()
		}
	}()
	main()
}

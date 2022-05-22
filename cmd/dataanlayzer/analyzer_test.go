package main

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func Test(t *testing.T) {
	_ = os.Setenv("POSTGRESQL_URL", "postgres://postgres:postgres@localhost:5432/postgres?sslmode=disable")
	go main()
}

func TestNew(t *testing.T) {
	_ = os.Setenv("POSTGRESQL_URL", "postgres://postgres:postgres@localhost:5432/postgres?sslmode=disable")
	analyzer := newDataAnalyzer()
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

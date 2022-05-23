package main

import (
  "fmt"
  "github.com/stretchr/testify/assert"
  "os"
  "testing"
)

func Test(t *testing.T) {
  err := os.Setenv("POSTGRESQL_URL", "postgres://postgres:postgres@localhost:5432/postgres?sslmode=disable")
  if err != nil {
    fmt.Println(err)
    t.Fail()
  }
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

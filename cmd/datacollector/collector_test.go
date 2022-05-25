package main

import (
  "fmt"
  "github.com/stretchr/testify/assert"
  "testing"
)

func Test(t *testing.T) {
  fmt.Println("testing main.")
  go main()
}

func TestNew(t *testing.T) {
  collector := newDataCollector()
  assert.NotNil(t, collector)
}

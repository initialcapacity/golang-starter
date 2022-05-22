package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test(t *testing.T) {
	go main()
}

func TestNew(t *testing.T) {
	collector := newDataCollector()
	assert.NotNil(t, collector)
}

package dataanalyzer_test

import (
	"testing"

	"github.com/initialcapacity/golang-starter/internal/dataanalyzer"
	"github.com/initialcapacity/golang-starter/pkg/databasesupport"
	"github.com/stretchr/testify/assert"
)

func TestDataGateway_Find(t *testing.T) {
	db, _ := databasesupport.Open("postgres://starter:starter@localhost:5432/starter_test?sslmode=disable")
	gateway := dataanalyzer.DataGateway{DB: db}

	find, err := gateway.Find()
	assert.NoError(t, err)
	assert.Equal(t, []byte("data"), find[0].Data)
}

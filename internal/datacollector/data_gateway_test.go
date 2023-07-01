package datacollector_test

import (
	"testing"

	"github.com/initialcapacity/golang-starter/internal/datacollector"
	"github.com/initialcapacity/golang-starter/pkg/databasesupport"
	"github.com/stretchr/testify/assert"
)

func TestDataGateway_Find(t *testing.T) {
	db, _ := databasesupport.Open("postgres://starter:starter@localhost:5432/starter_test?sslmode=disable")
	gateway := datacollector.DataGateway{DB: db}

	find, err := gateway.Find()
	assert.NoError(t, err)
	assert.Equal(t, "https://feed.infoq.com/", find[0].Url)
}

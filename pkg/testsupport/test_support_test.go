package testsupport_test

import (
	"context"
	"testing"

	"github.com/gorilla/mux"
	"github.com/initialcapacity/golang-starter/pkg/healthsupport"
	"github.com/initialcapacity/golang-starter/pkg/testsupport"
	"github.com/stretchr/testify/assert"
)

func TestHealth(t *testing.T) {
	r := mux.NewRouter()
	r.HandleFunc("/health", healthsupport.HandlerFunction).Methods("GET")
	server := testsupport.Server(r)
	assert.True(t, testsupport.WaitForHealthy(server, "health"))
	_ = server.Shutdown(context.Background())
}

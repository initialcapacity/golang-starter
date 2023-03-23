package healthsupport_test

import (
	"context"
	"fmt"
	"io"
	"net/http"
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

	resp, _ := http.Get(fmt.Sprintf("http://%s/health", server.Addr))
	body, _ := io.ReadAll(resp.Body)
	assert.Equal(t, "{\"status\":\"pass\"}", string(body))

	_ = server.Shutdown(context.Background())
}

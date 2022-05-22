package websupport_test

import (
	"bytes"
	"github.com/gorilla/mux"
	"github.com/initialcapacity/golang-starter/pkg/healthsupport"
	"github.com/initialcapacity/golang-starter/pkg/testsupport"
	"github.com/initialcapacity/golang-starter/pkg/websupport"
	"github.com/initialcapacity/golang-starter/pkg/websupport/test"
	"github.com/stretchr/testify/assert"
	"io"
	"net"
	"net/http/httptest"
	"testing"
)

func TestServer(t *testing.T) {
	listener, _ := net.Listen("tcp", "localhost:0")
	server := websupport.Create(listener.Addr().String(), func(router *mux.Router) {
		router.HandleFunc("/health", healthsupport.HandlerFunction).Methods("GET")
	})
	go websupport.Start(server, listener)
	assert.True(t, testsupport.WaitForHealthy(server, "health"))
	websupport.Stop(server)
}

func TestModelAndView(t *testing.T) {
	w := &httptest.ResponseRecorder{Body: new(bytes.Buffer)}
	_ = websupport.ModelAndView(w, &websupport_test.Resources, "test",
		websupport.Model{Map: map[string]interface{}{"name": "aName"}})
	body, _ := io.ReadAll(w.Body)
	assert.Equal(t, `
    <!DOCTYPE html>
    <html lang="en">
    <body>aName
    </body>
    </html>`, string(body))
}

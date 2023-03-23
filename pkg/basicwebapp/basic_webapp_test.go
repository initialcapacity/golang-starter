package basicwebapp_test

import (
	"context"
	"fmt"
	"io"
	"net"
	"net/http"
	"testing"

	"github.com/initialcapacity/golang-starter/pkg/basicwebapp"
	"github.com/initialcapacity/golang-starter/pkg/testsupport"
	"github.com/initialcapacity/golang-starter/pkg/websupport"
	"github.com/stretchr/testify/assert"
)

func TestApp(t *testing.T) {
	listener, _ := net.Listen("tcp", "localhost:0")

	basic := basicwebapp.NewBasicApp()
	server := websupport.Create(listener.Addr().String(), basic.LoadHandlers())

	go websupport.Start(server, listener)
	testsupport.WaitForHealthy(server, "health")

	resp, _ := http.Get(fmt.Sprintf("http://%s/", server.Addr))
	body, _ := io.ReadAll(resp.Body)
	assert.Contains(t, string(body), "AppContinuum[]")
	assert.Contains(t, string(body), "application with background workers.")

	_ = server.Shutdown(context.Background())
}

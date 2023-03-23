package main

import (
	"fmt"

	"github.com/initialcapacity/golang-starter/pkg/basicwebapp"
	"github.com/initialcapacity/golang-starter/pkg/websupport"

	"log"
	"net"
	"net/http"
	"os"
)

func App(addr string) *http.Server {
	return websupport.Create(addr, basicwebapp.NewBasicApp().LoadHandlers())
}

func newApp(addr string) (*http.Server, net.Listener) {
	if found := os.Getenv("PORT"); found != "" {
		host, _, _ := net.SplitHostPort(addr)
		addr = fmt.Sprintf("%v:%v", host, found)
	}
	log.Printf("Found server address %v", addr)
	listener, _ := net.Listen("tcp", addr)
	return App(listener.Addr().String()), listener
}

func main() {
	websupport.Start(newApp("0.0.0.0:8888"))
}

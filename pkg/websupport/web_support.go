package websupport

import (
	"context"
	"net"
	"net/http"

	"github.com/gorilla/mux"
)

func Create(addr string, handlers func(x *mux.Router)) *http.Server {
	router := mux.NewRouter()
	router.StrictSlash(true)
	handlers(router)
	server := http.Server{
		Addr:    addr,
		Handler: router,
	}
	return &server
}

func Start(server *http.Server, l net.Listener) {
	err := server.Serve(l)
	if err != nil {
		return
	}
}

func Stop(server *http.Server) {
	_ = server.Shutdown(context.Background())
}

package websupport

import (
	"context"
	"embed"
	"fmt"
	"github.com/gorilla/mux"
	"html/template"
	"net"
	"net/http"
	"path/filepath"
)

type Model struct {
	Map map[string]interface{}
}

func ModelAndView(w http.ResponseWriter, resources *embed.FS, view string, data Model) error {
	views := []string{
		fmt.Sprintf("resources/templates/%v.gohtml", view),
		"resources/templates/template.gohtml",
	}
	base := filepath.Base(views[0]) // to match template names in ParseFiles
	return template.Must(template.New(base).Funcs(template.FuncMap{}).ParseFS(resources, views...)).Execute(w, data)
}

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

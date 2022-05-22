package basicwebapp

import (
	"github.com/gorilla/mux"
	"github.com/initialcapacity/golang-starter/pkg/healthsupport"
	"github.com/initialcapacity/golang-starter/pkg/metricssupport"
	"github.com/initialcapacity/golang-starter/pkg/websupport"
	"io/fs"
	"net/http"
)

type BasicApp struct {
}

func NewBasicApp() BasicApp {
	return BasicApp{}
}

func (a BasicApp) LoadHandlers() func(x *mux.Router) {
	return func(router *mux.Router) {
		router.HandleFunc("/", a.dashboard).Methods("GET")
		router.HandleFunc("/health", healthsupport.HandlerFunction)
		router.HandleFunc("/metrics", metricssupport.HandlerFunction)
		router.Use(metricssupport.Middleware)

		static, _ := fs.Sub(Resources, "resources/static")
		fileServer := http.FileServer(http.FS(static))
		router.PathPrefix("/").Handler(http.StripPrefix("/", fileServer))
	}
}

func (a *BasicApp) dashboard(writer http.ResponseWriter, req *http.Request) {
	_ = websupport.ModelAndView(writer, &Resources, "index", websupport.Model{Map: map[string]interface{}{}})
}

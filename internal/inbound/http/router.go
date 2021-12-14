// Package http DCSwitch API
//
// The purpose of this service is to provide an application to help manage db switch task
//
//      Schemes: http
//      Host: 127.0.0.1
//      Version: 0.0.1
//
// swagger:meta
package http

import (
	handlers "dcswitch/internal/inbound/http/handlers"
	"dcswitch/internal/inbound/http/middlewares"
	"dcswitch/pkg/swagger"
	"github.com/gorilla/mux"
	"net/http"
)

// InitHandlers API入口
func InitHandlers() *mux.Router {

	r := mux.NewRouter()
	r.Use(middlewares.SlowRequestMiddleware, middlewares.LoggingMiddleware, middlewares.CORSMiddleware)

	r.HandleFunc("/healthz", handlers.HealthCheck).Methods("GET")

	r.HandleFunc("/{id:[0-9]+}", handlers.GetMockSlow).Methods("GET")
	r.HandleFunc("/config", handlers.GetConfig).Methods("GET")

	// biz
	r.HandleFunc("/switch/versions", handlers.GetAllSwitchVersions).Methods("GET")
	r.HandleFunc("/switch/version/name/{id:[0-9]+}", handlers.EditSwitchVersionName).Methods("PATCH")

	return r
}

// InitDocHandler API文档Handler
func InitDocHandler(r *mux.Router) {
	r.HandleFunc("/static/swagger.json", handlers.SwaggerFile).Methods("GET")
	r.HandleFunc("/docs", swagger.DefaultDocs).Methods("GET")
	r.HandleFunc("/redoc", swagger.DefaultReDoc).Methods("GET")
}

func HandleFuncWithAuth(path string, f func(http.ResponseWriter, *http.Request), auth func(http.ResponseWriter, *http.Request), methods ...string){

}

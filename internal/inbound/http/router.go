// Package http DCSwitch API
//
// The purpose of this service is to provide an application to help manage db switch task
//
//      Schemes: http
//      Host: 127.0.0.1
//      Version: 0.0.1
//
//      Security:
//      - api_key:
//
//      SecurityDefinitions:
//      api_key:
//           type: apiKey
//           name: TOKEN
//           in: header
//
// swagger:meta
package http

import (
	"dcswitch/internal/config"
	handlers "dcswitch/internal/inbound/http/handlers"
	"dcswitch/internal/inbound/http/middlewares"
	"dcswitch/pkg/swagger"
	"github.com/gorilla/mux"
	"net/http"
)

// InitHandlers API入口
func InitHandlers() *mux.Router {

	r := mux.NewRouter()
	r.Use(middlewares.RecoverWrap, middlewares.SlowRequestMiddleware, middlewares.LoggingMiddleware, middlewares.CORSMiddleware)

	r.HandleFunc("/healthz", handlers.HealthCheck).Methods("GET")

	r.HandleFunc("/{id:[0-9]+}", handlers.GetMockSlow).Methods("GET")
	r.HandleFunc("/config", handlers.GetConfig).Methods("GET")

	// biz
	r.HandleFunc("/switch/versions", handlers.GetAllSwitchVersions).Methods("GET")
	r.HandleFunc("/switch/version/date/{id:[0-9]+}", handlers.EditSwitchVersionDate).Methods("PATCH")
	r.HandleFunc("/task/switch/module/detail", handlers.CreateModuleDetailTask).Methods("POST")

	r.HandleFunc("/version/{id:[0-9]+}", handlers.GetVersionById).Methods("GET")
	r.HandleFunc("/switch/preparations/{versionId:[0-9]+}", handlers.GetAllSwitchPreparations).Methods("GET")
	r.HandleFunc("/switch/preparation/{id:[0-9]+}", handlers.GetSwitchPreparationById).Methods("GET")
	r.HandleFunc("/switch/preparation/edit/{id:[0-9]+}", handlers.EditSwitchPreparationById).Methods("GET")
	r.HandleFunc("/switch/modules/{versionId:[0-9]+}", handlers.GetAllSwitchModules).Methods("GET")
	r.HandleFunc("/switch/details/{versionId:[0-9]+}", handlers.GetAllSwitchDetails).Methods("GET")
	r.HandleFunc("/switch/module/progress/{moduleId:[0-9]+}", handlers.GetProgressByModuleId).Methods("GET")
	r.HandleFunc("/switch/version/progress/{versionId:[0-9]+}", handlers.GetProgressByVersionId).Methods("GET")
	r.HandleFunc("/switch/version/failed/details/{versionId:[0-9]+}", handlers.GetFailedDetailsByVersionId).Methods("GET")
	r.HandleFunc("/switch/module/details/{moduleId:[0-9]+}", handlers.GetDetailsByModuleId).Methods("GET")
	r.HandleFunc("/switch/detail/editComment/{id:[0-9]+}", handlers.EditDetailCommentById).Methods("GET")
	r.HandleFunc("/switch/detail/editEndTime/{id:[0-9]+}", handlers.EditDetailEndTimeById).Methods("GET")
	r.HandleFunc("/switch/range/upload/{moduleId:[0-9]+}", handlers.SwitchRangeUpload).Methods("POST")
	r.HandleFunc("/switch/modules/domain/{domainId:[0-9]+}", handlers.GetModulesByDomainId).Methods("GET")
	r.HandleFunc("/switch/module/domain/{moduleId:[0-9]+}", handlers.GetModuleByModuleId).Methods("GET")
	r.HandleFunc("/switch/domain/{id:[0-9]+}", handlers.GetSwitchDomainById).Methods("GET")
	return r
}

// InitDocHandler API文档Handler
func InitDocHandler(r *mux.Router) {
	fs := http.FileServer(http.FS(config.SwaggerFS))
	r.PathPrefix("/static").Handler(http.StripPrefix("/static", fs))
	r.HandleFunc("/docs", swagger.WrappedDocsHandler(swagger.DocsOpts{
		DocsSwaggerJsonURL: "/static/swagger.json",
		DocsJsURL:          "/static/swagger-ui-bundle.js",
		DocsCssURL:         "/static/swagger-ui.css",
		IconURL:            "/static/dcswitch.png",
	})).Methods("GET")
	r.HandleFunc("/redoc", swagger.WrappedReDocHandler(swagger.RedocOpts{
		SpecSwaggerJsonURL: "/static/swagger.json",
		RedocJsURL:         "/static/redoc.standalone.js",
		RedocCssURL:        "/static/redoc.standalone.css",
		IconURL:            "/static/dcswitch.png",
	})).Methods("GET")
}

func HandleFuncWithAuth(path string, f func(http.ResponseWriter, *http.Request), auth func(http.ResponseWriter, *http.Request), methods ...string) {

}

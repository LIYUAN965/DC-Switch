package handlers

import (
	"dcswitch/internal/config"
	log "github.com/sirupsen/logrus"
	"net/http"
)

func SwaggerFile(rw http.ResponseWriter, r *http.Request) {
	b, err := config.SwaggerFile.ReadFile("swagger.json")
	if err != nil {
		log.Error(err)
	}
	_, _ = rw.Write(b)
}

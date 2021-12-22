//go:generate swagger generate spec -o ../../internal/static/swagger.json
//go install github.com/go-swagger/go-swagger/cmd/swagger@latest

package main

import (
	"dcswitch/internal/config"
	httpApi "dcswitch/internal/inbound/http"
	"dcswitch/pkg/logger"
	log "github.com/sirupsen/logrus"
	"net/http"
	_ "net/http/pprof"
)

func main() {
	// init logger, config, db
	logger.InitLogger()
	config.GlobalConf.InitConfig()
	config.InitDB()
	// init servers
	servers := make(map[string]*http.Server)
	servers["BaseServer"] = httpApi.InitServer()
	servers["PProfServer"] = httpApi.InitPProfServer()
	for n, s := range servers {
		go func(n string, s *http.Server) {
			log.Infof("%s listen on %v\n", n, s.Addr)
			err := s.ListenAndServe()
			switch err {
			case http.ErrServerClosed:
				log.Error(err)
			default:
				log.Fatal(err)
			}
		}(n, s)
	}
	// Graceful shutdown
	httpApi.GracefulStop(servers)
}

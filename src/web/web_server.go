package web

import (
	"net/http"
	"log"
	_ "net/http/pprof"
	"github.com/xsank/EasyProxy/src/util"
	"github.com/xsank/EasyProxy/src/config"
)

type WebServer struct {
	host string
	port uint16
}

func (server *WebServer) Init(config *config.Config) {
	server.host = config.Host
	server.port = config.WebPort
}

func (server *WebServer) Start() {
	go func() {
		server.AddHandler()
		url := util.HostPortToAddress(server.host, server.port)
		err := http.ListenAndServe(url, nil)
		if err != nil {
			log.Println("create web server failed:", err)
		} else {
			log.Println("create web server success")
		}
	}()
}

func (server *WebServer) AddHandler() {
	http.HandleFunc(StatisticsUrl, Statistic)
}
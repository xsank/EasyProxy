package main

import (
	"os"
	"path/filepath"
	"runtime"
	"syscall"
	"os/signal"
	"github.com/xsank/EasyProxy/src/web"
	"github.com/xsank/EasyProxy/src/config"
	"github.com/xsank/EasyProxy/src/util"
	"github.com/xsank/EasyProxy/src/log"
	gw"github.com/xsank/EasyProxy/src/gateway"
)

const
(
	DefaultConfigFile = "conf/default.json"
	DefaultLogFile = "esayproxy.log"
)

type EasyServer struct {
	webServer   *web.WebServer
	proxyServer *gw.ProxyServer
}

func CreateEasyServer() *EasyServer {
	return &EasyServer{webServer:new(web.WebServer), proxyServer:new(gw.ProxyServer)}
}

func (easyServer *EasyServer)Init(config *config.Config) {
	easyServer.webServer.Init(config)
	easyServer.proxyServer.Init(config)
}

func (easyServer *EasyServer)Start() {
	easyServer.webServer.Start()
	easyServer.proxyServer.Start()
}

func (easyServer *EasyServer) CatchStopSignal() {
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGQUIT, syscall.SIGKILL)
	go func() {
		<-sig
		easyServer.Stop()
	}()
}

func (easyServer *EasyServer) Stop() {
	easyServer.proxyServer.Stop()
}

func main() {
	log.Init(DefaultLogFile)
	
	homePath := util.HomePath()
	config, err := config.Load(filepath.Join(homePath, DefaultConfigFile))

	if err == nil {
		runtime.GOMAXPROCS(config.MaxProcessor)

		easyServer := CreateEasyServer()
		easyServer.Init(config)
		easyServer.CatchStopSignal()
		easyServer.Start()
	}
}
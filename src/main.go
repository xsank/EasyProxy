package main

import (
	gw"github.com/xsank/EasyProxy/src/gateway"
	"github.com/xsank/EasyProxy/src/config"
	"github.com/xsank/EasyProxy/src/util"
	"path/filepath"
	"github.com/xsank/EasyProxy/src/log"
	"github.com/xsank/EasyProxy/src/web"
	"runtime"
	"syscall"
	"os/signal"
	"os"
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
	homePath := util.HomePath()
	log.Init(DefaultLogFile)
	config, err := config.Load(filepath.Join(homePath, DefaultConfigFile))

	if err == nil {
		runtime.GOMAXPROCS(config.MaxProcessor)

		easyServer := CreateEasyServer()
		easyServer.Init(config)
		easyServer.CatchStopSignal()
		easyServer.Start()
	}
}
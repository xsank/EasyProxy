package main

import (
	gw"github.com/xsank/EasyProxy/src/gateway"
	"github.com/xsank/EasyProxy/src/config"
	"github.com/xsank/EasyProxy/src/util"
	"path/filepath"
	"github.com/xsank/EasyProxy/src/log"
	"github.com/xsank/EasyProxy/src/web"
	"runtime"
)

const
(
	DefaultConfigFile = "conf/default.json"
	DefaultLogFile = "esayproxy.log"
)

func main() {
	runtime.GOMAXPROCS(2 * runtime.NumCPU())

	homePath := util.HomePath()
	log.Init(DefaultLogFile)
	config, err := config.Load(filepath.Join(homePath, DefaultConfigFile))

	if err == nil {
		webServer := new(web.WebServer)
		webServer.Init(config)
		webServer.Start()
		server := new(gw.ProxyServer)
		server.Init(config)
		server.Start()
	}
}
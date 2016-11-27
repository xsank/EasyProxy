package main

import (
	gw"github.com/xsank/EasyProxy/src/gateway"
	"github.com/xsank/EasyProxy/src/config"
	"github.com/xsank/EasyProxy/src/util"
	"path/filepath"
	"github.com/xsank/EasyProxy/src/log"
	"github.com/xsank/EasyProxy/src/web"
)

const DefaultConfigFile = "conf/default.json"
const DefaultLogFile = "esayproxy.log"

func main() {
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
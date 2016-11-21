package main

import (
	gw"github.com/xsank/EasyProxy/src/gateway"
	"github.com/xsank/EasyProxy/src/config"
	"github.com/xsank/EasyProxy/src/util"
	"path/filepath"
)

const DefaultConfigFile = "conf/default.json"

func main() {
	homePath := util.HomePath()
	config, err := config.Load(filepath.Join(homePath, DefaultConfigFile))
	if err == nil {
		server := new(gw.ProxyServer)
		server.Init(config)
		server.Start()
	}
}
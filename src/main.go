package main

import (
	"github.com/xsank/EasyProxy/src/net"
	"github.com/xsank/EasyProxy/src/config"
	"github.com/xsank/EasyProxy/src/util"
	"path/filepath"
)

const DefaultConfigFile = "conf/default.json"

func main() {
	homePath := util.HomePath()
	config, err := config.Load(filepath.Join(homePath, DefaultConfigFile))
	if err == nil {
		server := &net.ProxyServer{}
		server.Init(&config)
		server.Start()
	}
}
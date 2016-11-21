package gateway

import (
	"net"
	"log"
	"github.com/xsank/EasyProxy/src/proxy"
	"time"
	"github.com/xsank/EasyProxy/src/util"
	"github.com/xsank/EasyProxy/src/config"
)

const DefaultHeartBeatTime = 10

type ProxyServer struct {
	host  string
	port  uint16
	proxy proxy.Proxy
}

func (server *ProxyServer) Init(config *config.Config) {
	server.host = config.Host
	server.port = config.Port
	server.setProxy(config)
}

func (server *ProxyServer) setProxy(config *config.Config) {
	server.proxy = new(proxy.EasyProxy)
	server.proxy.Init(config)
}

func (server ProxyServer) Address() string {
	return util.HostPortToAddress(server.host, server.port)
}

func (server *ProxyServer) Start() {
	local, err := net.Listen("tcp", server.Address())
	if err != nil {
		log.Panic("proxy server start error:$s", err)
	}
	log.Println("easyproxy server start ok")
	defer local.Close()
	server.heartBeat()
	for {
		con, err := local.Accept()
		if (err == nil) {
			go server.proxy.Dispatch(con)
		} else {
			log.Println("client connect server error:", err)
		}
	}
}

func (server ProxyServer) heartBeat() {
	ticker := time.NewTicker(time.Second * DefaultHeartBeatTime)
	go func() {
		for {
			select {
			case <-ticker.C:
				server.proxy.Check()
			}
		}
	}()
}



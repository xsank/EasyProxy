package gateway

import (
	"net"
	"log"
	"time"
	"github.com/xsank/EasyProxy/src/proxy"
	"github.com/xsank/EasyProxy/src/util"
	"github.com/xsank/EasyProxy/src/config"
)

type ProxyServer struct {
	host     string
	port     uint16
	beattime int
	listener net.Listener
	proxy    proxy.Proxy
	on       bool
}

func (server *ProxyServer) Init(config *config.Config) {
	server.on = false
	server.host = config.Host
	server.port = config.Port
	server.beattime = config.Heartbeat
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
		log.Panic("proxy server start error:", err)
	}
	log.Println("easyproxy server start ok")
	server.listener = local
	server.on = true
	server.heartBeat()
	for server.on {
		con, err := server.listener.Accept()
		if (err == nil) {
			go server.proxy.Dispatch(con)
		} else {
			log.Println("client connect server error:", err)
		}
	}
	defer server.listener.Close()
}

func (server *ProxyServer) heartBeat() {
	ticker := time.NewTicker(time.Second * time.Duration(server.beattime))
	go func() {
		for {
			select {
			case <-ticker.C:
				server.proxy.Check()
			}
		}
	}()
}

func (server *ProxyServer) Stop() {
	server.listener.Close()
	server.proxy.Close()
	server.on = false
	log.Println("easyproxy server stop ok")
}



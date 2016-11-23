package proxy

import (
	"net"
	"io"
	"github.com/xsank/EasyProxy/src/proxy/schedule"
	"time"
	"log"
	"github.com/xsank/EasyProxy/src/config"
	s"github.com/xsank/EasyProxy/src/structure"
)

const (
	DefaultTimeoutTime = 3
)

type EasyProxy struct {
	data           *ProxyData
	strategy       schedule.Strategy
	channelManager *s.ChannelManager
}

func (proxy *EasyProxy) Init(config *config.Config) {
	proxy.data = new(ProxyData)
	proxy.data.Init(config)
	proxy.channelManager = new(s.ChannelManager)
	proxy.channelManager.Init()
	proxy.setStrategy(config.Strategy)
}

func (proxy *EasyProxy) setStrategy(name string) {
	switch name {
	case "random":
		proxy.strategy = new(schedule.Random)
	case "poll":
		proxy.strategy = new(schedule.Poll)
	default:
		proxy.strategy = new(schedule.Random)
	}
}

func (proxy *EasyProxy) Check() {
	for _, backend := range proxy.data.backends {
		_, err := net.Dial("tcp", backend.Url())
		if err != nil {
			proxy.Clean(backend.Url())
		}
	}
	for _, deadend := range proxy.data.backends {
		_, err := net.Dial("tcp", deadend.Url())
		if err == nil {
			proxy.Recover(deadend.Url())
		}
	}
}

func (proxy *EasyProxy) Dispatch(con net.Conn) {
	urls := proxy.data.BackendUrls()
	url := proxy.strategy.Choose(urls)
	proxy.transfer(con, url)
}

func (proxy *EasyProxy) safeCopy(local net.Conn, remote net.Conn) {
	io.Copy(local, remote)
	defer local.Close()
}

func (proxy *EasyProxy) transfer(local net.Conn, remote string) {
	remoteConn, err := net.DialTimeout("tcp", remote, DefaultTimeoutTime * time.Second)
	if err != nil {
		proxy.Clean(remoteConn.RemoteAddr().String())
		log.Println("connect error:%s", err)
		return
	}
	proxy.channelManager.PutChannelPair(local, remoteConn)
	go proxy.safeCopy(local, remoteConn)
	go proxy.safeCopy(remoteConn, local)
}

func (proxy *EasyProxy) Clean(url string) {
	proxy.data.cleanBackend(url)
}

func (proxy *EasyProxy) Recover(url string) {
	proxy.data.cleanDeadend(url)
}
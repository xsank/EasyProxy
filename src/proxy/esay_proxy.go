package proxy

import (
	"net"
	"io"
	"time"
	"log"
	"github.com/xsank/EasyProxy/src/proxy/schedule"
	"github.com/xsank/EasyProxy/src/config"
	"github.com/xsank/EasyProxy/src/structure"
)

const (
	DefaultTimeoutTime = 3
)

type EasyProxy struct {
	data     *ProxyData
	strategy schedule.Strategy
}

func (proxy *EasyProxy) Init(config *config.Config) {
	proxy.data = new(ProxyData)
	proxy.data.Init(config)
	proxy.setStrategy(config.Strategy)
	InitStatistic(proxy.data)
}

func (proxy *EasyProxy) setStrategy(name string) {
	proxy.strategy = schedule.GetStrategy(name)
	proxy.strategy.Init()
}

func (proxy *EasyProxy) Check() {
	for _, backend := range proxy.data.Backends {
		_, err := net.Dial("tcp", backend.Url())
		if err != nil {
			proxy.Clean(backend.Url())
		}
	}
	for _, deadend := range proxy.data.Deads {
		_, err := net.Dial("tcp", deadend.Url())
		if err == nil {
			proxy.Recover(deadend.Url())
		}
	}
}

func (proxy *EasyProxy) isBackendAvailable() bool {
	return len(proxy.data.Backends) > 0
}

func (proxy *EasyProxy) Dispatch(con net.Conn) {
	if proxy.isBackendAvailable() {
		servers := proxy.data.BackendUrls()
		url := proxy.strategy.Choose(con.RemoteAddr().String(), servers)
		proxy.transfer(con, url)
	} else {
		con.Close()
		log.Println("no backends available now,please check your server!")
	}
}

func (proxy *EasyProxy) safeCopy(from net.Conn, to net.Conn, sync chan int) {
	io.Copy(from, to)
	defer from.Close()
	sync <- 1
}

func (proxy *EasyProxy) putChannel(channel *structure.Channel) {
	proxy.data.ChannelManager.PutChannel(channel)
}

func (proxy *EasyProxy) closeChannel(channel *structure.Channel, sync chan int) {
	for i := 0; i < structure.ChannelPairNum; i++ {
		<-sync
	}
	proxy.data.ChannelManager.DeleteChannel(channel)
}

func (proxy *EasyProxy) transfer(local net.Conn, remote string) {
	remoteConn, err := net.DialTimeout("tcp", remote, DefaultTimeoutTime * time.Second)
	if err != nil {
		local.Close()
		proxy.Clean(remote)
		log.Println("connect backend error:%s", err)
		return
	}
	sync := make(chan int, 1)
	channel := structure.Channel{SrcConn:local, DstConn:remoteConn}
	go proxy.putChannel(&channel)
	go proxy.safeCopy(local, remoteConn, sync)
	go proxy.safeCopy(remoteConn, local, sync)
	go proxy.closeChannel(&channel, sync)
}

func (proxy *EasyProxy) Clean(url string) {
	proxy.data.cleanBackend(url)
}

func (proxy *EasyProxy) Recover(url string) {
	proxy.data.cleanDeadend(url)
}

func (proxy *EasyProxy) Close() {
	proxy.data.Clean()
}
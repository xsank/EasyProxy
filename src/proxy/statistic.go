package proxy

import (
	"github.com/xsank/EasyProxy/src/util"
)

var statistic = new(Statistic)

type Statistic struct {
	Clients   map[string]Client
	Services  map[string]Service
	proxyData *ProxyData
}

type Client struct {
	Host  string
	Count int
}

type Service struct {
	Url    string
	Count  int
	Status string
}

func InitStatistic(proxyData *ProxyData) {
	statistic.proxyData = proxyData
}

func StatisticData() *Statistic {
	return statistic
}

func Record() {
	statistic.Clients = make(map[string]Client)
	statistic.Services = make(map[string]Service)
	for _, server := range statistic.proxyData.Backends {
		statistic.Services[server.Url()] = Service{Url:server.Url(), Count:0, Status:"on"}
	}
	for _, server := range statistic.proxyData.Deads {
		statistic.Services[server.Url()] = Service{Url:server.Url(), Count:0, Status:"off"}
	}
	for _, channel := range statistic.proxyData.ChannelManager.GetChannels() {
		host := util.UrlToHost(channel.SrcUrl())
		serverUrl := channel.DstUrl()
		var client Client
		var service Service
		if _, ok := statistic.Clients[host]; ok {
			client = statistic.Clients[host]
			client.Count += 1
		} else {
			client = Client{Host:host, Count:1}
		}
		service = statistic.Services[serverUrl]
		service.Count += 1
		statistic.Clients[host] = client
		statistic.Services[serverUrl] = service
	}
}


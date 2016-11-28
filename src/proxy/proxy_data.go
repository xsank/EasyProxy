package proxy

import (
	"github.com/xsank/EasyProxy/src/config"
	"github.com/xsank/EasyProxy/src/structure"
)

type ProxyData struct {
	Service        string
	Host           string
	Port           uint16
	Backends       map[string]structure.Backend
	Deads          map[string]structure.Backend
	ChannelManager *structure.ChannelManager
}

func (proxyData *ProxyData) Init(config *config.Config) {
	proxyData.Service = config.Service
	proxyData.Host = config.Host
	proxyData.Port = config.Port
	proxyData.ChannelManager = new(structure.ChannelManager)
	proxyData.ChannelManager.Init()
	proxyData.setBackends(config.Backends)
}

func (proxyData *ProxyData) setBackends(backends []structure.Backend) {
	proxyData.Backends = make(map[string]structure.Backend)
	for _, backend := range backends {
		proxyData.Backends[backend.Url()] = backend
	}
	proxyData.Deads = make(map[string]structure.Backend)
}

func (proxyData ProxyData) BackendUrls() []string {
	_map := proxyData.Backends
	keys := make([]string, 0, len(_map))
	for k := range _map {
		keys = append(keys, k)
	}
	return keys
}

func (proxyData *ProxyData) cleanBackend(url string) {
	delete(proxyData.Backends, url)
	proxyData.Deads[url] = proxyData.Backends[url]
}

func (proxyData *ProxyData) cleanDeadend(url string) {
	delete(proxyData.Deads, url)
	proxyData.Backends[url] = proxyData.Deads[url]
}

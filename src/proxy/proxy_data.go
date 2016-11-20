package proxy

import (
	"github.com/xsank/EasyProxy/src/structure"
	"github.com/xsank/EasyProxy/src/config"
)

type ProxyData struct {
	service  string
	host     string
	port     uint16
	backends map[string]structure.Backend
	deads    map[string]structure.Backend
}

func (proxyData *ProxyData) Init(config *config.Config) {
	proxyData.service = config.Service
	proxyData.host = config.Host
	proxyData.port = config.Port
	proxyData.setBackends(config.Backends)
}

func (proxyData *ProxyData) setBackends(backends []structure.Backend) {
	proxyData.backends = make(map[string]structure.Backend)
	for _, backend := range backends {
		proxyData.backends[backend.Url()] = backend
	}
	proxyData.deads = make(map[string]structure.Backend)
}

func (proxyData ProxyData) BackendUrls() []string {
	_map := proxyData.backends
	keys := make([]string, 0, len(_map))
	for k := range _map {
		keys = append(keys, k)
	}
	return keys
}

func (proxyData *ProxyData) cleanBackend(url string) {
	delete(proxyData.backends, url)
	proxyData.deads[url] = proxyData.backends[url]
}

func (proxyData *ProxyData) cleanDeadend(url string) {
	delete(proxyData.deads, url)
	proxyData.backends[url] = proxyData.deads[url]
}

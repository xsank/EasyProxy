package proxy

import (
	"sync"
	"github.com/xsank/EasyProxy/src/structure"
	"github.com/xsank/EasyProxy/src/config"
)

type ProxyData struct {
	Service        string
	Host           string
	Port           uint16
	Backends       map[string]structure.Backend
	Deads          map[string]structure.Backend
	ChannelManager *structure.ChannelManager
	mutex          *sync.RWMutex
}

func (proxyData *ProxyData) Init(config *config.Config) {
	proxyData.Service = config.Service
	proxyData.Host = config.Host
	proxyData.Port = config.Port
	proxyData.ChannelManager = new(structure.ChannelManager)
	proxyData.ChannelManager.Init()
	proxyData.setBackends(config.Backends)
	proxyData.mutex = new(sync.RWMutex)
}

func (proxyData *ProxyData) setBackends(backends []structure.Backend) {
	proxyData.Backends = make(map[string]structure.Backend)
	for _, backend := range backends {
		proxyData.Backends[backend.Url()] = backend
	}
	proxyData.Deads = make(map[string]structure.Backend)
}

func (proxyData ProxyData) BackendUrls() []string {
	proxyData.mutex.RLock()
	_map := proxyData.Backends
	defer proxyData.mutex.RUnlock()
	keys := make([]string, 0, len(_map))
	for k := range _map {
		keys = append(keys, k)
	}
	return keys
}

func (proxyData *ProxyData) cleanBackend(url string) {
	proxyData.mutex.Lock()
	proxyData.Deads[url] = proxyData.Backends[url]
	delete(proxyData.Backends, url)
	defer proxyData.mutex.Unlock()
}

func (proxyData *ProxyData) cleanDeadend(url string) {
	proxyData.mutex.Lock()
	proxyData.Backends[url] = proxyData.Deads[url]
	delete(proxyData.Deads, url)
	defer proxyData.mutex.Unlock()
}

func cleanMap(_map map[string]structure.Backend) {
	for k := range _map {
		delete(_map, k)
	}
}

func (proxyData *ProxyData) Clean() {
	cleanMap(proxyData.Backends)
	cleanMap(proxyData.Deads)
	proxyData.ChannelManager.Clean()
}

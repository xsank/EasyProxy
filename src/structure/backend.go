package structure

import "github.com/xsank/EasyProxy/src/util"

type Backend struct {
	Host     string `json:"host"`
	Port     uint16 `json:"port"`
}

func (backend Backend) Url() string {
	return util.HostPortToAddress(backend.Host, backend.Port)
}

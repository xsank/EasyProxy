package proxy

import (
	"net"
	"github.com/xsank/EasyProxy/src/config"
)

type Proxy interface {
	Init(config *config.Config)
	Check()
	Clean(url string)
	Recover(url string)
	Dispatch(con net.Conn)
	Close()
}

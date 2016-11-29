package schedule

import (
	"github.com/xsank/EasyProxy/src/util"
)

type IpHash struct {
}

func (strategy *IpHash) Init() {}

func (strategy *IpHash) Choose(client string, servers []string) string {
	ip := util.UrlToHost(client)
	intIp := util.IP4ToInt(ip)
	length := len(servers)
	url := servers[intIp % length]
	return url
}

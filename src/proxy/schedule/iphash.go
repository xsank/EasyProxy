package schedule

import (
	"github.com/xsank/EasyProxy/src/util"
)

type IpHash struct {
}

func (strategy *IpHash) Choose(client string, servers []string) string {
	length := len(servers)
	if length == 0 {
		return ""
	}
	ip := util.UrlToHost(client)
	intIp := util.IP4ToInt(ip)
	url := servers[intIp%length]
	return url
}

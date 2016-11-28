package schedule

import (
	"time"
)

type Random struct {
}

func (strategy *Random) Choose(client string, servers []string) string {
	length := len(servers)
	if length == 0 {
		return ""
	}
	url := servers[int(time.Now().UnixNano())%length]
	return url
}

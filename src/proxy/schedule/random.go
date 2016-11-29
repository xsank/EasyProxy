package schedule

import (
	"time"
)

type Random struct {
}

func (strategy *Random) Init() {}

func (strategy *Random) Choose(client string, servers []string) string {
	length := len(servers)
	url := servers[int(time.Now().UnixNano()) % length]
	return url
}
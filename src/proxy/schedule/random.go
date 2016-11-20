package schedule

import (
	"time"
)

type Random struct {

}

func (strategy Random) Choose(urls []string) string {
	length := len(urls)
	url := urls[int(time.Now().UnixNano()) % length]
	return url
}
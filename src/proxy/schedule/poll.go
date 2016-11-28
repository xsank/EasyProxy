package schedule

const CycleCount = 1 << 21

type Poll struct {
	counter int
}

func (strategy *Poll) Choose(client string, servers []string) string {
	length := len(servers)
	if length == 0 {
		return ""
	}
	strategy.counter = (strategy.counter + 1) % CycleCount
	url := servers[strategy.counter%length]
	return url
}

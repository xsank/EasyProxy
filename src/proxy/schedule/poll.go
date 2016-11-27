package schedule

const CycleCount = 1 << 21

type Poll struct {
	counter int
}

func (strategy *Poll) Choose(client string, servers []string) string {
	strategy.counter = (strategy.counter + 1) % CycleCount
	length := len(servers)
	url := servers[strategy.counter % length]
	return url
}

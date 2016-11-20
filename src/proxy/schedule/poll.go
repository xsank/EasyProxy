package schedule

const CycleCount = 1 << 21

type Poll struct {
	counter int
}

func (strategy *Poll) Choose(urls []string) string {
	strategy.counter = (strategy.counter + 1) % CycleCount
	length := len(urls)
	url := urls[strategy.counter % length]
	return url
}

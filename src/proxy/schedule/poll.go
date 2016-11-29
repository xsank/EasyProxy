package schedule

import "sync"

const CycleCount = 1 << 21

type Poll struct {
	counter Counter
}

type Counter struct {
	count int
	mutex *sync.Mutex
}

func (counter *Counter) Inc() {
	counter.mutex.Lock()
	counter.count = (counter.count + 1) % CycleCount
	defer counter.mutex.Unlock()
}

func (counter *Counter) Get() int {
	return counter.count
}

func (strategy *Poll) Init() {
	strategy.counter = Counter{count:0, mutex:new(sync.Mutex)}

}

func (strategy *Poll) Choose(client string, servers []string) string {
	strategy.counter.Inc()
	length := len(servers)
	url := servers[strategy.counter.Get() % length]
	return url
}

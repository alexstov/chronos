package main

import (
	"fmt"
	"time"
)

type Limiter struct {
	maxCount int
	count    int
	ticker   *time.Ticker
	ch       chan struct{}
}

func (l *Limiter) run() {
	for {
		// if counter has reached 0: block until next tick
		if l.count <= 0 {
			<-l.ticker.C
			l.count = l.maxCount
		}

		// otherwise:
		select {
		case l.ch <- struct{}{}:
			l.count--

		case <-l.ticker.C:
			l.count = l.maxCount
		}
	}
}

func (l *Limiter) Await() {
	<-l.ch
}

func NewLimiter(d time.Duration, count int) *Limiter {
	l := &Limiter{
		maxCount: count,
		count:    count,
		ticker:   time.NewTicker(d),
		ch:       make(chan struct{}),
	}
	go l.run()

	return l
}

func bursts(l *Limiter) {
	start := time.Now()
	for i := 0; i < 10; i++ {
		l.Await()
		fmt.Println("burst ", i)
	}

	fmt.Println("--- bursts:", time.Now().Sub(start), "elapsed")
}

func slowIterations(l *Limiter) {
	start := time.Now()
	for i := 0; i < 10; i++ {
		<-time.After(250 * time.Millisecond)
		l.Await()
		fmt.Println("iteration ", i)
	}

	fmt.Println("--- slowIterations:", time.Now().Sub(start), "elapsed")
}

func main() {
	fmt.Println("Hello, 世界")

	l := NewLimiter(time.Second, 5)

	bursts(l)
	slowIterations(l)
}

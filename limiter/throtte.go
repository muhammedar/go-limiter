package limiter

import (
	"time"
)

// BuildChannel is responsible for building a channel which will be used to store and pass the requests
func BuildChannel(reqPerSec, bufferSize int) (c chan time.Time) {
	// building a the buffer using go-channels, and initializing it with time.now
	c = make(chan time.Time, bufferSize)
	for i := 0; i < bufferSize; i++ {
		c <- time.Now()
	}

	// ticker will help organizing the time accourding to the needed requests per seconds
	t := time.NewTicker(time.Second / time.Duration(reqPerSec))

	// add to channel on each tick
	go func() {
		for t := range t.C {
			select {
			case c <- t: // add the tick to channel
			default: // default will be executed when the channel is already full
			}
		}
		close(c) // close channel when the ticker is stopped
	}()

	return c
}

//GetLimiter gets a limiter from a map of limiters
func GetLimiter(limiters map[string]chan time.Time, id string) (limiter chan time.Time, res bool) {
	limiter, exists := limiters[id]

	if !exists {
		res = false
		return nil, res
	}

	return limiter, true
}

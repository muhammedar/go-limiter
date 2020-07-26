package limiter

import (
	"time"
)

// BuildLimiter is responsible for building a channel whcich will be used to store and pass the rquests
func BuildLimiter(reqPerSec, buffer int) (c chan time.Time) {

	// building a the buffer using go-channels, and initializing it with time.now
	c = make(chan time.Time, buffer)
	for i := 0; i < buffer; i++ {
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

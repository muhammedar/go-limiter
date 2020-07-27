package limiter

import (
	"time"
)

//Limiter is the body of the limiter object
type Limiter struct {
	//ID of the limiter
	ID string
	// Channel is the channel that manages the request limiting
	Channel chan time.Time
}

//NewLimiterObject creats a new limiter object
func NewLimiterObject(id string) *Limiter {
	return &Limiter{
		ID:      id,
		Channel: make(chan time.Time),
	}
}

// BuildChannel is responsible for building a channel which will be used to store and pass the requests
func (l *Limiter) BuildChannel(reqPerSec, bufferSize int) (c chan time.Time) {
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
func GetLimiter(limiters map[string]*Limiter, id string) (limiter *Limiter, res bool) {
	limiter, exists := limiters[id]

	if !exists {
		res = false
		return nil, res
	}

	return limiter, true
}

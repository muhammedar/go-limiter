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

// BuildLimiter is responsible for building a channel which will be used to store and pass the requests
func (l *Limiter) BuildLimiter(id string, reqPerSec, bufferSize int) (limiter *Limiter) {
	limiter.ID = id
	// building a the buffer using go-channels, and initializing it with time.now
	limiter.Channel = make(chan time.Time, bufferSize)
	for i := 0; i < bufferSize; i++ {
		limiter.Channel <- time.Now()
	}

	// ticker will help organizing the time accourding to the needed requests per seconds
	t := time.NewTicker(time.Second / time.Duration(reqPerSec))

	// add to channel on each tick
	go func() {
		for t := range t.C {
			select {
			case limiter.Channel <- t: // add the tick to channel
			default: // default will be executed when the channel is already full
			}
		}
		close(limiter.Channel) // close channel when the ticker is stopped
	}()

	return limiter
}

//GetLimiter gets a limiter from a map of limiters
func (l *Limiter) GetLimiter(limiters map[string]*Limiter, id string) (limiter *Limiter, res bool) {
	limiter, exists := limiters[id]

	if !exists {
		res = false
		return nil, res
	}

	return limiter, true
}

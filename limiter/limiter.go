package limiter

import "time"

//NewLimitWindow build the Queue which will hold the messeges
func NewLimitWindow(reqPerSec int) *LimitWindow {
	return &LimitWindow{
		ReqPerSec: reqPerSec,
		Queue:     make([]*Message, 0, reqPerSec),
	}
}

//Check is main (use time.Duration)
func (l *LimitWindow) Check() time.Duration {
	mutex.Lock()
	defer mutex.Unlock()
	m := NewMessage()
	full := l.checkSize()
	if !full {
		l.push(m)
	} else {
		l.remove()
		l.push(m)
		sleepTime := l.calculateSleepTime()
		return sleepTime
	}
	return 0
}

package limiter

import "time"

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

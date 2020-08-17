package limiter

import "time"

//NewLimitWindow build the Queue which will hold the messeges
func NewLimitWindow(reqPerSec int) *LimitWindow {
	return &LimitWindow{
		ReqPerSec: reqPerSec,
		Queue:     make([]time.Time, 0, reqPerSec),
		Debug:     false,
	}
}

//Check is used to check the queue and decide the time needed to sleep (use time.Duration)
func (l *LimitWindow) Check() time.Duration {
	l.Mutex.Lock()
	defer l.Mutex.Unlock()
	m := time.Now()
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

//CheckWithSleep is used to check the queue and decide the time needed to sleep (use time.Duration)
func (l *LimitWindow) CheckWithSleep() {
	l.Mutex.Lock()
	defer l.Mutex.Unlock()
	m := time.Now()
	full := l.checkSize()
	if !full {
		l.push(m)
	} else {
		l.remove()
		l.push(m)
		sleepTime := l.calculateSleepTime()
		time.Sleep(sleepTime)
	}
}

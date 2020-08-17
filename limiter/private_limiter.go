package limiter

import (
	"log"
	"sync"
	"time"
)

//LimitWindow is the Queue that holds the message object
type LimitWindow struct {
	ReqPerSec int
	Queue     []time.Time
	Mutex     sync.Mutex
	Debug     bool //default is false, set to true to see logs
}

func (l *LimitWindow) debugLogs(msg string, a ...interface{}) {
	if l.Debug {
		log.Printf(msg, a...)
	}
}

//checkSize checks if the window (queue) is full
func (l *LimitWindow) checkSize() bool {
	if len(l.Queue) == l.ReqPerSec {
		l.debugLogs("Rate Limiter: Queue is full")
		return true
	}
	return false
}

//push appends to the queue
func (l *LimitWindow) push(m time.Time) {
	l.Queue = append(l.Queue, m)
	var txt string
	for _, m := range l.Queue {
		txt += m.Format(time.StampMilli) + ", "
	}
	l.debugLogs("[%v]", txt)
}

//calaculates the amout of sleep time needed
func (l *LimitWindow) calculateSleepTime() (t time.Duration) {
	//here i need to sleep and remove the first element
	x := l.Queue[len(l.Queue)-1]
	y := l.Queue[0]
	ans := x.Sub(y)
	if ans < time.Second {
		l.debugLogs("%v", ans)
		t = time.Second - ans
		return t
	}
	return 0
}

//remove is removing the first element from a queue index 0
func (l *LimitWindow) remove() {
	l.Queue = l.Queue[1:]
}

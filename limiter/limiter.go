package main

import (
	"log"
	"sync"
	"time"
)

var mutex sync.Mutex

//Message represents a message request
type Message struct {
	Time time.Time
}

//LimitWindow is the Queue that holds the message object
type LimitWindow struct {
	ReqPerSec int
	Queue     []*Message
}

//NewMessage creates a new message object
func NewMessage() *Message {
	return &Message{
		Time: time.Now(),
	}
}

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

func (l *LimitWindow) checkSize() bool {
	if len(l.Queue) == l.ReqPerSec {
		log.Println("Queue is full, can't handle more, about to decide how much time to wait...")

		return true
	}
	return false
}

func (l *LimitWindow) push(m *Message) {
	l.Queue = append(l.Queue, m)
	var txt string
	for _, m := range l.Queue {
		txt += m.Time.Format(time.StampMilli) + ", "
	}
	log.Printf("[%v]", txt)
}

func (l *LimitWindow) calculateSleepTime() (t time.Duration) {
	//here i need to sleep and remove the first element
	x := l.Queue[len(l.Queue)-1].Time
	y := l.Queue[0].Time
	ans := x.Sub(y)
	if ans < time.Second {
		log.Printf("%v", ans)
		t = time.Second
		return t
	}
	return 0
}

func (l *LimitWindow) remove() {
	l.Queue = l.Queue[1:]
}

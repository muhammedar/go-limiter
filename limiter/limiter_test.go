package limiter

import (
	"log"
	"sync"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestNewLimitWindow(t *testing.T) {
	lw := NewLimitWindow(2)
	started := time.Now()
	var waitGroup sync.WaitGroup
	waitGroup.Add(5)
	for i := 0; i < 5; i++ {
		go func() {
			defer waitGroup.Done()
			for i := 0; i < 2; i++ {
				st := lw.Check()
				dur := time.Duration(st)
				{
					mutex.Lock()
					time.Sleep(dur)
					mutex.Unlock()
				}
				log.Printf("going to sleep: %v", dur)
			}
		}()
	}
	waitGroup.Wait()
	log.Println(time.Since(started))
	assert.True(t, time.Since(started) > time.Second*4)
}

func TestNewLimitWindow_1(t *testing.T) {
	lw := NewLimitWindow(10)
	started := time.Now()
	for i := 0; i < 50; i++ {
		dur := lw.Check()
		log.Printf("going to sleep: %v", dur)
		time.Sleep(dur)
	}
	assert.True(t, time.Since(started) > time.Second*3)
	log.Println(time.Since(started))
}

func TestNewLimitWindow_3(t *testing.T) {
	lw := NewLimitWindow(10)
	started := time.Now()
	var waitGroup sync.WaitGroup
	waitGroup.Add(15)
	for i := 0; i < 15; i++ {
		go func() {
			defer waitGroup.Done()
			st := lw.Check()
			log.Printf("going to sleep: %v", st)
			{
				mutex.Lock()
				time.Sleep(st)
				mutex.Unlock()
			}
		}()
	}
	waitGroup.Wait()
	log.Println(time.Since(started))
	assert.True(t, time.Since(started) >= time.Second)
}

func TestNewLimitWindow_4(t *testing.T) {
	lw := NewLimitWindow(3)
	started := time.Now()

	st := lw.Check()
	log.Printf("going to sleep: %v", st)
	time.Sleep(st)
	st = lw.Check()
	log.Printf("going to sleep: %v", st)
	time.Sleep(st)
	st = lw.Check()
	log.Printf("going to sleep: %v", st)
	time.Sleep(st)
	time.Sleep(time.Second)
	st = lw.Check()
	log.Printf("going to sleep: %v", st)
	time.Sleep(st)
	st = lw.Check()
	log.Printf("going to sleep: %v", st)
	time.Sleep(st)
	st = lw.Check()
	log.Printf("going to sleep: %v", st)
	time.Sleep(st)
	st = lw.Check()
	log.Printf("going to sleep: %v", st)
	time.Sleep(st)
	st = lw.Check()
	log.Printf("going to sleep: %v", st)
	time.Sleep(st)
	st = lw.Check()
	log.Printf("going to sleep: %v", st)
	time.Sleep(st)
	st = lw.Check()
	log.Printf("going to sleep: %v", st)
	time.Sleep(st)
	st = lw.Check()
	log.Printf("going to sleep: %v", st)
	time.Sleep(st)
	st = lw.Check()
	log.Printf("going to sleep: %v", st)
	time.Sleep(st)
	st = lw.Check()
	log.Printf("going to sleep: %v", st)
	time.Sleep(st)
	st = lw.Check()
	log.Printf("going to sleep: %v", st)
	time.Sleep(st)

	log.Println(time.Since(started))
	assert.True(t, time.Since(started) >= time.Second*4)
}

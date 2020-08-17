package limiter

import (
	"fmt"
	"log"
	"sync"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestNewLimitWindow(t *testing.T) {
	lw := NewLimitWindow(2)
	lw.Debug = false
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
					lw.Mutex.Lock()
					time.Sleep(dur)
					lw.Mutex.Unlock()
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
				lw.Mutex.Lock()
				time.Sleep(st)
				lw.Mutex.Unlock()
			}
		}()
	}
	waitGroup.Wait()
	log.Println(time.Since(started))
	assert.True(t, time.Since(started) >= time.Second)
}

func TestNewLimitWindow_4(t *testing.T) {
	lw := NewLimitWindow(3)
	lw.Debug = true
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

func TestNewLimitWindow_6(t *testing.T) {
	start := time.Now()
	count := 150
	lw := NewLimitWindow(5)
	for i := 0; i < count; i++ {
		lw.CheckWithSleep()
	}
	total := time.Since(start)
	fmt.Printf("total: %v messages in %v  (rate: %v)\n", count, total, float64(count)/total.Seconds())

}

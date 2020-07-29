package main

import (
	"log"
	"sync"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestNewLimitWindow(t *testing.T) {
	lw := NewLimitWindow(10)
	started := time.Now()
	var waitGroup sync.WaitGroup
	waitGroup.Add(6)
	for i := 0; i < 6; i++ {
		go func() {
			defer waitGroup.Done()
			for j := 0; j < 10; j++ {
				time.Sleep(1)
				st := lw.Check()
				dur := time.Second * time.Duration(st)
				log.Printf("going to sleep: %v", dur)
			}
		}()
	}
	waitGroup.Wait()
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
	assert.True(t, time.Since(started) < time.Second*3)

	// time.Sleep(time.Second)
	// st := lw.Check()
	// dur := time.Second * time.Duration(st)
	// log.Printf("going to sleep: %v", dur)
	// time.Sleep(dur)

	log.Println(time.Since(started))
}

func TestNewLimitWindow_3(t *testing.T) {
	lw := NewLimitWindow(3)
	started := time.Now()
	var waitGroup sync.WaitGroup
	waitGroup.Add(15)
	for i := 0; i < 15; i++ {
		go func() {
			defer waitGroup.Done()
			st := lw.Check()
			log.Printf("going to sleep: %v", st)
			time.Sleep(st)
		}()
	}
	waitGroup.Wait()
	log.Println(time.Since(started))
	assert.True(t, time.Since(started) > time.Second*2)
}

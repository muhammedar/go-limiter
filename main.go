package main

import "github.com/muhammedGabriel/limiter"

const (
	rps        = 50000
	bufferSize = 10
)

func main() {
	limiter.RateLimit(rps, bufferSize, 50000)
}

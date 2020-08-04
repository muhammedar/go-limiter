package limiter

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestBuildChannel(t *testing.T) {

	l := BuildChannel(3, 3)
	mockRequests := make(chan int, 10)
	for i := 1; i <= 10; i++ {
		mockRequests <- i
	}
	close(mockRequests)
	for req := range mockRequests {
		<-l
		fmt.Println("request", req, time.Now())

	}
	assert.NotEmpty(t, l)

}

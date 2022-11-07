package throt

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestLimiter(t *testing.T) {
	target := "NewLimiter"
	count := 3

	t.Logf("Given the need to test %s.", target)
	l := NewLimiter(time.Duration(time.Second), count)
	lim, ok := l.(*limiter)
	t.Logf("\t%s\tAfter calling NewLimiter()", succeed)
	assert.True(t, ok)
	assert.Equal(t, lim.maxCount, count)
	t.Logf("\t%s\tShould have max count set.", succeed)
}

func TestAwait(t *testing.T) {
	var result bool
	target := "Await"
	count := 3   // requests.
	seconds := 2 // per seconds.
	workers := 5
	var total int

	t.Logf("Given the need to test %s.", target)

	// Cancel after 5 sec.
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(5)*time.Second)
	defer cancel() // cancel after 5 sec.

	// a channel with guarantee for wait-for-result.
	ch := make(chan int)

	// Three requests per sec.
	l := NewLimiter(time.Duration(time.Duration(seconds)*time.Second), count)

	// Worker goroutines.
	for i := 0; i < workers; i++ {
		go func(ctx context.Context) {
			for {
				select {
				case <-ctx.Done():
					return
				default:
					l.Await() // limit the throttle.
					ch <- 1
				}
			}
		}(ctx)
	}

	ticker := time.NewTicker(time.Duration(1800) * time.Millisecond)
	i := 1
	for {
		// Wait for %count% workers to complete, checking the number.
		select {
		case num := <-ch:
			total += num
		case <-ctx.Done():
			switch {
			case i < 3:
				result = assert.Fail(t, "early cancellation")
			case i == 3:
				result = assert.Equal(t, count*i, total, fmt.Sprintf("expect workers to compete %d requests in %d iterations", count*i, i))
			default:
				result = assert.Fail(t, "late cancellation")
			}
			t.Logf(fmt.Sprintf("\t%s\tShould complete %d requests in %d iterations.", success(result), count*i, i))
			cancel()
			return
		case <-ticker.C:
			result = assert.Equal(t, count*i, total, fmt.Sprintf("expect workers to compete %d requests in %d iterations", count*i, i))
			t.Logf(fmt.Sprintf("\t%s\tShould complete %d requests in %d iterations.", success(result), count*i, i))
			i++
		}
	}
}

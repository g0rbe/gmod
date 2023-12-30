package slitu

import (
	"context"
	"time"
)

// Sleep is a context aware sleep function, which is sleep for d duration or return if ctx is cancelled.
func Sleep(ctx context.Context, d time.Duration) {

	select {
	case <-ctx.Done():
		return
	case <-time.After(d):
		return
	}
}

/*
Blocklist module provides a basic blocklisting mechanism:

- Add remote IP as a string with Block() and BlockTime()
- Check the status of the remote with IsBlocked().
- Deleting expired entries are happening automatically in the background.
*/
package blocklist

import (
	"sync"
	"time"
)

type Blocklist struct {
	blocklist map[string]int64

	maxLen int64 // Maximum length of the map, after exceeding this number, GarbageCollector() will run

	duration time.Duration // Default block time

	m *sync.RWMutex
}

// NewBlocklist creates a new Blocklist.
// t sets the default block time, l set the max length for the blocklist.
//
// If the size of the blocklist exceed l,
// an expensive garbage collector removes every expired entries.
// If l is 0, the garbage collector will never run.
func NewBlocklist(t time.Duration, l int64) *Blocklist {

	v := new(Blocklist)

	v.blocklist = make(map[string]int64, l)
	v.maxLen = l
	v.duration = t
	v.m = new(sync.RWMutex)

	return v
}

// BlockTime blocks ip for for time t (now + t).
//
// If the underlying size of the blocklist reaches the limit,
// run the GarbageCollector() before adding the new entry.
func (b *Blocklist) BlockTime(ip string, t time.Duration) {

	b.GarbageCollector()

	b.m.Lock()

	b.blocklist[ip] = time.Now().Add(b.duration).UnixMilli()

	b.m.Unlock()
}

// Block blocks ip using the default block time.
//
// If the underlying size of the blocklist reaches the limit,
// run the GarbageCollector() before adding the new entry.
func (b *Blocklist) Block(ip string) {

	b.BlockTime(ip, b.duration)
}

// Get returns the block end time in unix millisecond fir ip
// and a boolean to indicate if ip set.
func (b *Blocklist) Get(ip string) (int64, bool) {

	b.m.RLock()

	t, ok := b.blocklist[ip]

	b.m.RUnlock()

	return t, ok
}

// GetTime returns the block end time in unix millisecond.
func (b *Blocklist) GetTime(ip string) int64 {

	t, _ := b.Get(ip)

	return t
}

// GetRemaining returns the reaining millisecond from the block.
func (b *Blocklist) GetRemaining(ip string) int64 {

	t, _ := b.Get(ip)

	t = t - time.Now().UnixMilli()

	if t < 0 {
		return 0
	}

	return t
}

// GetLen returns the length of the underlying map.
func (b *Blocklist) GetLen() int64 {

	b.m.RLock()
	l := len(b.blocklist)
	b.m.RUnlock()

	return int64(l)
}

// Remove removes ip from the blocklist.
func (b *Blocklist) Remove(ip string) {

	b.m.Lock()

	delete(b.blocklist, ip)

	b.m.Unlock()
}

// IsBlocked returns whether ip is blocked.
func (b *Blocklist) IsBlocked(ip string) bool {

	t, set := b.Get(ip)

	if !set {
		return false
	}

	if t < time.Now().UnixMilli() {
		b.Remove(ip)
		return false
	}

	return true
}

// GarbageCollector removes the expired IPs from the blocklist
// if the size of the list exceed the limit set in NewBlocklist().
func (b *Blocklist) GarbageCollector() {

	// Do nothing if the limit is 0 or size not exceed the limit
	if b.maxLen == 0 || b.GetLen() < b.maxLen {
		return
	}

	t := time.Now().UnixMilli()

	b.m.Lock()

	for k, v := range b.blocklist {

		if v < t {
			delete(b.blocklist, k)
		}
	}

	b.m.Unlock()
}

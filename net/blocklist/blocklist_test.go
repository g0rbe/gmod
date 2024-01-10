package blocklist

import (
	"math/rand"
	"testing"
	"time"

	"github.com/g0rbe/gmod/net/ip"
)

func TestBlocklist(t *testing.T) {

	c := "0.0.0.0"

	b := NewBlocklist(1*time.Second, 0)

	b.Block(c)

	now := time.Now()

	for b.IsBlocked(c) {
		time.Sleep(1 * time.Millisecond)
	}

	dif := time.Since(now).Milliseconds() - 1000
	if dif > 10 {
		t.Fatalf("Took to much: %d millisecond\n", dif)
	} else {
		t.Logf("Difference: %d\n", dif)
	}
}

func TestBlocklistRace(t *testing.T) {

	c := "0.0.0.0"

	b := NewBlocklist(1*time.Millisecond, 1)

	go func() {

		for i := 0; i < 5000000; i++ {

			switch rand.Intn(9) {
			case 0:
				b.BlockTime(c, 10)
			case 1:
				b.Block(c)
			case 2:
				b.Get(c)
			case 3:
				b.GetTime(c)
			case 4:
				b.GetRemaining(c)
			case 5:
				b.GetLen()
			case 6:
				b.Remove(c)
			case 7:
				b.IsBlocked(c)
			case 8:
				b.GarbageCollector()
			}
		}
	}()

	for i := 0; i < 6000000; i++ {

		switch rand.Intn(9) {
		case 0:
			b.BlockTime(c, 10)
		case 1:
			b.Block(c)
		case 2:
			b.Get(c)
		case 3:
			b.GetTime(c)
		case 4:
			b.GetRemaining(c)
		case 5:
			b.GetLen()
		case 6:
			b.Remove(c)
		case 7:
			b.IsBlocked(c)
		case 8:
			b.GarbageCollector()
		}
	}
}

func BenchmarkBlockTime(b *testing.B) {

	list := NewBlocklist(20*time.Second, int64(b.N))

	ips := make([]string, 0, b.N)

	for i := 0; i < b.N; i++ {
		ips = append(ips, ip.GetRandom().String())
	}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		list.BlockTime(ips[i], 1000)
	}
}

func BenchmarkBlock(b *testing.B) {

	list := NewBlocklist(20*time.Second, int64(b.N))

	ips := make([]string, 0, b.N)

	for i := 0; i < b.N; i++ {
		ips = append(ips, ip.GetRandom().String())
	}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		list.Block(ips[i])
	}
}

func BenchmarkGet(b *testing.B) {

	list := NewBlocklist(20*time.Second, int64(b.N))

	ips := make([]string, 0, b.N)

	for i := 0; i < b.N; i++ {

		v := ip.GetRandom().String()

		ips = append(ips, v)
		list.Block(v)
	}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {

		list.Get(ips[i])
	}
}

func BenchmarkGetTime(b *testing.B) {

	list := NewBlocklist(20*time.Second, int64(b.N))

	ips := make([]string, 0, b.N)

	for i := 0; i < b.N; i++ {

		v := ip.GetRandom().String()

		ips = append(ips, v)
		list.Block(v)
	}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {

		list.GetTime(ips[i])
	}
}

func BenchmarkGetRemaining(b *testing.B) {

	list := NewBlocklist(20*time.Second, int64(b.N))

	ips := make([]string, 0, b.N)

	for i := 0; i < b.N; i++ {

		v := ip.GetRandom().String()

		ips = append(ips, v)
		list.Block(v)
	}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {

		list.GetRemaining(ips[i])
	}
}

func BenchmarkGetLen(b *testing.B) {

	list := NewBlocklist(20*time.Second, int64(b.N))

	for i := 0; i < b.N; i++ {

		list.Block(ip.GetRandom().String())
	}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {

		list.GetLen()
	}
}

func BenchmarkRemove(b *testing.B) {

	list := NewBlocklist(20*time.Second, int64(b.N))

	ips := make([]string, 0, b.N)

	for i := 0; i < b.N; i++ {

		v := ip.GetRandom().String()

		ips = append(ips, v)
		list.Block(v)
	}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {

		list.Remove(ips[i])
	}
}

func BenchmarkBlocked(b *testing.B) {

	list := NewBlocklist(20, 0)

	ips := make([]string, 0, b.N)

	for i := 0; i < b.N; i++ {

		v := ip.GetRandom().String()

		ips = append(ips, v)
		list.Block(v)
	}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		list.IsBlocked(ips[i])
	}
}

func BenchmarkGarbageCollector(b *testing.B) {

	lists := make([]*Blocklist, 0, b.N)

	for i := 0; i < b.N; i++ {

		l := NewBlocklist(10, 1000)
		for j := 0; j < 1000; j++ {
			l.Block(ip.GetRandom().String())
		}

		lists = append(lists, l)
	}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		lists[i].GarbageCollector()
	}
}

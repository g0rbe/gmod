package ctlog

import (
	"strings"
	"sync/atomic"
	"testing"
)

func TestLogByName(t *testing.T) {

	l := LogByName("argon2022")
	if l == nil {
		t.Fatalf("FAIL: argon2022 not found\n")
	}

	t.Logf("%#v\n", l)
}

func TestSize(t *testing.T) {

	for i := range Logs {

		if i > 9 {
			break
		}

		size, err := Size(Logs[i].URI)
		if err != nil {
			t.Fatalf("Size of %s failed: %s\n", Logs[i].Name, err)
		}

		t.Logf("Size of %s: %d\n", Logs[i].Name, size)
	}
}

func TestMaxBatchSize(t *testing.T) {

	for i := range Logs {

		if i > 9 {
			break
		}

		size, err := MaxBatchSize(Logs[i].URI)
		if err != nil && !strings.Contains(err.Error(), "429 Too Many Requests") {
			t.Fatalf("BatchSize of %s failed: %s\n", Logs[i].Name, err)
		}

		t.Logf("BatchSize of %s: %d\n", Logs[i].Name, size)
	}
}

func TestGetDomains(t *testing.T) {

	index := new(atomic.Int64)
	total := 0

	for {

		if index.Load() > 10000 {
			break
		}

		t.Logf("index: %d, total: %d\n", index.Load(), total)

		r, n, err := GetDomains("https://ct.googleapis.com/logs/argon2022/", index.Load())
		if err != nil {
			t.Fatalf("FAIL: %s\n", err)
		}

		index.Add(n)

		total += len(r)

	}

	t.Logf("index: %d, total: %d\n", index.Load(), total)

}

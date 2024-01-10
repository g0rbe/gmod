package hetzner

import (
	"errors"
	"os"
	"testing"
	"time"
)

func TestGetAllRecords(t *testing.T) {

	c := NewClientWithTimeout(os.Getenv("HETZNER_DNS_KEY"), 5*time.Second)

	zs, err := c.GetAllRecords()
	if err != nil {
		t.Fatalf("FAIL: %s\n", err)
	}

	t.Logf("%#v\n", zs)
}

func TestGetAllRecordsNoAPIKey(t *testing.T) {

	c := NewClientWithTimeout("", 5*time.Second)

	zs, err := c.GetAllRecords()

	if err == nil {
		t.Fatalf("FAIL: error is nil, returned: %#v\n", zs)
	}

	if !errors.Is(err, ErrNoAPIKey) {
		t.Fatalf("FAIL: error got: %s, want: %s\n", err, ErrNoAPIKey)
	}
}

func TestGetAllRecordsInvalidAPIKey(t *testing.T) {

	c := NewClientWithTimeout("invalid", 5*time.Second)

	zs, err := c.GetAllRecords()

	if err == nil {
		t.Fatalf("FAIL: error is nil, returned: %#v\n", zs)
	}

	if !errors.Is(err, ErrInvalidAPIKey) {
		t.Fatalf("FAIL: error got: %s, want: %s\n", err, ErrInvalidAPIKey)
	}
}

func TestGetAllRecordsByZoneZoneNotFound(t *testing.T) {

	c := NewClientWithTimeout(os.Getenv("HETZNER_DNS_KEY"), 5*time.Second)

	zs, err := c.GetAllRecordsByZone("notexists")

	if err == nil {
		t.Fatalf("FAIL: error is nil, returned: %#v\n", zs)
	}

	if !errors.Is(err, ErrZoneNotFound) {
		t.Fatalf("FAIL: error got: %s, want: %s\n", err, ErrNoAPIKey)
	}
}

func TestGetAllRecordsByZoneNoAPIKey(t *testing.T) {

	c := NewClientWithTimeout("", 5*time.Second)

	zs, err := c.GetAllRecordsByZone("invalid")

	if err == nil {
		t.Fatalf("FAIL: error is nil, returned: %#v\n", zs)
	}

	if !errors.Is(err, ErrNoAPIKey) {
		t.Fatalf("FAIL: error got: %s, want: %s\n", err, ErrInvalidAPIKey)
	}
}

func TestGetAllRecordsByZoneInvalidAPIKey(t *testing.T) {

	c := NewClientWithTimeout("invalid", 5*time.Second)

	zs, err := c.GetAllRecordsByZone("invalid")

	if err == nil {
		t.Fatalf("FAIL: error is nil, returned: %#v\n", zs)
	}

	if !errors.Is(err, ErrInvalidAPIKey) {
		t.Fatalf("FAIL: error got: %s, want: %s\n", err, ErrInvalidAPIKey)
	}
}

package pokecache

import (
	"bytes"
	"testing"
	"time"
)

func TestAddAndGet(t *testing.T) {
	cache := NewCache(time.Hour)
	cache.Add("first", []byte("one"))
	cache.Add("second", []byte("two"))

	for key, want := range map[string][]byte{
		"first":  []byte("one"),
		"second": []byte("two"),
	} {
		got, ok := cache.Get(key)
		if !ok || !bytes.Equal(got, want) {
			t.Errorf("Get(%q) = %q, %v; want %q, true", key, got, ok, want)
		}
	}
}

func TestGetMissing(t *testing.T) {
	cache := NewCache(time.Hour)
	got, ok := cache.Get("missing")
	if ok || got != nil {
		t.Errorf("Get(missing) = %q, %v; want nil, false", got, ok)
	}
}

func TestFreshEntry(t *testing.T) {
	cache := NewCache(time.Second)
	cache.Add("fresh", []byte("value"))

	got, ok := cache.Get("fresh")
	if !ok || !bytes.Equal(got, []byte("value")) {
		t.Errorf("Get(fresh) = %q, %v; want value, true", got, ok)
	}
}

func TestExpiredEntry(t *testing.T) {
	cache := NewCache(20 * time.Millisecond)
	cache.Add("old", []byte("value"))

	deadline := time.Now().Add(time.Second)
	for time.Now().Before(deadline) {
		if _, ok := cache.Get("old"); !ok {
			return
		}
		time.Sleep(5 * time.Millisecond)
	}
	t.Error("old entry was not removed after expiration")
}

func TestReplace(t *testing.T) {
	cache := NewCache(time.Hour)
	cache.Add("key", []byte("old"))
	cache.Add("other", []byte("unchanged"))
	cache.Add("key", []byte("new"))

	for key, want := range map[string][]byte{
		"key":   []byte("new"),
		"other": []byte("unchanged"),
	} {
		got, ok := cache.Get(key)
		if !ok || !bytes.Equal(got, want) {
			t.Errorf("Get(%q) = %q, %v; want %q, true", key, got, ok, want)
		}
	}
}

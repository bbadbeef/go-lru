package cache

import (
	"testing"
)

func TestLRU(t *testing.T) {
	c := NewCache(5)
	c.Set("1", "1")
	c.Set("2", "1")
	c.Set("3", "1")
	c.Set("4", "1")
	c.Set("5", "1")
	c.Set("6", "1")
	c.Set("7", "1")
	if c.Get("1") != nil {
		t.Fatal("get 2 is not nil")
	}

	c.Get("3")
	c.Set("8", "1")
	if c.Get("3") == nil {
		t.Fatal("get 3 is nil")
	}
}

package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCacheGetSet(t *testing.T) {
	testCases := []struct {
		description string
		key      string
		keyToGet string
		value    string
		ok       bool
		want     string
	}{
		{"just suceed to set and retrieve key", "key1", "key1", "value", true, "value"},
		{"should fail because key dont exist", "key2", "wrong-key", "value", false, ""},
	}
	c := NewCache()

	for _, tt := range testCases {
		t.Run(tt.description, func(t *testing.T) {
			t.Parallel()
			c.Set(tt.key, tt.value)

			val, ok := c.Get(tt.keyToGet)
			assert.Equal(t, tt.ok, ok)
			assert.Equal(t, tt.want, val)
		})
	}
}

package urna

import (
	"slices"
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
		{"just suceed to s and retrieve key", "key1", "key1", "value", true, "value"},
		{"should fail because key dont exist", "key2", "wrong-key", "value", false, ""},
	}
	c := NewCache()

	for _, tt := range testCases {
		t.Run(tt.description, func(t *testing.T) {
			c.Set(tt.key, tt.value)

			val, ok := c.Get(tt.keyToGet)
			assert.Equal(t, tt.ok, ok)
			assert.Equal(t, tt.want, val)
		})
	}
}

func TestCacheSetMultiple(t *testing.T) {
	c := NewCache()
	s := map[string]string{
			"key": "value",
			"key2": "value2", 
	}
	assert.Equal(t, c.Len(), 0)

	c.SetMultiple(s)

	assert.Equal(t, c.Len(), 2)
}

func TestCacheGetMultiple(t *testing.T) {
	c := NewCache()

	s := map[string]string{
			"key": "value",
			"key2": "value2", 
	}
	assert.Equal(t, c.Len(), 0)

	c.SetMultiple(s)

	assert.Equal(t, c.Len(), 2)

	results := c.GetMultiple([]string{"key", "key2"})


	for _, v := range s {
		assert.Equal(t, true, slices.Contains(results, v))
	}
}

func TestCacheClear(t *testing.T) {
	c := NewCache()

	c.Set("key", "value")
	c.Set("key2", "value2")

	assert.Equal(t, c.Len(), 2)

	c.Clear()
	assert.Equal(t, c.Len(), 0)
}

func TestCacheDelete(t *testing.T) {

	c := NewCache()

	c.Set("key", "value")
	c.Set("key2", "value2")
	assert.Equal(t, c.Len(), 2)

	c.Delete("key2")
	assert.Equal(t, c.Len(), 1)
}

func TestCacheLen(t *testing.T) {
	c := NewCache()


	assert.Equal(t, c.Len(), 0)

	c.Set("key", "value")

	assert.Equal(t, c.Len(), 1)
}

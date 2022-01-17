package redis

import (
	"testing"
	"time"

	"github.com/alicebob/miniredis/v2"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var extension = new(REDIS)

func TestRedis(t *testing.T) {
	mr, err := miniredis.Run()
	require.NoError(t, err)
	defer mr.Close()

	client := extension.NewClient(mr.Addr(), "", 0)

	// Set and Get:
	extension.Set(client, "foo", "bar", 0)
	gets := extension.Get(client, "foo")
	if gets != "bar" {
		t.Fatal("'bar' should have been returned")
	}

	// Del:
	extension.Del(client, "foo")
	if mr.Exists("foo") {
		t.Fatal("'foo' should not have existed anymore")
	}

	// Custom command:
	result, err := extension.Do(client, "PING")
	require.NoError(t, err)
	assert.Equal(t, "PONG", result)

	// Custom command:
	result, err = extension.Do(client, "SADD", "foo", "bar")
	require.NoError(t, err)
	assert.Equal(t, int64(1), result)

	// TTL and expiration:
	extension.Set(client, "foo", "bar", 5)
	mr.FastForward(10 * time.Second)
	if mr.Exists("foo") {
		t.Fatal("'foo' should not have existed anymore")
	}
}

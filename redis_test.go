package redis

import (
	"testing"
	"time"

	"github.com/alicebob/miniredis/v2"
)

var plugin = new(REDIS)

func TestRedis(t *testing.T) {
	mr, err := miniredis.Run()
	if err != nil {
		panic(err)
	}
	defer mr.Close()

	client := plugin.NewClient(mr.Addr(), "", 0)

	// Set and Get:
	plugin.Set(client, "foo", "bar", 0)
	gets := plugin.Get(client, "foo")
	if gets != "bar" {
		t.Fatal("'bar' should have been returned")
	}

	// Del:
	plugin.Del(client, "foo")
	if mr.Exists("foo") {
		t.Fatal("'foo' should not have existed anymore")
	}

	// Custom command:
	gets = plugin.Do(client, "PING", "")
	if gets != "PONG" {
		t.Fatal("'PONG' should have been returned")
	}

	// TTL and expiration:
	plugin.Set(client, "foo", "bar", 5)
	mr.FastForward(10 * time.Second)
	if mr.Exists("foo") {
		t.Fatal("'foo' should not have existed anymore")
	}
}

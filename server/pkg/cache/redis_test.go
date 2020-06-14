package cache

import (
	"github.com/shiniao/gtodo/pkg/redis"
	"reflect"
	"testing"
	"time"
)

func TestRedisCacheGetSet(t *testing.T) {
	redis.InitTestRedis()

	redisclient := redis.Client
	cache := NewRedisCache(redisclient, "unit-test", JSONEncoding{})

	// test set

	type setArgs struct {
		key   string
		value interface{}
		exp   time.Duration
	}

	setTests := []struct {
		name    string
		cache   Driver
		args    setArgs
		wantErr bool
	}{
		{
			"test redis set",
			cache, // redis cache
			setArgs{"key-001", "value-001", 60 * time.Second},
			false,
		},
	}

	for _, test := range setTests {
		t.Run(test.name, func(t *testing.T) {
			c := test.cache
			err := c.Set(test.args.key, test.args.value, test.args.exp)
			if (err != nil) != test.wantErr {
				t.Errorf("Set() error = %v, wantErr %v", err, test.wantErr)
			}
		})
	}

	// test get
	type args struct {
		key string
	}

	tests := []struct {
		name    string
		cache   Driver
		args    args
		wantVal interface{}
		wantErr bool
	}{
		{
			"test redis get",
			cache,
			args{"key-001"},
			"value-001",
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := tt.cache
			var gotVal interface{}
			err := c.Get(tt.args.key, &gotVal)
			if (err != nil) != tt.wantErr {
				t.Errorf("Get() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			t.Log("gotval", gotVal)
			if !reflect.DeepEqual(gotVal, tt.wantVal) {
				t.Errorf("Get() gotVal = %v, want %v", gotVal, tt.wantVal)
			}
		})
	}
}

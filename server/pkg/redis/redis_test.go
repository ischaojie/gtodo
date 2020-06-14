package redis

import (
	"testing"
	"time"
)

func TestRedisInit(t *testing.T)  {
	InitTestRedis()

	err := Client.Ping().Err()
	if err != nil{
		t.Error("ping redis server err:", err)
		return
	}
	t.Log("ping redis server pass")
}

func TestRedisSetGet(t *testing.T){
	InitTestRedis()

	var key = "test-set"
	var value = "test-content"

	Client.Set(key, value, time.Second*100)

	expValue := Client.Get(key).Val()
	if expValue != value{
		t.Logf("expect %s, but get %s", expValue, value)
		return
	}
	t.Log("redis set get test pass")
}

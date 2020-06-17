// vcode_service 验证码服务接口
package service

import (
	"fmt"
	"github.com/pkg/errors"
	"github.com/shiniao/gtodo/pkg/log"
	"github.com/shiniao/gtodo/pkg/redis"
	"math/rand"
	"strconv"
	"time"
)

const (
	maxDurationTime = time.Minute // 过期时间60s
	vcodeRedisKey   = "app:login:vcode:%d"
)

type VCodeService interface {
	GenLoginVCode(phone int) (int, error) // 生成验证码
	CheckLoginVCode(phone, code int) bool // 检查验证码
	GetLoginVCode(phone int) (int, error) // 获取验证码
}

type vcodeSvc struct{}

// GenLoginVCode 生成验证码
func (s *vcodeSvc) GenLoginVCode(phone int) (int, error) {
	// 1. 生成随机数
	vcodeStr := fmt.Sprintf("%06v", rand.New(rand.NewSource(time.Now().UnixNano())).Int31n(1000000))
	// 2. 存入redis
	// 构造存入code的key，保持唯一
	key := fmt.Sprintf(vcodeRedisKey, phone)
	err := redis.Client.Set(key, vcodeStr, maxDurationTime).Err()
	if err != nil {
		return 0, errors.Wrapf(err, "generator login code err")
	}

	vcode, err := strconv.Atoi(vcodeStr)
	if err != nil {
		return 0, errors.Wrapf(err, "str convert int err")
	}
	return vcode, nil
}

// GetLoginVCode 获得校验码
func (s *vcodeSvc) GetLoginVCode(phone int) (int, error) {
	// 直接从redis里获取
	key := fmt.Sprintf(vcodeRedisKey, phone)
	vcode, err := redis.Client.Get(key).Result()
	if err == redis.Nil {
		return 0, nil
	} else if err != nil {
		return 0, errors.Wrap(err, "redis get login vcode err")
	}

	verifyCode, err := strconv.Atoi(vcode)
	if err != nil {
		return 0, errors.Wrap(err, "strconv err")
	}

	return verifyCode, nil
}

// CheckLoginVCode 检查验证码是否正确
func (s *vcodeSvc) CheckLoginVCode(phone, vcode int) bool {
	oldVCode, err := s.GetLoginVCode(phone)
	if err != nil {
		log.Warnf("[vcode_service] get verify code err, %v", err)
		return false
	}

	if vcode != oldVCode {
		return false
	}

	return true
}

func NewVCodeSvc() VCodeService {
	return &vcodeSvc{}
}

// VCodeSvc 验证码服务接口
var VCodeSvc = NewVCodeSvc()

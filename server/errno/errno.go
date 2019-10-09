package errno

import "fmt"

// * 自定义错误代码（code+message）
// * 这是展示给用户的
type Errno struct {
	Code int
	Message string
}

// * 实现Error接口，自定义错误类型
/*
type error interface {
	Error() string
}
*/

func (err Errno) Error() string {
	return err.Message
}

// * 定义错误
// * 这是展示给后台的
type Err struct {
	Code int
	Message string
	Err error
}

func (err *Err) Error() string {
	return fmt.Sprintf("Err - code: %d, message: %s, error: %s", err.Code, err.Message, err.Err)
}

// * 新建一个自定义错误（传入code+message, err）
func New(errno *Errno, err error) *Err {
	return &Err{
		Code:    errno.Code,
		Message: errno.Message,
		Err:     err,
	}
}

// * 解析定制的错误
func DecodeErr(err error) (int, string) {
	// * 没有错误返回ok
	if err == nil {
		return OK.Code, OK.Message
	}

	switch typed := err.(type) {
	case *Err:
		return typed.Code, typed.Message
	case *Errno:
		return typed.Code, typed.Message
	default:
	}

	return InternalServerError.Code, err.Error()
}
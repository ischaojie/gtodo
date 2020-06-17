/*定义错误代码*/

package errno

var (
	// Common errors
	OK                  = &Errno{Code: 0, Message: "OK"}
	InternalServerError = &Errno{Code: 10001, Message: "Internal server error."}
	ErrBind             = &Errno{Code: 10002, Message: "Error occurred while binding the request body to the struct."}
	ErrParam            = &Errno{Code: 10003, Message: "The params incorrect."}

	// user errors
	ErrUserNotFound = &Errno{Code: 20102, Message: "The user was not found."}

	ErrValidation            = &Errno{Code: 20001, Message: "Validation failed."}
	ErrTwicePasswordNotMatch = &Errno{Code: 20002, Message: "The twice password not match."}
	ErrDatabase              = &Errno{Code: 20003, Message: "Database error."}
	ErrToken                 = &Errno{Code: 20004, Message: "Error occurred while signing the JSON web token."}
	ErrRegister              = &Errno{Code: 20005, Message: "Register error."}
	ErrEmailOrPassword       = &Errno{Code: 20006, Message: "Email or Password error."}
	ErrVCode                 = &Errno{Code: 20007, Message: "验证码错误."}

	// user errors
	ErrEncrypt      = &Errno{Code: 20101, Message: "Error occurred while encrypting the user password."}
	ErrTokenInvalid = &Errno{Code: 20103, Message: "The token was invalid."}
	ErrKeyIncorrect = &Errno{Code: 20104, Message: "The key was incorrect"}
)

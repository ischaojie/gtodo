package auth

import "golang.org/x/crypto/bcrypt"

// Compare 比较 加密前后密码是否一致
func Compare(hpwd, pwd string) error {
	return bcrypt.CompareHashAndPassword([]byte(hpwd), []byte(pwd))
}

// Encrypt 加密密码文本
func Encrypt(pwd string) (string, error) {
	hpwd, err := bcrypt.GenerateFromPassword([]byte(pwd), bcrypt.DefaultCost)
	return string(hpwd), err
}

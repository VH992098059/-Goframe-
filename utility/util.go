package utility

import "github.com/gogf/gf/v2/crypto/gmd5"

// EncryptPassword 密码加密（md5+加密盐）
func EncryptPassword(password, salt string) string {
	return gmd5.MustEncryptString(gmd5.MustEncryptString(password) + gmd5.MustEncryptString(salt))
}

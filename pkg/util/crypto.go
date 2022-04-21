package util

import (
	"crypto/sha256"
	"fmt"

	"example/pkg/logging"
	"example/pkg/setting"
)

func GetSha256Code(pwd string) string {
	h := sha256.New()
	sec, err := setting.Cfg.GetSection("app")
	if err != nil {
		logging.Fatal(2, "Fail to get section 'app': %v", err)
	}
	// 加盐
	salt := sec.Key("JWT_SALT").String()
	h.Write([]byte(pwd + salt))
	return fmt.Sprintf("%x", h.Sum(nil))
}

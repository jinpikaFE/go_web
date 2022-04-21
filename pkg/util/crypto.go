package util

import (
	"crypto/sha256"
	"fmt"
	"log"

	"example/pkg/setting"
)

func GetSha256Code(pwd string) string {
	h := sha256.New()
	sec, err := setting.Cfg.GetSection("app")
	if err != nil {
		log.Fatal(2, "Fail to get section 'app': %v", err)
	}
	// 加盐
	salt := sec.Key("JWT_SALT").String()
	h.Write([]byte(pwd + salt))
	return fmt.Sprintf("%x", h.Sum(nil))
}

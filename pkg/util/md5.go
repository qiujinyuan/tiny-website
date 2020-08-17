package util

import (
	"crypto/md5"
	"encoding/hex"

	"github.com/yrjkqq/tiny-website/pkg/logging"
)

func EncodeMD5(value string) string {
	m := md5.New()
	_, err := m.Write([]byte(value))
	if err != nil {
		logging.Info(err)
		return ""
	}

	return hex.EncodeToString(m.Sum(nil))
}

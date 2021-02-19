package utility

import (
	"crypto/md5"
	"encoding/hex"
)

func GetMd5(buffer []byte) string {
	h := md5.New()
	h.Write(buffer)
	return hex.EncodeToString(h.Sum(nil))
}
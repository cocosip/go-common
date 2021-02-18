package utility

import (
	"crypto/hmac"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/hex"
)

func GetSha1(buffer []byte) string {
	h:=sha1.New()
	h.Write(buffer)
	//return fmt.Sprintf("%x",h.Sum(nil))
	return hex.EncodeToString(h.Sum(nil))
}

func GetSha256(buffer []byte) string {
	h:=sha256.New()
	h.Write(buffer)
	return hex.EncodeToString(h.Sum(nil))
}

func GetSha512(buffer []byte) string {
	h:=sha512.New()
	h.Write(buffer)
	return hex.EncodeToString(h.Sum(nil))
}


func GetHmacSha1(buffer []byte,key []byte) string {
	h := hmac.New(sha1.New, key)
	h.Write(buffer)
	//base64.StdEncoding.EncodeToString(h.Sum(nil))
	return hex.EncodeToString(h.Sum(nil))
}

func GetHmacSha256(buffer []byte,key []byte) string {
	h:=hmac.New(sha256.New,key)
	h.Write(buffer)
	return hex.EncodeToString(h.Sum(nil))
}

package public

import (
	"crypto/md5"
	"encoding/hex"
)

func Md5Enc(str, salt string) string {
	h := md5.New()
	h.Write([]byte(str))
	hash := h.Sum([]byte(salt))
	return hex.EncodeToString(hash)
}

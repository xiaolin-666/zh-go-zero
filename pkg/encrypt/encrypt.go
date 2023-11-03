package encrypt

import (
	"crypto/md5"
	"encoding/hex"
	"github.com/zeromicro/go-zero/core/codec"
)

const (
	mobileAesKey  = "BIN3FMTR1GFP3VVULDBQ5G5BEOWJ71AW"
	PasswdEncSeed = "JnB)Y+4hpp"
)

func EncMobile(mobile string) (string, error) {
	return codec.EcbEncryptBase64(mobileAesKey, mobile)
}

func EncPasswd(passwd string) string {
	md5Sum := md5.Sum([]byte(passwd + PasswdEncSeed))
	return hex.EncodeToString(bytes16toBytes(md5Sum))
}

func bytes16toBytes(in [16]byte) []byte {
	var tmp = make([]byte, 16)
	for i, val := range in {
		tmp[i] = val
	}
	return tmp
}

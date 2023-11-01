package encrypt

import "github.com/zeromicro/go-zero/core/codec"

const (
	mobileAesKey = "BIN3FMTR1GFP3VVULDBQ5G5BEOWJ71AW"
)

func EncMobile(mobile string) (string, error) {
	return codec.EcbEncryptBase64(mobileAesKey, mobile)
}

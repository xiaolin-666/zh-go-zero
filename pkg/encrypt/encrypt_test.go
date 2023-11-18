package encrypt

import (
	"testing"
)

func TestEncPasswd(t *testing.T) {
	passwd := EncPasswd("12345")
	t.Log(passwd)
}

func TestEncMobile(t *testing.T) {
	mobile, err := EncMobile("123456789")
	t.Log(mobile, err)
}

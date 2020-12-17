package validator

import (
	"github.com/TraiOi/translator"
)

func PrintErr(content string) string {
	return translator.ErrCode(content)
}

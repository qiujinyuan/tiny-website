package app

import (
	"fmt"

	"github.com/astaxie/beego/validation"
	"github.com/yrjkqq/tiny-website/pkg/logging"
)

func MarkErrors(errors []*validation.Error) (validErrorMsg string) {
	for _, err := range errors {
		logging.Info(err.Key, err.Message)
		validErrorMsg += fmt.Sprintf("Valid error: %v -> %v", err.Key, err.Message)
	}
	return
}

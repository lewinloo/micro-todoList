package controllers

import (
	"errors"
	logging "gateway/pkg/logging"
)

// 包装错误
func panicIfError(err error) {
	if err != nil {
		err = errors.New(err.Error())
		logging.Info(err)
		panic(any(err))
	}
}

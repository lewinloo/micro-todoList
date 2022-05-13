package controllers

import (
	"errors"
	logging "gateway/pkg/logging"
)

// 包装错误
func PanicIfTestError(err error) {
	if err != nil {
		err = errors.New("testService --" + err.Error())
		logging.Info(err)
		panic(any(err))
	}
}

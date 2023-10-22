package helper

import (
	"fmt"

	"github.com/sirupsen/logrus"
)

func LogError(errorMessage string, err error) {
	logrus.Error(fmt.Sprintf("%s:%s", errorMessage, err.Error()))
}

func LogInfo(info string) {
	logrus.Info(info)
}

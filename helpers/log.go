package helpers

import "go.uber.org/zap"

func LogError(logger *zap.Logger, ctrl string, method string, message string, a ...interface{}) {
	m := ctrl + ":" + method + ":" + message
	if len(a) > 0 {
		logger.Sugar().Error(m, "-", a)

	} else {
		logger.Sugar().Error(m)
	}
}

func LogInfo(logger *zap.Logger, ctrl string, method string, message string, a ...interface{}) {
	m := ctrl + ":" + method + ":" + message
	if len(a) > 0 {
		logger.Sugar().Info(m, "-", a)

	} else {
		logger.Sugar().Info(m)
	}
}

func LogWarning(logger *zap.Logger, ctrl string, method string, message string, a ...interface{}) {
	m := ctrl + ":" + method + ":" + message
	if len(a) > 0 {
		logger.Sugar().Warn(m, "-", a)

	} else {
		logger.Sugar().Warn(m)
	}
}

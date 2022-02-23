package artifact

import (
	"go.uber.org/zap"
)

var globalLog Logger

// Logger structure
type Logger struct {
	*zap.SugaredLogger
}

// GetLogger gets the global instance of the logger
func GetLogger() Logger {
	return globalLog
}

// NewLogger sets up logger
func NewLogger(env interface{}) Logger {

	config := zap.NewDevelopmentConfig()

	//if env.Environment == "local" {
	//	config.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	//} else {
	//	config.Level.SetLevel(zap.PanicLevel)
	//}

	logger, _ := config.Build()

	globalLog := logger.Sugar()

	return Logger{
		globalLog,
	}

}

package zapLogger

import (
	"fmt"
	"time"

	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"

	"github.com/AmitKarnam/KeyCloak/internal/utlis/logger"
)

var KClogger *zap.Logger

type ZapLogging struct{}

var _ logger.KClogging = &ZapLogging{}

func (zp *ZapLogging) GenerateLogger() {
	KClogger = initLogger()

	// Sync method is used to flush any buffered log entries to the underlying output.
	defer KClogger.Sync()
}

func initLogger() *zap.Logger {
	currentDate := time.Now().Format("2023-01-05")
	logFileName := fmt.Sprintf("logs/%s-log-%s.log", "keycloak", currentDate)

	logRotator := &lumberjack.Logger{
		Filename:   logFileName,
		MaxSize:    10, // Max size in MB before log rotation
		MaxBackups: 7,  // Max number of old log files to retain
		MaxAge:     7,  // Max number of days to retian the old log files
	}

	// Configurations required for the logs
	config := zap.NewProductionConfig()
	config.OutputPaths = []string{logFileName}
	config.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder

	// Building the core using the required core configuraitons
	core := zapcore.NewCore(
		zapcore.NewConsoleEncoder(config.EncoderConfig),
		zapcore.AddSync(logRotator),
		config.Level,
	)

	// Building the logger using the core
	logger := zap.New(core, zap.AddCaller())
	return logger
}

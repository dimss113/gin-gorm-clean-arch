package app

import (
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// var Logger *zap.Logger

/*
Debug - Information that is diagnostically helpful to people more than just developers (IT, sysadmins, etc.).
Info - Generally useful information to log (service start/stop, configuration assumptions, etc). Info I want to always have available but usually don't care about under normal circumstances. This is my out-of-the-box config level.
Warn - Anything that can potentially cause application oddities, but for which I am automatically recovering. (Such as switching from a primary to backup server, retrying an operation, missing secondary data, etc.)
Error - Any error which is fatal to the operation, but not the service or application (can't open a required file, missing data, etc.). These errors will force user (administrator, or direct user) intervention. These are usually reserved (in my apps) for incorrect connection strings, missing services, etc.
Fatal - Any error that is forcing a shutdown of the service or application to prevent data loss (or further data loss). I reserve these only for the most heinous errors and situations where there is guaranteed to have been data corruption or loss.

console : debug, info, warn, error, fatal
file : info, warn, error, fatal
*/
func InitializeLogger() *zap.Logger {
	config := zap.NewProductionEncoderConfig()
	config.EncodeTime = zapcore.ISO8601TimeEncoder

	fileEncoder := zapcore.NewJSONEncoder(config)
	consoleEncoder := zapcore.NewConsoleEncoder(config)

	logFileInfo, _ := os.OpenFile("logs/info.json", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	logFileWarn, _ := os.OpenFile("logs/warn.json", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	logFileError, _ := os.OpenFile("logs/error.json", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	logFileFatal, _ := os.OpenFile("logs/fatal.json", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)

	writerInfo := zapcore.AddSync(logFileInfo)
	writerWarn := zapcore.AddSync(logFileWarn)
	writerError := zapcore.AddSync(logFileError)
	writerFatal := zapcore.AddSync(logFileFatal)

	core := zapcore.NewTee(
		zapcore.NewCore(consoleEncoder, zapcore.AddSync(os.Stdout), zapcore.DebugLevel),
		zapcore.NewCore(fileEncoder, writerInfo, zapcore.InfoLevel),
		zapcore.NewCore(fileEncoder, writerWarn, zapcore.WarnLevel),
		zapcore.NewCore(fileEncoder, writerError, zapcore.ErrorLevel),
		zapcore.NewCore(fileEncoder, writerFatal, zapcore.FatalLevel),
	)
	logger := zap.New(core, zap.AddCaller(), zap.AddStacktrace(zapcore.ErrorLevel))
	return logger
}

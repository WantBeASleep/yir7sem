package custom

import (
	"fmt"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// Глобальный логгер
var Logger *zap.Logger

// Инициализация кастомного логгера
func Init() error {
	cfg := zap.Config{
		Level:       zap.NewAtomicLevelAt(zap.InfoLevel), // Уровень логирования
		Encoding:    "console",                           // Формат вывода (json)
		OutputPaths: []string{"stdout", "logs/logs.txt"}, // Путь к файлу и стандартный вывод
		// ErrorOutputPaths: []string{"stderr"},                   // Вывод ошибок
		EncoderConfig: zapcore.EncoderConfig{
			MessageKey:     "message",
			LevelKey:       "lvl",
			TimeKey:        "time",
			NameKey:        "logger",
			CallerKey:      "source",
			StacktraceKey:  "stacktrace",
			LineEnding:     zapcore.DefaultLineEnding,
			EncodeLevel:    zapcore.LowercaseLevelEncoder,
			EncodeTime:     zapcore.ISO8601TimeEncoder,
			EncodeDuration: zapcore.StringDurationEncoder,
			EncodeCaller:   zapcore.ShortCallerEncoder, // file.go:42 OR zapcore.FullCallerEncoder -> /path/to/file.go:42
		},
	}

	var err error
	Logger, err = cfg.Build()
	if err != nil {
		return fmt.Errorf("init logger failed: %w", err)
	}

	return nil
}

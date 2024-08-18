package log

import (
	"fmt"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

/*

1. Почекайте библиотеку zap

2. Логировать парами перед этим сообщив о самом логе, например берем данные с DB по ID:

	log.Debug("Response data from med_worker db", zap.Int("id": 1))

	Уровни логгирования Debug/Info/Warn/Error

	Логи все летят в консоль(можно перенастроить в файл, но оставим для прода)

*/

// develop config only
func New() (*zap.Logger, error) {
	encoderCfg := zap.NewDevelopmentEncoderConfig()
	encoderCfg.EncodeLevel = zapcore.CapitalColorLevelEncoder
	encoderCfg.EncodeTime = func(t time.Time, pae zapcore.PrimitiveArrayEncoder) {
		pae.AppendString(t.Format(time.DateTime))
	}

	cfg := zap.NewDevelopmentConfig()
	cfg.DisableCaller = true
	cfg.DisableStacktrace = true
	cfg.EncoderConfig = encoderCfg

	logger, err := cfg.Build()
	if err != nil {
		return nil, fmt.Errorf("build logger: %w", err)
	}
	return logger, nil
}
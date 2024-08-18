package log

import (
	"fmt"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

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

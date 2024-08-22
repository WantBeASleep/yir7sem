// ПОКУРИТЬ НАД ЭТИМ КРИНЖОМ, ВЫХОД ВВИЖУ ТОЛЬКО В ОБЩЕМ КОНИФГЕ НА ВСЕ ЧТО ФУЛ ЗАЛУПИЧИ
package log

import (
	"fmt"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// panic if have error!
// only DEV && PROD env accepted
// if env == DEV, outPath will be ommited
func New(env string, outPath string) *zap.Logger {
	var cfg zap.Config
	var encoderCfg zapcore.EncoderConfig

	switch env {
	case "DEV":
		encoderCfg = zap.NewDevelopmentEncoderConfig()
		encoderCfg.EncodeLevel = zapcore.CapitalColorLevelEncoder
		encoderCfg.EncodeTime = func(t time.Time, pae zapcore.PrimitiveArrayEncoder) {
			pae.AppendString(t.Format(time.DateTime))
		}

		cfg = zap.NewDevelopmentConfig()
		cfg.DisableCaller = true
		cfg.DisableStacktrace = true
		cfg.EncoderConfig = encoderCfg

	case "PROD":
		panic("not implemented")

	default:
		panic("wrong env type")
	}

	logger, err := cfg.Build()
	if err != nil {
		panic(fmt.Errorf("build logger: %w", err))
	}
	return logger
}

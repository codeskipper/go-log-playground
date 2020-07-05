package main

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"time"

	mysubpack "github.com/codeskipper/go-log-playground/zap-playground/mySubPack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var (
	_, b, _, _ = runtime.Caller(0)

	// Root folder of this project, it's one level up in this example
	// from Oleksiy Chechel at https://stackoverflow.com/a/38644571/4326287
	//projectFolder = filepath.Dir(b)
	projectFolder = filepath.Join(filepath.Dir(b), "..")
)

func main() {
	if err := do(); err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}
}

func do() error {
	config := zap.NewDevelopmentConfig()
	config.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	config.DisableStacktrace = true
	config.EncoderConfig.TimeKey = "timestamp"

	//config.EncoderConfig.EncodeTime = zapcore.RFC3339TimeEncoder
	config.EncoderConfig.EncodeTime = RFC3339MilliTimeEncoder
	logger, err := config.Build()
	if err != nil {
		return err
	}
	logger.Debug("test", zap.Any("foo", foo{One: "one", Two: "two"}))
	logger.Info("test with a longer message", zap.Any("foo", foo{One: "one", Two: "two"}))
	logger.Warn("test")
	logger.Error("test")
	mysubpack.Try(logger)
	return nil
}

type foo struct {
	One string
	Two string
}

// RFC3339MilliTimeEncoder serializes a time.Time to a RFC3339-formatted string with millisecond precision
//
// If enc supports AppendTimeLayout(t time.Time,layout string), it's used
// instead of appending a pre-formatted string value.
// adapted from zap/zapcore/encoder.go func RFC3339TimeEncoder
func RFC3339MilliTimeEncoder(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	//encodeTimeLayout(t, time.RFC3339, enc)
	encodeTimeLayout(t, "2006-01-02T15:04:05.999Z07:00", enc)
}

// copied from zap/zapcore/encoder.go to satisfy dependency for RFC3339MilliTimeEncoder
func encodeTimeLayout(t time.Time, layout string, enc zapcore.PrimitiveArrayEncoder) {
	type appendTimeEncoder interface {
		AppendTimeLayout(time.Time, string)
	}

	if enc, ok := enc.(appendTimeEncoder); ok {
		enc.AppendTimeLayout(t, layout)
		return
	}

	enc.AppendString(t.Format(layout))
}

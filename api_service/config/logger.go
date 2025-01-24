package config

import (
	"github.com/rs/zerolog"
	"gopkg.in/natefinch/lumberjack.v2"
)

var Logger zerolog.Logger

func CreateLogger(filename string) zerolog.Logger {
	z := zerolog.New(&lumberjack.Logger{
		Filename:   filename,
		MaxSize:    100,
		MaxBackups: 5,
		MaxAge:     30,
		Compress:   false,
	})
	z = z.With().Caller().Timestamp().Logger()
	Logger = z

	return z
}

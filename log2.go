package jccclient

import (
	"encoding/json"

	"go.uber.org/zap"
)

var mainLogger *zap.Logger

// NewLogger - new a zap.Logger
func NewLogger(logfn string) (*zap.Logger, error) {
	cfg := zap.NewProductionConfig()
	cfg.OutputPaths = []string{
		logfn,
	}

	return cfg.Build()
}

// SetMainLogger - set main logger
func SetMainLogger(logger *zap.Logger) {
	mainLogger = logger
}

func init() {
	logger, _ := zap.NewProduction()
	mainLogger = logger
}

// InitLogger - init logger
func InitLogger() error {
	logger, err := zap.NewProduction()
	if err != nil {
		return err
	}

	SetMainLogger(logger)

	return nil
}

// JSON - make json to field
func JSON(key string, obj interface{}) zap.Field {
	s, err := json.Marshal(obj)
	if err != nil {
		return zap.Error(err)
	}

	return zap.String(key, string(s))
}

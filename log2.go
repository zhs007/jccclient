package jccclient

import "go.uber.org/zap"

// NewLogger - new a zap.Logger
func NewLogger(logfn string) (*zap.Logger, error) {
	cfg := zap.NewProductionConfig()
	cfg.OutputPaths = []string{
	  logfn,
	}
	
	return cfg.Build()
  }
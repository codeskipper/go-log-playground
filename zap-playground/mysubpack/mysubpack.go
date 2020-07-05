package mysubpack

// test package for caller path logging

import (
	"go.uber.org/zap"
)

// Try will log test entry
func Try(logger *zap.Logger) {
	logger.Info("from file in mySubPack")
}

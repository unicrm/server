package logging

import (
	"testing"

	"go.uber.org/zap"
)

var TEST_LOGGER *zap.Logger

func TestLogger(t *testing.T) {
	config := Logger{"debug", "log", "console", "stacktrace", "[unicrm]", true, 28, true}
	TEST_LOGGER = LoggerInit(config)
	TEST_LOGGER.Debug("Debug")
	TEST_LOGGER.Info("Info")
	TEST_LOGGER.Warn("Warn")
	TEST_LOGGER.Error("Error")
}

package logging

import (
	"github.com/geomyidia/zylog/logger"
	"github.com/hexagram30/raster/components/config"
	log "github.com/sirupsen/logrus"
)

// Setup ...
func Setup(cfg *config.Config) {
	logger.SetupLogging(cfg.Logging)
}

// Load pretends that the global is more functional in nature ...
func Load(cfg *config.Config) *log.Logger {
	Setup(cfg)
	return log.StandardLogger()
}

// SetupClient ...
func SetupClient(cfg *config.Config) {
	logger.SetupLogging(cfg.Client.Logging)
}

// LoadClient pretends that the global is more functional in nature ...
func LoadClient(cfg *config.Config) *log.Logger {
	SetupClient(cfg)
	return log.StandardLogger()
}

package components

import (
	"github.com/geomyidia/reverb"
	"github.com/hexagram30/raster/components/config"
	"github.com/hexagram30/raster/components/db"
	logger "github.com/sirupsen/logrus"
)

// Base component collection
type Base struct {
	Config *config.Config
	Logger *logger.Logger
}

// TestBase component that keeps stdout clean
type TestBase struct {
	Config *config.Config
}

// Default component collection
type Default struct {
	Base
	GRPCD *reverb.Reverb
	DB    db.DB
}

// Add more components here that have more or less than what's done above. This
// is useful for testing or runnning in different binaries/executables, etc.

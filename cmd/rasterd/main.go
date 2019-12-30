package main

import (
	"github.com/geomyidia/reverb"
	"github.com/hexagram30/raster/app"
	"github.com/hexagram30/raster/components/config"
	"github.com/hexagram30/raster/components/logging"
)

func main() {
	// Create the application object and assign components to it
	a := new(app.Application)
	a.Config = config.NewConfig()
	a.Logger = logging.Load(a.Config)
	a.GRPCD = reverb.New()

	// Perform application setup and then start the services)
	a.Start()
}

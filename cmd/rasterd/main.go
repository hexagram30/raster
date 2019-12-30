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
	a.Config = config.New()
	a.Logger = logging.Load(a.Config)
	a.GRPCD = reverb.New()

	// Connect to storage server
	a.DB = a.SetupDB()
	defer a.DB.Close()

	// Perform application setup and then start the services)
	a.Start()
	defer a.Close()
}

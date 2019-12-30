package app

import (
	"github.com/geomyidia/reverb"
	"github.com/hexagram30/raster/app/handlers"
	"github.com/hexagram30/raster/components"
	log "github.com/sirupsen/logrus"
)

// Application ...
type Application struct {
	components.Default
}

// SetupgRPCImplementation ...
func (a *Application) SetupgRPCImplementation(r *reverb.Reverb) {
	log.Debug("Setting up gRPC implementation ...")
	s := handlers.NewGRPCHandlerServer()
	s.RegisterServer(r.GRPCServer)
	log.Info("gRPC implementation set up.")
}

// StartgRPCD ...
func (a *Application) StartgRPCD() {
	log.Debug("Starting gRPC daemon ...")
	serverOpts := a.Config.GRPCConnectionString()
	server := a.GRPCD.Start(serverOpts)
	a.SetupgRPCImplementation(server)
	go server.Serve()
	log.Infof("gRPC daemon started on %s.", serverOpts)
}

// Start ...
func (a *Application) Start() {
	a.StartgRPCD()
	log.Info("System started.")
}

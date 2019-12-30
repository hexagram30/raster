package app

import (
	"github.com/geomyidia/reverb"
	"github.com/hexagram30/raster/app/handlers"
	"github.com/hexagram30/raster/components"
	"github.com/hexagram30/raster/components/db"
	log "github.com/sirupsen/logrus"
)

// Application ...
type Application struct {
	components.Default
}

// Close ...
func (a *Application) Close() {
	a.GRPCD.TCPListener.Close()
}

// SetupgRPCImplementation ...
func (a *Application) SetupgRPCImplementation(r *reverb.Reverb) {
	log.Debug("Setting up gRPC implementation ...")
	s := handlers.NewGRPCHandlerServer()
	s.RegisterServer(r.GRPCServer)
	log.Info("gRPC implementation set up.")
}

// SetupDB ...
func (a *Application) SetupDB() db.DB {
	log.Debug("Setting up database ...")
	conn, err := db.New(a.Config)
	if err != nil {
		log.Panic(err)
	}
	log.Info("DB store is set up.")
	return conn
}

// StartgRPCD ...
func (a *Application) StartgRPCD() {
	log.Debug("Starting gRPC daemon ...")
	serverOpts := a.Config.GRPCConnectionString()
	server := a.GRPCD.Start(serverOpts)
	a.SetupgRPCImplementation(server)
	server.Serve()
	log.Infof("gRPC daemon started on %s.", serverOpts)
}

// Start ...
func (a *Application) Start() {
	a.StartgRPCD()
	log.Info("System started.")
}

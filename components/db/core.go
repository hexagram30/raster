package db

import (
	"fmt"

	"github.com/hexagram30/raster/components/config"
	log "github.com/sirupsen/logrus"
)

// db-wide constants
const (
	// Connection types
	REDIX string = "redix"
)

// DB ...
type DB interface {
	Close() error
	DBType() string
}

// New dispatches the creation of a specific DB connection type
func New(cfg *config.Config) (DB, error) {

	switch cfg.DB.Type {
	case REDIX:
		conn, err := NewRedixConnection(cfg)
		if err != nil {
			log.Error("Could not connect to RedixDB")
			return conn, err
		}
		return conn, nil
	default:
		return nil, fmt.Errorf(
			"database connection type %s not supported", cfg.DB.Type)
	}
}

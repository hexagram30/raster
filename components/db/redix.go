package db

import (
	"fmt"

	"github.com/go-redis/redis"
	"github.com/hexagram30/raster/components/config"
	log "github.com/sirupsen/logrus"
)

// RedixConnector ...
type RedixConnector struct {
	ConnType string
	Options  *redis.Options
	Conn     *redis.Client
}

// NewRedixConnection is for use with unit tests
func NewRedixConnection(cfg *config.Config) (*RedixConnector, error) {
	addr := fmt.Sprintf("%s:%d", cfg.DB.Redix.Host, cfg.DB.Redix.Port)
	opts := &redis.Options{Addr: addr}
	log.Tracef("Redix connection options: %#v", opts)
	conn := redis.NewClient(opts)
	log.Tracef("Redix connection: %v", conn)
	connector := &RedixConnector{
		ConnType: REDIX,
		Options:  opts,
		Conn:     conn,
	}
	log.Debug("Checking connection with Redix ...")
	pong, err := connector.Conn.Ping().Result()
	if err != nil {
		log.Error("Couldn't create Redix session ...")
		log.Fatal(err)
	}
	if pong != "PONG" {
		log.Fatalf("Unexpected pong result '%s' during connection check", pong)
	}
	return connector, err
}

// Close closes the connection to the database
func (c *RedixConnector) Close() error {
	c.Conn.Close()
	return nil
}

// DBType returns the database type
func (c *RedixConnector) DBType() string {
	return c.ConnType
}

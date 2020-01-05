package client

import (
	"time"

	pb "github.com/hexagram30/raster/api"
	"github.com/hexagram30/raster/common"
	"github.com/hexagram30/raster/components"
	"github.com/hexagram30/raster/components/config"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

const (
	shortTimeout = (20 * time.Millisecond) + 1*time.Minute
)

//Client is the wrapper for the client obj
type Client struct {
	components.Base
	components.BaseApp
	components.BaseRPCClient
}

// New ...
func New() *Client {
	client := Client{}
	client.AppName = config.AppName
	client.AppAbbv = config.AppName
	client.ProjectPath = common.CallerPaths().DotPath
	return &client
}

// Close the gRPC connection
func (c *Client) Close() {
	c.RPCConn.Close()
}

// SetupConnection ...
func (c *Client) SetupConnection() {
	connectionOpts := c.Config.GRPCConnectionString()
	conn, err := grpc.Dial(connectionOpts, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect to gRPC server: %v", err)
	}
	c.RPCConn = conn
	c.RPCClient = pb.NewServiceAPIClient(conn)
}

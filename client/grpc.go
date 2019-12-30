package client

import (
	"context"

	"github.com/golang/protobuf/proto"
	pb "github.com/hexagram30/raster/api"
	log "github.com/sirupsen/logrus"
)

/////////////////////////////////////////////////////////////////////////////
///   Client Implementation for the gRPC Server API    //////////////////////
/////////////////////////////////////////////////////////////////////////////

// Echo ...
func (c *Client) Echo(msg string) string {

	ctx, cancel := context.WithTimeout(context.Background(), shortTimeout)
	defer cancel()

	r, err := c.RPCClient.Echo(ctx, &pb.GenericData{
		Data: *proto.String(msg),
	})
	if err != nil {
		log.Fatalf("Could not get echo reply: %v", err)
	}

	return r.GetData()
}

// Health ...
func (c *Client) Health() string {

	ctx, cancel := context.WithTimeout(context.Background(), shortTimeout)
	defer cancel()

	r, err := c.RPCClient.Health(ctx, &pb.HealthRequest{})
	if err != nil {
		log.Fatalf("Could not get health reply: %v", err)
	}
	log.Printf("Errors: %s", r.GetErrors())

	return r.GetComponents()
}

// Ping ...
func (c *Client) Ping() string {

	ctx, cancel := context.WithTimeout(context.Background(), shortTimeout)
	defer cancel()

	r, err := c.RPCClient.Ping(ctx, &pb.PingRequest{})
	if err != nil {
		log.Fatalf("Could not get ping reply: %v", err)
	}

	return r.GetData()
}

// Version ...
func (c *Client) Version() *pb.VersionReply {

	ctx, cancel := context.WithTimeout(context.Background(), shortTimeout)
	defer cancel()

	r, err := c.RPCClient.Version(ctx, &pb.VersionRequest{})
	if err != nil {
		log.Fatalf("Could not get version: %v", err)
	}
	log.Debugf("Version data: %v", r)

	return r
}

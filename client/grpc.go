package client

import (
	"context"

	"github.com/hexagram30/protocols/src/golang/common"
	log "github.com/sirupsen/logrus"
)

/////////////////////////////////////////////////////////////////////////////
///   Client Implementation for the gRPC Server API    //////////////////////
/////////////////////////////////////////////////////////////////////////////

// Ping ...
func (c *Client) Ping() string {

	ctx, cancel := context.WithTimeout(context.Background(), shortTimeout)
	defer cancel()

	r, err := c.RPCClient.Ping(ctx, &common.PingRequest{})
	if err != nil {
		log.Fatalf("Could not get ping reply: %v", err)
	}

	return r.GetData()
}

// Version ...
func (c *Client) Version() *common.VersionReply {

	ctx, cancel := context.WithTimeout(context.Background(), shortTimeout)
	defer cancel()

	r, err := c.RPCClient.Version(ctx, &common.VersionRequest{})
	if err != nil {
		log.Fatalf("Could not get version: %v", err)
	}
	log.Debugf("Version data: %v", r)

	return r
}

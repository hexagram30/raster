package handlers

import (
	"context"

	"google.golang.org/grpc"

	"github.com/golang/protobuf/proto"
	base "github.com/hexagram30/protocols/src/golang/common"
	"github.com/hexagram30/raster/api"
	"github.com/hexagram30/raster/common"
	log "github.com/sirupsen/logrus"
)

// GRPCHandlerServer ...
type GRPCHandlerServer struct {
	api.UnimplementedServiceAPIServer
}

// NewGRPCHandlerServer ...
func NewGRPCHandlerServer() *GRPCHandlerServer {
	return &GRPCHandlerServer{}
}

// RegisterServer ...
func (s *GRPCHandlerServer) RegisterServer(grpcServer *grpc.Server) {
	api.RegisterServiceAPIServer(grpcServer, s)
}

// Ping ...
func (s *GRPCHandlerServer) Ping(ctx context.Context, in *base.PingRequest) (*base.PingReply, error) {
	log.Debugf("Received: %v", in)
	return &base.PingReply{Data: "PONG"}, nil
}

// Version ...
func (s *GRPCHandlerServer) Version(
	ctx context.Context, in *base.VersionRequest) (*base.VersionReply, error) {
	log.Tracef("Received: %v", in)
	version := common.VersionData()
	return &base.VersionReply{
		Version:    *proto.String(version.Semantic),
		BuildDate:  *proto.String(version.BuildDate),
		GitCommit:  *proto.String(version.GitCommit),
		GitBranch:  *proto.String(version.GitBranch),
		GitSummary: *proto.String(version.GitSummary),
	}, nil
}

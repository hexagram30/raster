package handlers

import (
	"context"

	"google.golang.org/grpc"

	"github.com/golang/protobuf/proto"
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

// Echo ...
func (s *GRPCHandlerServer) Echo(ctx context.Context, in *api.GenericData) (*api.GenericData, error) {
	log.Debugf("Received: %v", in)
	return &api.GenericData{Data: in.GetData()}, nil
}

// Health ...
func (s *GRPCHandlerServer) Health(ctx context.Context, in *api.HealthRequest) (*api.HealthReply, error) {
	log.Debugf("Received: %v", in)
	return &api.HealthReply{Components: "OK", Errors: "NULL"}, nil
}

// Ping ...
func (s *GRPCHandlerServer) Ping(ctx context.Context, in *api.PingRequest) (*api.PingReply, error) {
	log.Debugf("Received: %v", in)
	return &api.PingReply{Data: "PONG"}, nil
}

// Version ...
func (s *GRPCHandlerServer) Version(
	ctx context.Context, in *api.VersionRequest) (*api.VersionReply, error) {
	log.Tracef("Received: %v", in)
	version := common.VersionData()
	return &api.VersionReply{
		Version:    *proto.String(version.Semantic),
		BuildDate:  *proto.String(version.BuildDate),
		GitCommit:  *proto.String(version.GitCommit),
		GitBranch:  *proto.String(version.GitBranch),
		GitSummary: *proto.String(version.GitSummary),
	}, nil
}

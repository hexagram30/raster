package handlers

import (
	"context"

	"google.golang.org/grpc"

	"github.com/golang/protobuf/proto"
	pb "github.com/hexagram30/raster/api"
	"github.com/hexagram30/raster/common"
	log "github.com/sirupsen/logrus"
)

// GRPCHandlerServer ...
type GRPCHandlerServer struct {
	pb.UnimplementedServiceAPIServer
}

// NewGRPCHandlerServer ...
func NewGRPCHandlerServer() *GRPCHandlerServer {
	return &GRPCHandlerServer{}
}

// RegisterServer ...
func (s *GRPCHandlerServer) RegisterServer(grpcServer *grpc.Server) {
	pb.RegisterServiceAPIServer(grpcServer, s)
}

// Echo ...
func (s *GRPCHandlerServer) Echo(ctx context.Context, in *pb.GenericData) (*pb.GenericData, error) {
	log.Debugf("Received: %v", in)
	return &pb.GenericData{Data: in.GetData()}, nil
}

// Health ...
func (s *GRPCHandlerServer) Health(ctx context.Context, in *pb.HealthRequest) (*pb.HealthReply, error) {
	log.Debugf("Received: %v", in)
	return &pb.HealthReply{Components: "OK", Errors: "NULL"}, nil
}

// Ping ...
func (s *GRPCHandlerServer) Ping(ctx context.Context, in *pb.PingRequest) (*pb.PingReply, error) {
	log.Debugf("Received: %v", in)
	return &pb.PingReply{Data: "PONG"}, nil
}

// Version ...
func (s *GRPCHandlerServer) Version(
	ctx context.Context, in *pb.VersionRequest) (*pb.VersionReply, error) {
	log.Tracef("Received: %v", in)
	version := common.VersionData()
	return &pb.VersionReply{
		Version:    *proto.String(version.Semantic),
		BuildDate:  *proto.String(version.BuildDate),
		GitCommit:  *proto.String(version.GitCommit),
		GitBranch:  *proto.String(version.GitBranch),
		GitSummary: *proto.String(version.GitSummary),
	}, nil
}

package handlers

import (
	"context"

	"google.golang.org/grpc"

	pb "github.com/hexagram30/raster/api"
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

// RegisterServer ...
func (s *GRPCHandlerServer) RegisterServer(grpcServer *grpc.Server) {
	pb.RegisterServiceAPIServer(grpcServer, s)
}

// Code generated by github.com/izumin5210/grapi. DO NOT EDIT.

package server

import (
	"context"

	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"google.golang.org/grpc"

	api_pb "emergence/api"
)

// RegisterWithServer implements grapiserver.Server.RegisterWithServer.
func (s *hospitalServiceServerImpl) RegisterWithServer(grpcSvr *grpc.Server) {
	api_pb.RegisterHospitalServiceServer(grpcSvr, s)
}

// RegisterWithHandler implements grapiserver.Server.RegisterWithHandler.
func (s *hospitalServiceServerImpl) RegisterWithHandler(ctx context.Context, mux *runtime.ServeMux, conn *grpc.ClientConn) error {
	return api_pb.RegisterHospitalServiceHandler(ctx, mux, conn)
}
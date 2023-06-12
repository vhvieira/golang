package main

import (
	"context"
	"golang/grpc/proto"
	"log"
	"net"

	"google.golang.org/grpc"
)

// should implement the interface myPkgName.InvoicerServer
type myGRPCInvoiceServer struct {
	// type embedded to comply with Google lib
	proto.UnimplementedInvoicerServer
}

// should implement the interface myPkgName.InvoicerServer
type myGRPCHealthServer struct {
	// type embedded to comply with Google lib
	proto.UnimplementedHealthCheckServer
}

func (m *myGRPCInvoiceServer) Create(ctx context.Context, request *proto.CreateRequest) (*proto.CreateResponse, error) {
	log.Println("Received request", request)
	return &proto.CreateResponse{Pdf: []byte("TODO")}, nil
}

func (m *myGRPCHealthServer) Ping(ctx context.Context, request *proto.HealthRequest) (*proto.HealthResponse, error) {
	log.Println("Received request", request)
	return &proto.HealthResponse{Status: "OK"}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":3333")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	myGRPCServerReference := &myGRPCInvoiceServer{}
	myGRPCHealthReference := &myGRPCHealthServer{}
	proto.RegisterInvoicerServer(s, myGRPCServerReference)
	proto.RegisterHealthCheckServer(s, myGRPCHealthReference)
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

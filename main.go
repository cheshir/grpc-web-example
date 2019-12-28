package main

import (
	"context"
	"log"
	"net"
	"net/http"
	"time"
	pb "webgrpc/api/gen/currenttime"

	"github.com/golang/protobuf/ptypes/empty"
	"github.com/golang/protobuf/ptypes/timestamp"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

const (
	HTTPAddress = ":80"
	GRPCAddress = ":9090"
)

func main() {
	go func() {
		if err := runGRPCServer(); err != nil {
			panic(err)
		}
	}()

	if err := runHTTPServer(); err != nil {
		panic(err)
	}
}

func runGRPCServer() error {
	listener, err := net.Listen("tcp", GRPCAddress)
	if err != nil {
		return err
	}

	grpcServer := grpc.NewServer()
	pb.RegisterCurrentTimerServer(grpcServer, &CurrentTimerServer{})
	log.Printf("Starting GRPC server on %s", GRPCAddress)

	return grpcServer.Serve(listener)
}

type CurrentTimerServer struct{}

func (s *CurrentTimerServer) GetCurrentTime(ctx context.Context, empty *empty.Empty) (*pb.CurrentTimeResponse, error) {
	t := time.Now()

	// Fail with some probability.
	failFactor := 8
	if t.Second()%failFactor == 0 {
		return nil, status.Error(codes.Unknown, "Unknown")
	}

	response := &pb.CurrentTimeResponse{
		Timestamp: &timestamp.Timestamp{
			Seconds: t.Unix(),
			Nanos:   int32(t.Nanosecond()),
		},
	}

	return response, nil
}

func runHTTPServer() error {
	http.Handle("/", http.FileServer(http.Dir("static/dist")))
	log.Printf("Starting HTTP server on %s", HTTPAddress)

	return http.ListenAndServe(HTTPAddress, nil)
}

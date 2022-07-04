package main

import (
	"fmt"
	"net"

	"github.com/RomanoLu/myaktion-go/src/banktransfer/grpc/banktransfer"
	"github.com/RomanoLu/myaktion-go/src/banktransfer/service"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

func init() {
	log.SetLevel(log.TraceLevel)
	log.SetFormatter(&log.TextFormatter{})
	log.SetReportCaller(true)
}

var grpcPort = 9111

func main() {
	log.Info("Starting Banktransfer server")
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", grpcPort))
	if err != nil {
		log.Fatalf("failed to listen on grpc port %d: %v", grpcPort, err)
	}
	grpcServer := grpc.NewServer()
	banktransfer.RegisterBankTransferServer(grpcServer, service.NewBankTransferService())
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

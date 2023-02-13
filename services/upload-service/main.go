package main

import (
	"fmt"
	"google.golang.org/grpc"
	"log"
	"net"
	"uploadService/services/upload-service/config"
	upload_service "uploadService/services/upload-service/proto"
	pb "uploadService/services/upload-service/proto/upload-service"
)

func main() {
	conf := config.UploadConfig()

	uploadServiceHost := conf.HostUpload
	uploadServicePort := conf.PortUpload

	uploadServiceAddress := fmt.Sprintf("%s:%s", uploadServiceHost, uploadServicePort)

	lis, err := net.Listen("tcp", uploadServiceAddress)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	server := grpc.NewServer()

	pb.RegisterUploadServiceServer(
		server,
		&upload_service.Server{},
	)

	log.Printf("server listening at %v", lis.Addr())

	err = server.Serve(lis)

	if err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}

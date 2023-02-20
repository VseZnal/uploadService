package main

import (
	"google.golang.org/grpc"
	"log"
	"net"
	"uploadService/cmd/server/config"
	upload2 "uploadService/cmd/server/upload"
	proto_list_album_service "uploadService/proto"
)

func main() {
	conf := config.UploadServerConfig()
	port := conf.PortServer
	static := conf.Static

	lis, err := net.Listen("tcp", ":"+port)
	if err != nil {
		log.Fatal(err)
	}
	defer lis.Close()

	uplSrv := upload2.NewServer(upload2.New(static))
	rpcSrv := grpc.NewServer()

	proto_list_album_service.RegisterUploadServiceServer(rpcSrv, uplSrv)
	log.Fatal(rpcSrv.Serve(lis))
}

package main

import (
	"google.golang.org/grpc"
	"log"
	"net"
	upload2 "uploadService/cmd/server/upload"
	proto_list_album_service "uploadService/proto"
)

func main() {
	lis, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal(err)
	}
	defer lis.Close()

	uplSrv := upload2.NewServer(upload2.New("tmp/"))

	rpcSrv := grpc.NewServer()

	proto_list_album_service.RegisterUploadServiceServer(rpcSrv, uplSrv)
	log.Fatal(rpcSrv.Serve(lis))
}

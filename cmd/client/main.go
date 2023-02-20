package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/metadata"
	"log"
	"uploadService/cmd/client/config"
	"uploadService/cmd/client/upload"
)

func main() {
	conf := config.UploadClientConfig()
	hostClient := conf.HostUpload
	portClient := conf.PortUpload

	clientServiceAddress := fmt.Sprintf("%s:%s", hostClient, portClient)

	conn, err := grpc.Dial(clientServiceAddress,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		log.Fatalln(err)
	}
	defer conn.Close()

	client := upload.NewClient(conn)
	header := metadata.New(map[string]string{"name": "zxc.jpg"})
	ctx := metadata.NewOutgoingContext(context.Background(), header)

	name, err := client.Upload(ctx, "/home/vseznal/t4k6licnFdc.jpg")

	if err != nil {
		log.Fatalln(err)
	}
	log.Println(name)
}

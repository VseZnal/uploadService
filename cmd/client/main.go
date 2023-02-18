package main

import (
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/metadata"
	"log"
	"uploadService/cmd/client/upload"
)

func main() {
	//flag.Parse()
	//if flag.NArg() == 0 {
	//	log.Fatalln("Missing file path")
	//}
	conn, err := grpc.Dial("localhost:8080",
		grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalln(err)
	}
	defer conn.Close()

	client := upload.NewClient(conn)
	//"/home/vseznal/t4k6licnFdc.jpg"

	header := metadata.New(map[string]string{"name": "zxc.jpg"})
	ctx := metadata.NewOutgoingContext(context.Background(), header)

	name, err := client.Upload(ctx, "/home/vseznal/t4k6licnFdc.jpg")

	if err != nil {
		log.Fatalln(err)
	}
	log.Println(name)
}

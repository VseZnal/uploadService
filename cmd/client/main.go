package main

import (
	"context"
	"flag"
	"google.golang.org/grpc"
	"log"
	"uploadService/cmd/client/upload"
)

func main() {
	flag.Parse()
	if flag.NArg() == 0 {
		log.Fatalln("Missing file path")
	}

	conn, err := grpc.Dial(":50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalln(err)
	}
	defer conn.Close()

	client := upload.NewClient(conn)
	name, err := client.Upload(context.Background(), flag.Arg(0))
	if err != nil {
		log.Fatalln(err)
	}
	log.Println(name)
}

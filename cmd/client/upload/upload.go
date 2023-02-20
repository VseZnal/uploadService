package upload

import (
	"context"
	"io"
	"os"
	"time"
	"uploadService/libs/errors"
	proto_list_album_service "uploadService/proto"

	"google.golang.org/grpc"
)

type Client struct {
	client proto_list_album_service.UploadServiceClient
}

func NewClient(conn grpc.ClientConnInterface) Client {
	return Client{
		client: proto_list_album_service.NewUploadServiceClient(conn),
	}
}

func (c Client) Upload(ctx context.Context, file string) (string, error) {
	ctx, cancel := context.WithDeadline(ctx, time.Now().Add(10*time.Second))
	defer cancel()

	stream, err := c.client.Upload(ctx)
	if err != nil {
		return "", errors.LogError(err)
	}

	fil, err := os.Open(file)
	if err != nil {
		return "", errors.LogError(err)
	}

	buf := make([]byte, 1024)

	for {
		num, err := fil.Read(buf)
		if err == io.EOF {
			break
		}
		if err != nil {
			return "", errors.LogError(err)
		}

		if err := stream.Send(&proto_list_album_service.UploadRequest{Chunk: buf[:num]}); err != nil {
			return "", errors.LogError(err)
		}
	}

	res, err := stream.CloseAndRecv()
	if err != nil {
		return "", err
	}

	return res.GetName(), nil
}

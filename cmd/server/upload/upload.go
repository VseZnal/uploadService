package upload

import (
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
	"io"
	"uploadService/cmd/server/repository"
	"uploadService/libs/errors"
	upload_service "uploadService/proto"
)

type Server struct {
	upload_service.UnimplementedUploadServiceServer
	storage Manager
	limit   bool
}

func NewServer(storage Manager) Server {
	return Server{
		storage: storage,
	}
}

var db repository.Database

func Init() error {
	var err error

	db, err = repository.NewDatabase()
	return err
}

func (s Server) Upload(stream upload_service.UploadService_UploadServer) error {
	md, _ := metadata.FromIncomingContext(stream.Context())

	name := md.Get("name")[0]
	file := NewFile(name)
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			err := s.storage.Store(file)
			if err != nil {
				return status.Error(codes.Internal, err.Error())
			}

			err = db.CreateImageInfo(name)
			if err != nil {
				return errors.LogError(err)
			}

			return stream.SendAndClose(&upload_service.UploadResponse{
				Name:      req.GetName(),
				CreatedAt: timestamppb.Now()})
		}
		if err != nil {
			return status.Error(codes.Internal, err.Error())
		}

		if err := file.Write(req.GetChunk()); err != nil {
			return status.Error(codes.Internal, err.Error())
		}
	}

}

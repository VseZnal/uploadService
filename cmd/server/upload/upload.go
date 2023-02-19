package upload

import (
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"io"
	proto_list_album_service "uploadService/proto"
)

type Server struct {
	proto_list_album_service.UnimplementedUploadServiceServer
	storage Manager
}

func NewServer(storage Manager) Server {
	return Server{
		storage: storage,
	}
}

func (s Server) Upload(stream proto_list_album_service.UploadService_UploadServer) error {
	md, _ := metadata.FromIncomingContext(stream.Context())

	name := md.Get("name")[0]
	file := NewFile(name)
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			if err := s.storage.Store(file); err != nil {
				return status.Error(codes.Internal, err.Error())
			}
			return stream.SendAndClose(&proto_list_album_service.UploadResponse{Name: req.GetName()})
		}
		if err != nil {
			return status.Error(codes.Internal, err.Error())
		}

		if err := file.Write(req.GetChunk()); err != nil {
			return status.Error(codes.Internal, err.Error())
		}
	}

}

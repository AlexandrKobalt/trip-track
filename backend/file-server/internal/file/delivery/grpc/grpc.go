package grpc

import (
	"context"

	"github.com/AlexandrKobalt/trip-track/backend/file-server/internal/file/service"
	fileserviceproto "github.com/AlexandrKobalt/trip-track/backend/file-server/proto"
	"google.golang.org/protobuf/types/known/emptypb"
)

type Server struct {
	service service.IService
	fileserviceproto.FileServer
}

func New(service service.IService) fileserviceproto.FileServer {
	return &Server{service: service}
}

func (s *Server) Upload(
	_ context.Context,
	request *fileserviceproto.UploadRequest,
) (response *emptypb.Empty, err error) {
	return s.service.Upload(request)
}

func (s *Server) GetURL(
	_ context.Context,
	request *fileserviceproto.GetURLRequest,
) (response *fileserviceproto.GetURLResponse, err error) {
	return s.service.GetURL(request)
}

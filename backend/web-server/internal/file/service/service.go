package service

import "github.com/AlexandrKobalt/trip-track/backend/web-server/internal/file/models"

type service struct {
}

func New() IService {
	return &service{}
}

func (s *service) Upload(params models.UploadParams) error {
	return nil
}

func (s *service) GetURL(params models.GetURLParams) (result models.GetURLResult, err error) {
	return result, nil
}

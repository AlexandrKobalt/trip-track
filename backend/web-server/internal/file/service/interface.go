package service

import "github.com/AlexandrKobalt/trip-track/backend/web-server/internal/file/models"

type IService interface {
	Upload(params models.UploadParams) (err error)
	GetURL(params models.GetURLParams) (result models.GetURLResult, err error)
}

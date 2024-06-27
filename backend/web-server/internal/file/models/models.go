package models

import "mime/multipart"

type UploadParams struct {
	File *multipart.FileHeader
}

type GetURLParams struct {
	Key string `json:"key"`
}

type GetURLResult struct {
	URL string `json:"url"`
}

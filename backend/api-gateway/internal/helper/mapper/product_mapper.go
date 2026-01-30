package mapper

import (
	"io"
	"mime/multipart"

	productpb "api-gateway/internal/pb/product"
)

func MapFileToProductThumbnail(fileHeader *multipart.FileHeader) *productpb.ImageRequest {
	file, err := fileHeader.Open()
	if err != nil {
		return nil
	}

	bytes, err := io.ReadAll(file)
	file.Close()
	if err != nil {
		return nil
	}

	return &productpb.ImageRequest{
		Filename:    fileHeader.Filename,
		ContentType: fileHeader.Header.Get("Content-Type"),
		Data:        bytes,
	}
}

func MapFilesToProductImages(files []*multipart.FileHeader) []*productpb.ImageRequest {
	var grpcImages []*productpb.ImageRequest

	for _, fileHeader := range files {
		grpcImages = append(grpcImages, MapFileToProductThumbnail(fileHeader))
	}

	return grpcImages
}

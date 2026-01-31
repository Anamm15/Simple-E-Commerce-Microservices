package service

import (
	"context"

	"product-service/internal/dto"
	"product-service/internal/helper"
	"product-service/internal/helper/mapper"
	"product-service/internal/infrastructure/cloud"
	"product-service/internal/model"
	productpb "product-service/internal/pb/product"
	"product-service/internal/repository"
	"product-service/pkg/util"
)

type ImageService interface {
	AddImageProduct(ctx context.Context, request dto.AddImageProductRequestDTO) (*productpb.ImageProductResponse, error)
	DeleteImageProduct(ctx context.Context, imageID string) error
}

type imageService struct {
	imageRepo    repository.ImageRepository
	cloudStorage cloud.CloudinaryService
}

func NewImageService(imageRepo repository.ImageRepository, cloudStorage cloud.CloudinaryService) ImageService {
	return &imageService{imageRepo: imageRepo, cloudStorage: cloudStorage}
}

func (r *imageService) AddImageProduct(ctx context.Context, request dto.AddImageProductRequestDTO) (*productpb.ImageProductResponse, error) {
	var images []model.Image
	for _, image := range request.Images {
		productID, err := util.StringToUUID(request.ProductID)
		if err != nil {
			return nil, err
		}

		uniqueFileName := helper.GenerateRandomFilename()
		fileReader := util.ByteToIOReader(image.Data)

		fileUrl, err := r.cloudStorage.UploadFile(ctx, fileReader, uniqueFileName)
		if err != nil {
			return nil, err
		}
		images = append(images, model.Image{URL: fileUrl, ProductID: productID})
	}

	err := r.imageRepo.CreateBatchImage(ctx, images)
	if err != nil {
		return nil, err
	}

	return &productpb.ImageProductResponse{
		Images: mapper.MapImagesListToPB(images),
	}, nil
}

func (r *imageService) DeleteImageProduct(ctx context.Context, imageID string) error {
	err := r.cloudStorage.DeleteFile(ctx, imageID)
	if err != nil {
		return err
	}

	return nil
}

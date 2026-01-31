package mapper

import (
	"product-service/internal/dto"
	"product-service/internal/model"
	productpb "product-service/internal/pb/product"
)

func MapProductToImages(product *model.Product) []*productpb.Image {
	var images []*productpb.Image
	for _, image := range product.Images {
		images = append(images, &productpb.Image{
			Url: image.URL,
		})
	}
	return images
}

func MapImagesListToPB(imgs []model.Image) []*productpb.Image {
	var images []*productpb.Image
	for _, img := range imgs {
		images = append(images, &productpb.Image{
			Url: img.URL,
		})
	}
	return images
}

func MapImagePBToDTO(img *productpb.ImageRequest) *dto.ProductImageRequestDTO {
	return &dto.ProductImageRequestDTO{
		Filename:    img.Filename,
		ContentType: img.ContentType,
		Data:        img.Data,
		Size:        int64(len(img.Data)),
	}
}

func MapImagesListToDTO(imgs []*productpb.ImageRequest) []dto.ProductImageRequestDTO {
	var images []dto.ProductImageRequestDTO
	for _, img := range imgs {
		images = append(images, *MapImagePBToDTO(img))
	}
	return images
}

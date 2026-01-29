package mapper

import (
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

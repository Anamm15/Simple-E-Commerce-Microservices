package mapper

import (
	"product-service/internal/model"
	productpb "product-service/internal/pb/product"
)

func MapProductToReviews(product *model.Product) []*productpb.Review {
	var reviews []*productpb.Review
	for _, review := range product.Reviews {
		reviews = append(reviews, &productpb.Review{
			Id:       review.ID.String(),
			UserId:   review.UserID.String(),
			UserName: review.UserName,
			Rating:   review.Rating,
			Comment:  review.Comment,
			// CreatedAt: util.TimeToTimestamp(review.CreatedAt),
		})
	}
	return reviews
}

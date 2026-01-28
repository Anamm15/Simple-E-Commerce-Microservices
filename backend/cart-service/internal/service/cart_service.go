package service

import (
	"context"

	"cart-service/internal/dto"
	"cart-service/internal/helper/mapper"
	cartpb "cart-service/internal/pb/cart"
	productpb "cart-service/internal/pb/product"
	"cart-service/internal/repository"
	"cart-service/internal/util"
)

type CartService interface {
	GetCart(ctx context.Context, userID string) (*cartpb.CartDetails, error)
	AddItem(ctx context.Context, item dto.AddCartItemRequestDTO) (*cartpb.CartItem, error)
	RemoveItem(ctx context.Context, item dto.RemoveCartItemRequestDTO) error
	UpdateItem(ctx context.Context, item dto.UpdateCartItemRequestDTO) error
	ClearCart(ctx context.Context, userID string) error
}

type cartService struct {
	cartRepo      repository.CartRepository
	productClient productpb.ProductServiceClient
}

func NewCartService(cartRepo repository.CartRepository, productClient productpb.ProductServiceClient) CartService {
	return &cartService{
		cartRepo:      cartRepo,
		productClient: productClient,
	}
}

func (s *cartService) GetCart(ctx context.Context, userID string) (*cartpb.CartDetails, error) {
	cartItems, err := s.cartRepo.GetCart(ctx, userID)
	if err != nil {
		return nil, err
	}

	var productIDs []string
	for _, cartItem := range cartItems {
		productIDs = append(productIDs, cartItem.ProductID.String())
	}

	productItems, err := s.productClient.GetProductBatch(ctx, &productpb.GetProductBatchRequest{ProductIds: productIDs})
	if err != nil {
		return nil, err
	}
	return &cartpb.CartDetails{Items: mapper.MapCartListResponse(cartItems, productItems.Products)}, nil
}

func (s *cartService) AddItem(ctx context.Context, item dto.AddCartItemRequestDTO) (*cartpb.CartItem, error) {
	cartItem := item.ToModel()
	err := s.cartRepo.AddItem(ctx, cartItem)
	if err != nil {
		return nil, err
	}

	productItem, err := s.productClient.GetProductDetail(ctx, &productpb.GetProductDetailRequest{
		Id: item.ProductID,
	})

	return mapper.MapCartResponse(*cartItem, productItem.Product), nil
}

func (s *cartService) RemoveItem(ctx context.Context, item dto.RemoveCartItemRequestDTO) error {
	userID, err := util.StringToUUID(item.UserID)
	if err != nil {
		return err
	}

	productID, err := util.StringToUUID(item.ProductID)
	if err != nil {
		return err
	}

	return s.cartRepo.RemoveItem(ctx, userID, productID)
}

func (s *cartService) UpdateItem(ctx context.Context, item dto.UpdateCartItemRequestDTO) error {
	userID, err := util.StringToUUID(item.UserID)
	if err != nil {
		return err
	}

	productID, err := util.StringToUUID(item.ProductID)
	if err != nil {
		return err
	}

	return s.cartRepo.UpdateItem(ctx, userID, productID, item.Quantity)
}

func (s *cartService) ClearCart(ctx context.Context, userID string) error {
	return s.cartRepo.ClearCart(ctx, userID)
}

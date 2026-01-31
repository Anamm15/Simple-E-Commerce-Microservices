package controller

import (
	"context"

	"cart-service/internal/dto"
	cartpb "cart-service/internal/pb/cart"
	"cart-service/internal/service"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

type cartController struct {
	cartpb.UnimplementedCartServiceServer
	cartService service.CartService
}

func NewCartController(cartService service.CartService) *cartController {
	return &cartController{cartService: cartService}
}

func (c *cartController) GetCart(ctx context.Context, req *cartpb.GetCartRequest) (*cartpb.CartDetails, error) {
	return c.cartService.GetCart(ctx, req.UserId)
}

func (c *cartController) AddItem(ctx context.Context, req *cartpb.AddItemRequest) (*cartpb.CartItem, error) {
	input := dto.AddCartItemRequestDTO{
		UserID:    req.UserId,
		ProductID: req.ProductId,
		Quantity:  req.Quantity,
	}

	return c.cartService.AddItem(ctx, input)
}

func (c *cartController) RemoveItem(ctx context.Context, req *cartpb.RemoveItemRequest) (*emptypb.Empty, error) {
	input := dto.RemoveCartItemRequestDTO{
		UserID:    req.UserId,
		ProductID: req.ProductId,
	}

	err := c.cartService.RemoveItem(ctx, input)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &emptypb.Empty{}, nil
}

func (c *cartController) UpdateItem(ctx context.Context, req *cartpb.UpdateItemRequest) (*emptypb.Empty, error) {
	input := dto.UpdateCartItemRequestDTO{
		UserID:    req.UserId,
		ProductID: req.ProductId,
		Quantity:  req.Quantity,
	}

	err := c.cartService.UpdateItem(ctx, input)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &emptypb.Empty{}, nil
}

func (c *cartController) ClearCart(ctx context.Context, req *cartpb.ClearCartRequest) (*emptypb.Empty, error) {
	err := c.cartService.ClearCart(ctx, req.UserId)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &emptypb.Empty{}, nil
}

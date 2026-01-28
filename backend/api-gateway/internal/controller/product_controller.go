package controller

import (
	"net/http"

	"api-gateway/internal/dto"
	"api-gateway/internal/helpers/constants"
	productpb "api-gateway/internal/pb/product"
	"api-gateway/internal/utils"

	"github.com/gin-gonic/gin"
)

type ProductController struct {
	productClient productpb.ProductServiceClient
}

func NewProductController(productClient productpb.ProductServiceClient) *ProductController {
	return &ProductController{
		productClient: productClient,
	}
}

func (pc *ProductController) GetProducts(c *gin.Context) {
	var req dto.SearchProductRequestDTO
	if err := c.ShouldBindQuery(&req); err != nil {
		res := utils.BuildResponseFailed(constants.MsgInvalidRequest, err.Error(), nil)
		c.JSON(http.StatusBadRequest, res)
		return
	}

	grpcReq := &productpb.SearchFilterRequest{
		SearchQuery:  req.SearchQuery,
		CategorySlug: req.CategorySlug,
		MinPrice:     req.MinPrice,
		MaxPrice:     req.MaxPrice,
		Page:         req.Page,
		Limit:        req.Limit,
	}

	grpcRes, err := pc.productClient.GetProducts(c, grpcReq)
	if err != nil {
		res := utils.BuildResponseFailed(constants.MsgInternalServerError, err.Error(), nil)
		c.JSON(http.StatusInternalServerError, res)
		return
	}

	res := utils.BuildResponseSuccess(constants.MsgSuccess, grpcRes)
	c.JSON(http.StatusOK, res)
}

func (pc *ProductController) GetProductDetail(c *gin.Context) {
	productID := c.Param("id")

	grpcReq := &productpb.GetProductDetailRequest{
		Id: productID,
	}

	grpcRes, err := pc.productClient.GetProductDetail(c, grpcReq)
	if err != nil {
		res := utils.BuildResponseFailed(constants.MsgInternalServerError, err.Error(), nil)
		c.JSON(http.StatusInternalServerError, res)
		return
	}

	res := utils.BuildResponseSuccess(constants.MsgSuccess, grpcRes)
	c.JSON(http.StatusOK, res)
}

func (pc *ProductController) CreateProduct(c *gin.Context) {
	var req dto.CreateProductRequestDTO
	if err := c.ShouldBindJSON(&req); err != nil {
		res := utils.BuildResponseFailed(constants.MsgInvalidRequest, err.Error(), nil)
		c.JSON(http.StatusBadRequest, res)
		return
	}

	grpcReq := &productpb.CreateProductRequest{
		CategoryId:          req.CategoryID,
		Name:                req.Name,
		Description:         req.Description,
		Price:               req.Price,
		WeightG:             req.WeightG,
		ImageUrl:            req.ImageURL,
		AdditionalImageUrls: req.AdditionalImageURLs,
		InitialStock:        req.InitialStock,
	}

	grpcRes, err := pc.productClient.CreateProduct(c, grpcReq)
	if err != nil {
		res := utils.BuildResponseFailed(constants.MsgInternalServerError, err.Error(), nil)
		c.JSON(http.StatusInternalServerError, res)
		return
	}

	res := utils.BuildResponseSuccess(constants.MsgSuccess, grpcRes)
	c.JSON(http.StatusCreated, res)
}

func (pc *ProductController) UpdateProduct(c *gin.Context) {
	productID := c.Param("id")

	var req dto.UpdateProductRequestDTO
	if err := c.ShouldBindJSON(&req); err != nil {
		res := utils.BuildResponseFailed(constants.MsgInvalidRequest, err.Error(), nil)
		c.JSON(http.StatusBadRequest, res)
		return
	}

	grpcReq := &productpb.UpdateProductRequest{
		Id:                  productID,
		CategoryId:          req.CategoryID,
		Name:                req.Name,
		Description:         req.Description,
		Price:               req.Price,
		WeightG:             req.WeightG,
		ImageUrl:            req.ImageURL,
		AdditionalImageUrls: req.AdditionalImageURLs,
	}

	grpcRes, err := pc.productClient.UpdateProduct(c, grpcReq)
	if err != nil {
		res := utils.BuildResponseFailed(constants.MsgInternalServerError, err.Error(), nil)
		c.JSON(http.StatusInternalServerError, res)
		return
	}

	res := utils.BuildResponseSuccess(constants.MsgSuccess, grpcRes)
	c.JSON(http.StatusOK, res)
}

func (pc *ProductController) DeleteProduct(c *gin.Context) {
	productID := c.Param("id")

	grpcReq := &productpb.DeleteProductRequest{
		Id: productID,
	}

	_, err := pc.productClient.DeleteProduct(c, grpcReq)
	if err != nil {
		res := utils.BuildResponseFailed(constants.MsgInternalServerError, err.Error(), nil)
		c.JSON(http.StatusInternalServerError, res)
		return
	}

	res := utils.BuildResponseSuccess(constants.MsgSuccess, nil)
	c.JSON(http.StatusOK, res)
}

package controller

import (
	"io"
	"net/http"

	"api-gateway/internal/dto"
	"api-gateway/internal/helper/constant"
	productpb "api-gateway/internal/pb/product"
	"api-gateway/pkg/util"

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
		res := util.BuildResponseFailed(constant.MsgInvalidRequest, err.Error(), nil)
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
		res := util.BuildResponseFailed(constant.MsgInternalServerError, err.Error(), nil)
		c.JSON(http.StatusInternalServerError, res)
		return
	}

	res := util.BuildResponseSuccess(constant.MsgSuccess, grpcRes)
	c.JSON(http.StatusOK, res)
}

func (pc *ProductController) GetProductDetail(c *gin.Context) {
	productID := c.Param("id")

	grpcReq := &productpb.GetProductDetailRequest{
		Id: productID,
	}

	grpcRes, err := pc.productClient.GetProductDetail(c, grpcReq)
	if err != nil {
		res := util.BuildResponseFailed(constant.MsgInternalServerError, err.Error(), nil)
		c.JSON(http.StatusInternalServerError, res)
		return
	}

	res := util.BuildResponseSuccess(constant.MsgSuccess, grpcRes)
	c.JSON(http.StatusOK, res)
}

func (pc *ProductController) CreateProduct(c *gin.Context) {
	c.Request.Body = http.MaxBytesReader(
		c.Writer,
		c.Request.Body,
		20<<20,
	)

	var req dto.CreateProductRequestDTO
	if err := c.ShouldBind(&req); err != nil {
		res := util.BuildResponseFailed(constant.MsgInvalidRequest, err.Error(), nil)
		c.JSON(http.StatusBadRequest, res)
		return
	}

	form, err := c.MultipartForm()
	if err != nil {
		res := util.BuildResponseFailed("invalid multipart form", err.Error(), nil)
		c.JSON(http.StatusBadRequest, res)
		return
	}

	files := form.File["images"] // [] *multipart.FileHeader
	if len(files) == 0 {
		res := util.BuildResponseFailed("images required", "image required minimal 1", nil)
		c.JSON(http.StatusBadRequest, res)
		return
	}

	var grpcImages []*productpb.ImageRequest

	for _, fileHeader := range files {
		file, err := fileHeader.Open()
		if err != nil {
			continue
		}

		bytes, err := io.ReadAll(file)
		file.Close()
		if err != nil {
			continue
		}

		grpcImages = append(grpcImages, &productpb.ImageRequest{
			Filename:    fileHeader.Filename,
			ContentType: fileHeader.Header.Get("Content-Type"),
			Data:        bytes,
		})
	}

	grpcReq := &productpb.CreateProductRequest{
		CategoryId:     req.CategoryID,
		Name:           req.Name,
		Description:    req.Description,
		Price:          req.Price,
		WeightG:        req.WeightG,
		InitialStock:   req.InitialStock,
		Thumbnail:      grpcImages[0],
		AdditionalImgs: grpcImages[1:],
	}

	grpcRes, err := pc.productClient.CreateProduct(c, grpcReq)
	if err != nil {
		res := util.BuildResponseFailed(constant.MsgInternalServerError, err.Error(), nil)
		c.JSON(http.StatusInternalServerError, res)
		return
	}

	res := util.BuildResponseSuccess(constant.MsgSuccess, grpcRes)
	c.JSON(http.StatusCreated, res)
}

func (pc *ProductController) UpdateProduct(c *gin.Context) {
	productID := c.Param("id")

	var req dto.UpdateProductRequestDTO
	if err := c.ShouldBindJSON(&req); err != nil {
		res := util.BuildResponseFailed(constant.MsgInvalidRequest, err.Error(), nil)
		c.JSON(http.StatusBadRequest, res)
		return
	}

	grpcReq := &productpb.UpdateProductRequest{
		Id:          productID,
		CategoryId:  req.CategoryID,
		Name:        req.Name,
		Description: req.Description,
		Price:       req.Price,
		WeightG:     req.WeightG,
	}

	grpcRes, err := pc.productClient.UpdateProduct(c, grpcReq)
	if err != nil {
		res := util.BuildResponseFailed(constant.MsgInternalServerError, err.Error(), nil)
		c.JSON(http.StatusInternalServerError, res)
		return
	}

	res := util.BuildResponseSuccess(constant.MsgSuccess, grpcRes)
	c.JSON(http.StatusOK, res)
}

func (pc *ProductController) DeleteProduct(c *gin.Context) {
	productID := c.Param("id")

	grpcReq := &productpb.DeleteProductRequest{
		Id: productID,
	}

	_, err := pc.productClient.DeleteProduct(c, grpcReq)
	if err != nil {
		res := util.BuildResponseFailed(constant.MsgInternalServerError, err.Error(), nil)
		c.JSON(http.StatusInternalServerError, res)
		return
	}

	res := util.BuildResponseSuccess(constant.MsgSuccess, nil)
	c.JSON(http.StatusOK, res)
}

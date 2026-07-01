package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"

	entity "github.com/ascend/phoenix-multichannel-mkp/domain"
	"github.com/ascend/phoenix-multichannel-mkp/internal/service"
	"github.com/ascend/phoenix-multichannel-mkp/pkg/response"
)

type ProductHandler struct {
	products service.ProductService
}

func NewProductHandler(products service.ProductService) *ProductHandler {
	return &ProductHandler{products: products}
}

func (h *ProductHandler) RegisterRoutes(rg *gin.RouterGroup) {
	products := rg.Group("/products")
	products.POST("", h.Create)
	products.GET("", h.List)
	products.GET("/:id", h.GetByID)
}

type createProductRequest struct {
	SellerID    string           `json:"seller_id" binding:"required"`
	Name        string           `json:"name"      binding:"required"`
	Description string           `json:"description"`
	Price       float64          `json:"price"     binding:"required,gte=0"`
	Stock       int              `json:"stock"     binding:"gte=0"`
	Channels    []entity.Channel `json:"channels"  binding:"required,min=1"`
}

func (h *ProductHandler) Create(c *gin.Context) {
	var req createProductRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	out, err := h.products.Create(c.Request.Context(), service.CreateProductInput{
		SellerID:    req.SellerID,
		Name:        req.Name,
		Description: req.Description,
		Price:       req.Price,
		Stock:       req.Stock,
		Channels:    req.Channels,
	})
	if err != nil {
		response.Fail(c, err)
		return
	}
	response.Created(c, out.Product)
}

func (h *ProductHandler) GetByID(c *gin.Context) {
	p, err := h.products.Get(c.Request.Context(), c.Param("id"))
	if err != nil {
		response.Fail(c, err)
		return
	}
	response.OK(c, p)
}

func (h *ProductHandler) List(c *gin.Context) {
	out, err := h.products.List(c.Request.Context(), service.ListProductsInput{
		SellerID: c.Query("seller_id"),
		Channel:  entity.Channel(c.Query("channel")),
		Status:   entity.ProductStatus(c.Query("status")),
	})
	if err != nil {
		response.Fail(c, err)
		return
	}
	response.OK(c, response.PaginatedResponse{
		Items: out.Products,
		Total: out.Total,
		Page:  out.Page,
		Limit: out.Limit,
	})
}

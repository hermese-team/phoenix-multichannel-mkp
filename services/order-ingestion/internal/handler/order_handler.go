package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"

	entity "github.com/ascend/phoenix-multichannel-mkp/domain"
	"github.com/ascend/phoenix-multichannel-mkp/internal/service"
	"github.com/ascend/phoenix-multichannel-mkp/pkg/response"
)

type OrderHandler struct {
	orders service.OrderService
}

func NewOrderHandler(orders service.OrderService) *OrderHandler {
	return &OrderHandler{orders: orders}
}

func (h *OrderHandler) RegisterRoutes(rg *gin.RouterGroup) {
	orders := rg.Group("/orders")
	orders.POST("", h.Create)
}

type createOrderItemRequest struct {
	ProductID string `json:"product_id" binding:"required"`
	Quantity  int    `json:"quantity"   binding:"required,min=1"`
}

type createOrderRequest struct {
	BuyerID string                   `json:"buyer_id" binding:"required"`
	Channel entity.Channel           `json:"channel"  binding:"required"`
	Items   []createOrderItemRequest  `json:"items"    binding:"required,min=1"`
}

func (h *OrderHandler) Create(c *gin.Context) {
	var req createOrderRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	items := make([]service.CreateOrderItemInput, len(req.Items))
	for i, item := range req.Items {
		items[i] = service.CreateOrderItemInput{
			ProductID: item.ProductID,
			Quantity:  item.Quantity,
		}
	}

	out, err := h.orders.Create(c.Request.Context(), service.CreateOrderInput{
		BuyerID: req.BuyerID,
		Channel: req.Channel,
		Items:   items,
	})
	if err != nil {
		response.Fail(c, err)
		return
	}
	response.Created(c, out.Order)
}

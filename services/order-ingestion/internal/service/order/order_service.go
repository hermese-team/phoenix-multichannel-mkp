package order

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/google/uuid"

	entity "github.com/ascend/phoenix-multichannel-mkp/domain"
	"github.com/ascend/phoenix-multichannel-mkp/internal/repository"
	"github.com/ascend/phoenix-multichannel-mkp/internal/service"
	domainerrors "github.com/ascend/phoenix-multichannel-mkp/pkg/errors"
	"github.com/ascend/phoenix-multichannel-mkp/pkg/logger"
)

type OrderService struct {
	orderRepo         repository.OrderRepository
	productRepo       repository.ProductRepository
	publisher         service.EventPublisher
	orderCreatedTopic string
}

func NewOrderService(
	orderRepo repository.OrderRepository,
	productRepo repository.ProductRepository,
	publisher service.EventPublisher,
	orderCreatedTopic string,
) *OrderService {
	return &OrderService{
		orderRepo:         orderRepo,
		productRepo:       productRepo,
		publisher:         publisher,
		orderCreatedTopic: orderCreatedTopic,
	}
}

func (s *OrderService) Create(ctx context.Context, input service.CreateOrderInput) (*service.CreateOrderOutput, error) {
	if len(input.Items) == 0 {
		return nil, domainerrors.New("EMPTY_ORDER", "order must have at least one item", domainerrors.ErrInvalidInput)
	}

	var orderItems []entity.OrderItem
	var sellerID string

	for _, item := range input.Items {
		product, err := s.productRepo.GetByID(ctx, item.ProductID)
		if err != nil {
			return nil, fmt.Errorf("fetching product %s: %w", item.ProductID, err)
		}
		if product.Status != entity.ProductStatusActive {
			return nil, domainerrors.New("PRODUCT_UNAVAILABLE", fmt.Sprintf("product %s is not available", item.ProductID), domainerrors.ErrInvalidInput)
		}
		if err := product.DeductStock(item.Quantity); err != nil {
			return nil, err
		}
		if err := s.productRepo.Update(ctx, product); err != nil {
			return nil, fmt.Errorf("updating product stock: %w", err)
		}
		if sellerID == "" {
			sellerID = product.SellerID
		}
		orderItems = append(orderItems, entity.OrderItem{
			ProductID: product.ID,
			Name:      product.Name,
			Price:     product.Price,
			Quantity:  item.Quantity,
		})
	}

	order, err := entity.NewOrder(input.BuyerID, sellerID, input.Channel, orderItems)
	if err != nil {
		return nil, err
	}
	order.ID = uuid.NewString()

	if err := s.orderRepo.Create(ctx, order); err != nil {
		return nil, err
	}

	s.publishOrderCreated(ctx, order)

	return &service.CreateOrderOutput{Order: order}, nil
}

// publishOrderCreated emits an order.created event. A publish failure is logged
// but does not fail the order, which is already persisted.
func (s *OrderService) publishOrderCreated(ctx context.Context, order *entity.Order) {
	if s.publisher == nil {
		return
	}
	payload, err := json.Marshal(order)
	if err != nil {
		logger.Error("marshal order event", "order_id", order.ID, "error", err)
		return
	}
	if err := s.publisher.Publish(ctx, s.orderCreatedTopic, order.ID, payload); err != nil {
		logger.Error("publish order.created", "order_id", order.ID, "error", err)
	}
}

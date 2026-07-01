package product

import (
	"context"

	"github.com/google/uuid"

	entity "github.com/ascend/phoenix-multichannel-mkp/domain"
	"github.com/ascend/phoenix-multichannel-mkp/internal/repository"
	"github.com/ascend/phoenix-multichannel-mkp/internal/service"
	"github.com/ascend/phoenix-multichannel-mkp/pkg/logger"
)

type ProductService struct {
	productRepo  repository.ProductRepository
	marketplaces service.MarketplaceRegistry
}

func NewProductService(repo repository.ProductRepository, marketplaces service.MarketplaceRegistry) *ProductService {
	return &ProductService{productRepo: repo, marketplaces: marketplaces}
}

func (s *ProductService) Create(ctx context.Context, input service.CreateProductInput) (*service.CreateProductOutput, error) {
	product, err := entity.NewProduct(
		input.SellerID,
		input.Name,
		input.Description,
		input.Price,
		input.Stock,
		input.Channels,
	)
	if err != nil {
		return nil, err
	}
	product.ID = uuid.NewString()

	if err := s.productRepo.Create(ctx, product); err != nil {
		return nil, err
	}

	s.syncToMarketplaces(ctx, product)

	return &service.CreateProductOutput{Product: product}, nil
}

// syncToMarketplaces pushes the product to each external channel it targets.
// Failures are logged but do not fail creation; channels without a gateway
// (e.g. web/mobile/api) are skipped.
func (s *ProductService) syncToMarketplaces(ctx context.Context, p *entity.Product) {
	if s.marketplaces == nil {
		return
	}
	for _, ch := range p.Channels {
		gw, err := s.marketplaces.Get(ch)
		if err != nil {
			continue
		}
		externalID, err := gw.PushProduct(ctx, p)
		if err != nil {
			logger.Error("push product to marketplace", "channel", ch, "product_id", p.ID, "error", err)
			continue
		}
		logger.Info("pushed product to marketplace", "channel", ch, "product_id", p.ID, "external_id", externalID)
	}
}

func (s *ProductService) Get(ctx context.Context, id string) (*entity.Product, error) {
	return s.productRepo.GetByID(ctx, id)
}

func (s *ProductService) List(ctx context.Context, input service.ListProductsInput) (*service.ListProductsOutput, error) {
	if input.Page <= 0 {
		input.Page = 1
	}
	if input.Limit <= 0 || input.Limit > 100 {
		input.Limit = 20
	}

	products, total, err := s.productRepo.List(ctx, repository.ProductFilter{
		SellerID: input.SellerID,
		Channel:  input.Channel,
		Status:   input.Status,
		Page:     input.Page,
		Limit:    input.Limit,
	})
	if err != nil {
		return nil, err
	}
	return &service.ListProductsOutput{
		Products: products,
		Total:    total,
		Page:     input.Page,
		Limit:    input.Limit,
	}, nil
}

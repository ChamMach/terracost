package costestimation

import (
	"context"

	"github.com/cycloidio/cost-estimation/price"
	"github.com/cycloidio/cost-estimation/product"
)

//go:generate go run github.com/golang/mock/mockgen -destination=mock/ingester.go -mock_names=Ingester=Ingester -package mock github.com/cycloidio/cost-estimation Ingester
//go:generate go run github.com/golang/mock/mockgen -destination=mock/backend.go -mock_names=Backend=Backend -package mock github.com/cycloidio/cost-estimation Backend

// Ingester represents a vendor-specific mechanism to load pricing data.
type Ingester interface {
	// Ingest downloads pricing data from a cloud provider and sends prices with their associated products
	// on the returned channel.
	Ingest(ctx context.Context, chSize int) <-chan *price.WithProduct

	// Err returns any potential error.
	Err() error
}

// Backend represents a storage method used to store pricing data. It must include concrete implementations
// of all repositories.
type Backend interface {
	Product() product.Repository
	Price() price.Repository
}

// IngestPricing uses the Ingester to load the pricing data and stores it into the Backend.
func IngestPricing(ctx context.Context, backend Backend, ingester Ingester) error {
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	skuProductID := make(map[string]product.ID)
	for pp := range ingester.Ingest(ctx, 8) {
		if id, ok := skuProductID[pp.Product.SKU]; ok {
			pp.Product.ID = id
		} else {
			var err error
			pp.Product.ID, err = backend.Product().Upsert(ctx, pp.Product)
			if err != nil {
				return err
			}
			skuProductID[pp.Product.SKU] = pp.Product.ID
		}

		if _, err := backend.Price().Upsert(ctx, pp); err != nil {
			return err
		}
	}

	if err := ingester.Err(); err != nil {
		return err
	}
	return nil
}

package saveasset

import (
	"context"

	domain "markettracker.com/replicator/internal"
)

type AssetService struct {
	assetRepository domain.AssetRepository
}

func NewAssetService(assetRepository domain.AssetRepository) AssetService {
	return AssetService{
		assetRepository: assetRepository,
	}
}

func (a AssetService) Save(ctx context.Context, id string, date string, exchange string, price float32) error {
	asset, err := domain.NewAsset(id, date, exchange, price)
	if err != nil {
		return err
	}
	return a.assetRepository.Save(ctx, asset)
}

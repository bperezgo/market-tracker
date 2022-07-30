package domain

import "context"

// TODO: Use mock library to manage the next tests
type AssetRepositoryMock struct{}

func (r *AssetRepositoryMock) Save(ctx context.Context, asset Asset) error {
	return nil
}

package replicate

import (
	"context"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
	"markettracker.com/pkg/event"
	"markettracker.com/tracker/internal/domain"
)

func Test_Should_Send_The_Events(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockBus := event.NewMockBus(ctrl)
	// Expect one event in the array, and is called 1 times
	mockBus.EXPECT().Publish(gomock.Any(), gomock.Len(1)).Times(1)

	ctx := context.Background()

	repo := domain.AssetRepositoryMock{}
	replicator := New(&repo, mockBus)
	err := replicator.Replicate(ctx, uuid.NewString(), time.Date(2021, 8, 15, 14, 30, 45, 100, time.Local), "exchange", 123.23)
	require.NoError(t, err, "replicating error")

}

package repository

import (
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"markettracker.com/tracker/internal/domain"
)

func initConfig() PostgresqlConfig {
	return PostgresqlConfig{
		Host:     "localhost",
		Port:     5432,
		User:     "tracker",
		Password: "dummy-pasword",
		Dbname:   "tracker",
	}
}

func Test_Validate_Connection_With_Postgresql(t *testing.T) {
	_, err := NewPostgresql("dummy_asset", initConfig())
	assert.NoError(t, err, "error was not expected")
}

func Test_Save_Value(t *testing.T) {
	repo, err := NewPostgresql("dummy_asset", initConfig())
	require.NoError(t, err, "error was not expected in the constructor")

	asset, err := domain.NewAsset(uuid.NewString(), time.Now(), "dummy_asset", 123.4)
	require.NoError(t, err, "failed the business rules")

	err = repo.Save(asset)
	assert.NoError(t, err, "unexpected error when is saving the asset ina dummy table")
}

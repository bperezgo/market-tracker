package domain

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Should_Asset_Date(t *testing.T) {
	t.Run("return nil error", func(t *testing.T) {
		_, err := NewDate("2022-07-04T15:48:58.4Z")
		assert.Equal(t, nil, err, "expect no error")
	})

	t.Run("return ErrDateDoesNotMatchTheRightFormat", func(t *testing.T) {
		_, err := NewDate("")
		assert.Equal(t, ErrDateDoesNotMatchTheRightFormat, err, "expect error")
	})
}

package service

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestGetAggregatedCosts(t *testing.T) {
	t.Run("GroupByService", func(t *testing.T) {
		costs, err := GetAggregatedCosts("service")

		assert.NoError(t, err)
		assert.NotNil(t, costs)
		assert.Equal(t, 2, len(costs))
		assert.Equal(t, "AmazonEC2", costs[0].GroupKey)
	})

	t.Run("GroupByCloud", func(t *testing.T) {
		costs, err := GetAggregatedCosts("cloud")

		assert.NoError(t, err)
		assert.NotNil(t, costs)
		assert.Equal(t, 2, len(costs))
		assert.Equal(t, "AWS", costs[0].GroupKey)
	})
}

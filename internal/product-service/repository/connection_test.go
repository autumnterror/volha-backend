package repository

import (
	"testing"

	"github.com/autumnterror/volha-backend/internal/product-service/config"
	"github.com/stretchr/testify/assert"
)

func TestConnectDisconnect(t *testing.T) {
	d, err := NewConnect(config.Test())
	assert.NoError(t, err)
	assert.NoError(t, d.Disconnect())
}

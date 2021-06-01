package config_test

import (
	"testing"

	"github.com/pursuit/gateway/internal/config"

	"github.com/stretchr/testify/assert"
)

func TestInstance(t *testing.T) {
	assert.NotNil(t, config.Instance("./"))
	assert.NotNil(t, config.Instance("./mock/no_prefix/"))
	assert.Panics(t, func() { config.Instance("../url_factory/imagination/") })
	assert.Panics(t, func() { config.Instance("../url_factory/invalid_format/") })
}

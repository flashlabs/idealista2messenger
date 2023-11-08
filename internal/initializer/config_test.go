package initializer_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/flashlabs/idealista2messenger/internal/initializer"
	_ "github.com/flashlabs/idealista2messenger/internal/test"
)

func TestCfg(t *testing.T) {
	var config = initializer.Cfg("config")

	t.Run("test application config", func(t *testing.T) {
		assert.NotEmpty(t, config.Application.Name)
	})

	t.Run("test Google config", func(t *testing.T) {
		assert.NotEmpty(t, config.Google.AccessTokenFile)
		assert.NotEmpty(t, config.Google.CredentialsFile)
	})

	t.Run("test Messenger config", func(t *testing.T) {
		assert.NotEmpty(t, config.Messenger.PageAccessTokenFile)
	})
}

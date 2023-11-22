package config

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

var strConfig = `
TELEGRAM_BOT_ADMINS=1
TELEGRAM_TOKEN=2
PINTEREST_TOKEN=3
PINTEREST_CLIENT_ID=4
PINTEREST_CLIENT_SECRET=5
DATABASE_URL=6
`

func TestConfig(t *testing.T) {
	t.Parallel()

	testcase := []struct {
		name     string
		setupEnv func(t *testing.T) (*Config, error)
		checkEnv func(t *testing.T, config *Config, err error)
	}{
		{
			name: "test",
			setupEnv: func(t *testing.T) (*Config, error) {
				file, err := os.CreateTemp("", "temp-env")
				assert.NoError(t, err)

				_, err = file.WriteString(strConfig)
				assert.NoError(t, err)

				defer file.Close()

				config, err := NewConfig()
				assert.NoError(t, err)

				return config, nil
			},
			checkEnv: func(t *testing.T, config *Config, err error) {
				assert.NoError(t, err)

				assert.Equal(t, "1", config.TgToken)
				assert.Equal(t, "2", config.ApiHash)
				assert.Equal(t, "3", config.TgEnabled)
				assert.Equal(t, "4", config.TgBotAdmins)
				assert.Equal(t, "5", config.DatabaseURL)
				assert.Equal(t, "6", config.PUserToken)
				assert.Equal(t, "7", config.PClientId)
				assert.Equal(t, "8", config.PClientSecret)
			},
		},
	}
	for _, tt := range testcase {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			tt.setupEnv(t)

			cfg, err := tt.setupEnv(t)
			tt.checkEnv(t, cfg, err)
		})
	}
}

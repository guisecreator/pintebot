package config

import (
	"github.com/celestix/gotgproto"
	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
	"os"
	"path/filepath"
)

type Config struct {
	TgEnabled     	bool
	TgBotAdmins   	string `required:"true"   envconfig:"TELEGRAM_BOT_ADMINS"`
	TgToken       	string `required:"true"   envconfig:"TELEGRAM_TOKEN"`
	AppID         	int    `required:"false"  envconfig:"APP_ID"`
	ApiHash         string `required:"false"  envconfig:"API_HASH"`
	PAccessToken    string `required:"false"  envconfig:"PINTEREST_ACCESS_TOKEN"`
	PClientId     	string `required:"false"  envconfig:"PINTEREST_CLIENT_ID"`
	PClientSecret 	string `required:"false"  envconfig:"PINTEREST_CLIENT_SECRET"`
	DatabaseURL   	string `required:"false"  envconfig:"DATABASE_URL"`
	CType        	  gotgproto.ClientType
	CmdText       	CommandsText
}

var processEnv = envconfig.Process

func NewConfig() (*Config, error) {
	var (
		newCfg Config
		errCfg error
	)

	wd, errCfg := os.Getwd()
	if errCfg != nil {
		return nil, errCfg
	}

	envPath := filepath.Join(wd, ".env")

	_ = godotenv.Overload(envPath)
	if errCfg = processEnv("", &newCfg); errCfg != nil {
		return nil, errCfg
	}

	return &newCfg, nil
}

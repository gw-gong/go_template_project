package service01

import (
	"fmt"
	"sync"

	"github.com/gw-gong/go-template-project/config/types"
	"github.com/gw-gong/go-template-project/internal/pkg/components/component01"
	"github.com/gw-gong/go-template-project/internal/pkg/components/component02"

	"github.com/gw-gong/gwkit-go/hot_cfg"
	"github.com/gw-gong/gwkit-go/log"
)

var (
	Cfg  *Config
	once sync.Once
)

type Config struct {
	*hot_cfg.BaseConfig
	HttpServer  *types.HttpServer              `yaml:"http_server" mapstructure:"http_server"`
	Logger      *log.LoggerConfig              `yaml:"logger" mapstructure:"logger"`
	Component01 component01.Component01Options `yaml:"component01" mapstructure:"component01"`
	Component02 component02.Component02Options `yaml:"component02" mapstructure:"component02"`
}

func InitConfig(filePath, fileName, fileType string) error {
	var err error
	once.Do(func() {
		Cfg = &Config{}
		Cfg.BaseConfig, err = hot_cfg.NewBaseConfig(
			hot_cfg.WithLocalConfig(filePath, fileName, fileType),
		)
		if err != nil {
			err = fmt.Errorf("init base config failed: %w", err)
		}
	})
	return err
}

func (c *Config) LoadConfig() {
	if err := c.BaseConfig.Viper.Unmarshal(&c); err != nil {
		log.Error("unmarshal config failed", log.Err(err))
		return
	}

	if c.Logger == nil {
		log.Error("logger config is nil, use default logger config")
		c.Logger = log.NewDefaultLoggerConfig()
	}

	log.Info("LoadConfig", log.Any("config", c))
}

package localcfg

import (
	"github.com/gw-gong/go-template-project/internal/config/types"

	"github.com/gw-gong/gwkit-go/hotcfg"
	"github.com/gw-gong/gwkit-go/log"
)

type Config struct {
	hotcfg.BaseConfigCapable
	HttpServer *types.HttpServer `yaml:"http_server" mapstructure:"http_server"`
	Logger     *log.LoggerConfig `yaml:"logger" mapstructure:"logger"`
}

func (c *Config) LoadConfig() {
	if err := c.Unmarshal(&c); err != nil {
		log.Error("unmarshal config failed", log.Err(err))
		return
	}

	log.Info("LoadConfig", log.Any("config", c))
}

func NewConfig(localConfigOption *hotcfg.LocalConfigOption) (config *Config, err error) {
	config = &Config{}
	config.BaseConfigCapable, err = hotcfg.NewLocalBaseConfigCapable(localConfigOption)
	return config, err
}

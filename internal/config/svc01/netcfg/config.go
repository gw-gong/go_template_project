package netcfg

import (
	"github.com/gw-gong/boilerplate-go/internal/pkg/db/mysql"

	"github.com/gw-gong/gwkit-go/hotcfg"
	"github.com/gw-gong/gwkit-go/log"
)

type Config struct {
	hotcfg.BaseConfigCapable
	Test01DbManager *mysql.Test01DbManagerOptions `yaml:"test01_db_manager" mapstructure:"test01_db_manager"`
	Test02DbManager *mysql.Test02DbManagerOptions `yaml:"test02_db_manager" mapstructure:"test02_db_manager"`
}

func (c *Config) LoadConfig() {
	if err := c.Unmarshal(&c); err != nil {
		log.Error("unmarshal config failed", log.Err(err))
		return
	}

	log.Info("LoadConfig", log.Any("config", c))
}

func NewConfig(consulConfigOption *hotcfg.ConsulConfigOption) (config *Config, err error) {
	config = &Config{}
	config.BaseConfigCapable, err = hotcfg.NewConsulBaseConfigCapable(consulConfigOption)
	return config, err
}

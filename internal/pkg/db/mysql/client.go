package mysql

import (
	"context"

	"github.com/gw-gong/gwkit-go/log"
)

type MysqlClientOptions struct {
	Host     string `yaml:"host" mapstructure:"host"`
	Port     int    `yaml:"port" mapstructure:"port"`
	User     string `yaml:"user" mapstructure:"user"`
	Password string `yaml:"password" mapstructure:"password"`
	Database string `yaml:"database" mapstructure:"database"`
}

type (
	Test01DbManagerOptions MysqlClientOptions
	Test02DbManagerOptions MysqlClientOptions
)

type mysqlClient struct {
	options *MysqlClientOptions
}

func newMysqlClient(options *MysqlClientOptions) (*mysqlClient, error) {
	log.Info("newMysqlClient", log.Any("options", options))
	return &mysqlClient{
		options: options,
	}, nil
}

func NewTest01DbManager(options *Test01DbManagerOptions) (Test01DbManager, error) {
	return newMysqlClient((*MysqlClientOptions)(options))
}

func NewTest02DbManager(options *Test02DbManagerOptions) (Test02DbManager, error) {
	return newMysqlClient((*MysqlClientOptions)(options))
}

func (m *mysqlClient) Close() error {
	log.Info("mysqlClient Close")
	return nil
}

func (m *mysqlClient) Setxxxx(ctx context.Context) {
	log.Infoc(ctx, "mysqlClient Setxxxx")
}

func (m *mysqlClient) Getxxxx(ctx context.Context) {
	log.Infoc(ctx, "mysqlClient Getxxxx")
}

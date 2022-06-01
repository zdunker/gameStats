package config

import (
	"github.com/spf13/viper"
	"github.com/uptrace/bun"
)

var c conf

type conf struct {
	Version string

	DotaConfig         *dotaConfig           `mapstructure:"dota_config"`
	InfrastructureConf *infrastructureConfig `mapstructure:"infrastructure"`

	// golbal instances after initializing the config
	infras infrastructures
}

type infrastructureConfig struct {
	DBConfig *dbConfig `mapstructure:"db_config"`
}

func (infraConf infrastructureConfig) initialize() error {
	return infraConf.DBConfig.connect()
}

// -----------------------------------------------------------------------------

func Load(v *viper.Viper) error {
	err := v.Unmarshal(&c)
	if err != nil {
		//TODO: log error
		return err
	}

	err = c.InfrastructureConf.initialize()
	if err != nil {
		//TODO: log error
		return err
	}

	return nil
}

func GetConfig() *conf {
	return &c
}

func (c *conf) GetVersion() string {
	return c.Version
}

func (c *conf) Dota() *dotaConfig {
	return c.DotaConfig
}

// ------ infrastructure instances ---------

type infrastructures struct {
	db *bun.DB
}

func GetInfrastructure() *infrastructures {
	return &c.infras
}

func (infras *infrastructures) DB() *bun.DB {
	return infras.db
}

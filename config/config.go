package config

var c conf

type conf struct {
	DotaConfig     *dotaConfig
	Infrastructure *infrastructure
}

type infrastructure struct {
	DBConfig *dbConfig
}

// -----------------------------------------------------------------------------

func LoadConfig(filepath string) error {
	dotaconfig, err := loadDotaConfig(filepath)
	if err != nil {
		return err
	}
	c.DotaConfig = dotaconfig

	return nil
}

func GetConfig() *conf {
	return &c
}

func (c *conf) Dota() *dotaConfig {
	return c.DotaConfig
}

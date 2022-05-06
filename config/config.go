package config

var c conf

type conf struct {
	dotaConfig *dotaConfig
}

func LoadConfig(filepath string) error {
	dotaconfig, err := loadDotaConfig(filepath)
	if err != nil {
		return err
	}
	c.dotaConfig = dotaconfig

	return nil
}

func GetConfig() *conf {
	return &c
}

func (c *conf) Dota() *dotaConfig {
	return c.dotaConfig
}

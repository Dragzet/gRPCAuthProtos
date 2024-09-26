package CustomLib

import "github.com/ilyakaznacheev/cleanenv"

type Config struct {
	Name string `yaml:"name"`
}

func MustLoad() *Config {
	var cfg *Config

	if err := cleanenv.ReadConfig("config.yaml", cfg); err != nil {
		return nil
	}

	return cfg
}

package conf

type Config struct {
}

func NewConfig() (*Config, error) {
	return &Config{}, nil
}

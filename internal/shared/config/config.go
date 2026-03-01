package config

type Config struct {
	AppPort      string  `env:"APP_PORT"`
	InterestRate float64 `env:"INTEREST_RATE"`
	InterestType string  `env:"INTEREST_TYPE"`
	TenorWeeks   int16   `env:"TENOR_WEEKS"`
	Database     DatabaseConfig
}

type DatabaseConfig struct {
	Host                   string `env:"DB_HOST"`
	Port                   string `env:"DB_PORT"`
	User                   string `env:"DB_USER"`
	Password               string `env:"DB_PASSWORD"`
	Name                   string `env:"DB_NAME"`
	MaxIdleConns           int    `env:"DB_MAX_IDLE_CONNS"`
	MaxOpenConns           int    `env:"DB_MAX_OPEN_CONNS"`
	ConnMaxLifetimeMinutes int    `env:"DB_CONN_MAX_LIFETIME_MINUTES"`
}

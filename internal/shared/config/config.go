package config

type Config struct {
	AppPort      string
	InterestRate float64
	InterestType string
	TenorWeeks   int16
	Database     DatabaseConfig
}

type DatabaseConfig struct {
	Host                   string
	Port                   string
	User                   string
	Password               string
	Name                   string
	MaxIdleConns           int
	MaxOpenConns           int
	ConnMaxLifetimeMinutes int
}

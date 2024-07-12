package config

// DatabaseConfig provides a structure containing the fields needed to connect to postgres.
type DatabaseConfig struct {
	Host     string `envconfig:"DATABASE_HOST" required:"true"`
	Port     int    `envconfig:"DATABASE_PORT" required:"true"`
	Database string `envconfig:"DATABASE_NAME" required:"true"`
	Username string `envconfig:"DATABASE_USERNAME" required:"true"`
	Password string `envconfig:"DATABASE_PASSWORD" required:"true"`
	Sslmode  string `envconfig:"DATABASE_SSLMODE" required:"true"`
}

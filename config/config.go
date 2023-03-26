package config

import (
	"flag"
	"os"

	_ "github.com/joho/godotenv/autoload"
)

type Config struct {
	BotToken string
	Env      string
	Log      struct {
		Level string
	}
	DB struct {
		DSN          string
		MaxOpenConns int
		MaxIdleConns int
		MaxIdleTime  string
	}
}

var (
	dsn      = os.Getenv("DB_DSN")
	botToken = os.Getenv("TG_BOT_TOKEN")
)

// NewConfig returns app config.
func NewConfig() (*Config, error) {
	cfg := &Config{}
	flag.StringVar(&cfg.Env, "env", "development", "Environment (development|staging|production)")

	flag.IntVar(&cfg.DB.MaxOpenConns, "db-max-open-conns", 25, "PostgreSQL max open connections")
	flag.IntVar(&cfg.DB.MaxIdleConns, "db-max-idle-conns", 25, "PostgreSQL max idle connections")
	flag.StringVar(&cfg.DB.MaxIdleTime, "db-max-idle-time", "15m", "PostgreSQL max connection idle time")
	flag.StringVar(&cfg.DB.DSN, "db-dsn", dsn, "PostgreSQL DSN")

	flag.StringVar(&cfg.Log.Level, "log-level", "DEBUG", "Log level")

	flag.StringVar(&cfg.BotToken, "bot-token", botToken, "Telegram bot token")

	flag.Parse()

	return cfg, nil
}

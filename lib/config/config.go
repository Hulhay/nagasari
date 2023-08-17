package config

import (
	"os"
	"strconv"
	"time"

	"github.com/hulhay/nagasari/lib/database"
	"github.com/joho/godotenv"
)

type Config struct {
	httpServerAddress string
	dbSqlx            database.SqlxDatabase

	expJwtToken  time.Duration
	jwtSecretKey string
}

func NewConfig() *Config {
	err := godotenv.Load(".env")
	if err != nil {
		panic("Failed to load env file")
	}

	cfg := new(Config)

	cfg.ConnectSqlxDB()
	return cfg
}

func (c *Config) HTTPServerAddress() string {
	c.httpServerAddress = os.Getenv("HTTP_SERVER_ADDRESS")
	return c.httpServerAddress
}

func (c *Config) ConnectSqlxDB() {
	c.dbSqlx = database.InitSQLX()
}

func (c *Config) DisconnectDB() {
	c.dbSqlx.Close()
}

func (c *Config) DB() database.SqlxDatabase {
	return c.dbSqlx
}

func (c *Config) JWTSecretKey() string {
	c.jwtSecretKey = os.Getenv("JWT_SECRET_KEY")
	return c.jwtSecretKey
}

func (c *Config) ExpJwtToken() time.Duration {
	val, _ := strconv.Atoi(os.Getenv("EXP_JWT_TOKEN"))
	c.expJwtToken = time.Duration(val) * time.Hour
	return c.expJwtToken
}

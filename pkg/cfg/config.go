package cfg

import (
	"github.com/kelseyhightower/envconfig"
	"gorm.io/gorm"
)

// Config to hold global environment variables
type Config struct {
	APIVersion        string `envconfig:"AIFTD_API_VERSION" default:"v1"`
	APIPort           string `envconfig:"AIFTD_API_PORT" default:"1111"`
	Debug             bool   `envconfig:"AIFTD_DEBUG" default:"true"`
	ExpirationInHours int    `envconfig:"AIFTD_EXPIRATION_IN_HOURS" default:"24"`
	Dsn               string `envconfig:"AIFTD_MYSQL_DSN" default:"root:amede2a2@tcp(mysql:3306)/aifitted?charset=utf8mb4&parseTime=True&loc=Local"`
	Secret            string `envconfig:"AIFTD_SECRET" dafault:"My Awesome AI Fitted"`
}

// Spec is the specification of this instance
var Spec *Config

// Db is for dabase connection
var Db *gorm.DB

// Init config
func Init() error {
	var c Config
	err := envconfig.Process("aiftd", &c)
	if err != nil {
		return err
	}
	Spec = &c

	return nil
}

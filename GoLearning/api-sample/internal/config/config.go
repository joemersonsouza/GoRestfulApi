package config

import (
	"fmt"
	"io/ioutil"
	"os"
	"time"

	"github.com/kelseyhightower/envconfig"
	"gopkg.in/yaml.v2"
)

// ServiceConfig ...
type ServiceConfig struct {
	Environment string       `envconfig:"GAPI_ENVIRONMENT" yaml:"environment" json:"environment" split_words:"true"`
	Server      serverInfo   `yaml:"server" json:"server"`
	Cors        corsInfo     `yaml:"cors" json:"cors"`
	DB          dataBaseInfo `yaml:"databaseInfo" json:"databaseInfo"`
	RestAPI     restAPIInfo  `yaml:"rest_api" json:"rest_api"`
	ServiceName string       `envconfig:"GAPI_SERVICE_NAME" yaml:"service_name" json:"service_name" split_words:"true"`
}

type serverInfo struct {
	Address         string        `envconfig:"GAPI_SERVER_ADDRESS" yaml:"address" json:"address"`
	WriteTimeout    time.Duration `envconfig:"GAPI_SERVER_WRITE_TIMEOUT" yaml:"write_timeout" json:"write_timeout" split_words:"true"`
	ReadTimeout     time.Duration `envconfig:"GAPI_SERVER_READ_TIMEOUT" yaml:"read_timeout" json:"read_timeout" split_words:"true"`
	IdleTimeout     time.Duration `envconfig:"GAPI_SERVER_IDLE_TIMEOUT" yaml:"idle_timeout" json:"idle_timeout" split_words:"true"`
	ShutdownTimeout time.Duration `envconfig:"GAPI_SERVER_SHUTDOWN_TIMEOUT" yaml:"shutdown_timeout" json:"shutdown_timeout" split_words:"true"`
}

type corsInfo struct {
	AllowedHeaders []string `envconfig:"GAPI_CORS_ALLOWED_HEADERS" yaml:"allowed_headers" json:"allowed_headers" split_words:"true"`
	AllowedMethods []string `envconfig:"GAPI_CORS_ALLOWED_METHODS" yaml:"allowed_methods" json:"allowed_methods" split_words:"true"`
	AllowedOrigins []string `envconfig:"GAPI_CORS_ALLOWED_ORIGINS" yaml:"allowed_origins" json:"allowed_origins" split_words:"true"`
	ExposedHeaders []string `envconfig:"GAPI_CORS_EXPOSED_HEADERS" yaml:"exposed_headers" json:"exposed_headers" split_words:"true"`
	MaxAge         int      `envconfig:"GAPI_CORS_MAX_AGE" yaml:"max_age" json:"max_age" split_words:"true"`
}

type restAPIInfo struct {
	Address        string        `envconfig:"GAPI_REST_SERVER_ADDRESS" yaml:"address" json:"address"`
	RequestTimeout time.Duration `envconfig:"GAPI_REST_SERVER_REQUEST_TIMEOUT" yaml:"request_timeout" json:"request_timeout" split_words:"true"`
}

type dataBaseInfo struct {
	Database string `envconfig:"GAPI_DB_DATABASE" yaml:"database" json:"database"`
	Host     string `envconfig:"GAPI_DB_HOST" yaml:"host" json:"host"`
	Port     string `envconfig:"GAPI_DB_PORT" yaml:"port" json:"port"`
	User     string `envconfig:"GAPI_DB_USER" yaml:"user" json:"user"`
	Password string `envconfig:"GAPI_DB_PASS" yaml:"password" json:"password"`
	Dialect  string `envconfig:"GAPI_DB_DIALECT" yaml:"dialect" json:"dialect"`
}

// Address Return DB Address
func (db dataBaseInfo) HostAddress() string {
	return fmt.Sprintf("%s:%s", db.Host, db.Port)
}

// LoadServiceConfig ...
func LoadServiceConfig(configFile string) (*ServiceConfig, error) {
	var cfg ServiceConfig

	if err := loadServiceConfigFromFile(configFile, &cfg); err != nil {
		return nil, err
	}

	if err := envconfig.Process("", &cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}

func loadServiceConfigFromFile(configFile string, cfg *ServiceConfig) error {
	_, err := os.Stat(configFile)
	if err != nil {
		if os.IsNotExist(err) {
			return nil
		}
		return err
	}

	yamlFile, err := ioutil.ReadFile(configFile)
	if err != nil {
		return err
	}

	return yaml.Unmarshal(yamlFile, &cfg)
}

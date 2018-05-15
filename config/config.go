package config

import (
	"github.com/BurntSushi/toml"
)

var appConfig *AppConfig

// AppConfig will include all config items for use later
type AppConfig struct {
	GlobalConfig   globalConfig   `toml:"global"`
	DatabaseConfig databaseConfig `toml:"database"`
	GeoDataConfig  geoDataConfig  `toml:"geodata"`
	HTTPConfig     httpConfig     `toml:"http"`
}

type globalConfig struct {
	ListenAt   string `toml:"listenAt"`
	DataFolder string `toml:"dataFolder"`
	LogLevel   string `toml:"logLevel"`
	LogFile    string `toml:"logFile"`
}

type databaseConfig struct {
	HostAndPort string `toml:"hostAndPort"`
	Database    int    `toml:"database"`
}

type geoDataConfig struct {
	Geolite2CityURL    string `toml:"geolite2_city"`
	Geolite2CountryURL string `toml:"geolite2_country"`
	Update             bool   `toml:"update"`
	Crond              string `toml:"crond"`
}

// Using for download geoip database only
type httpConfig struct {
	Connect   string `toml:"connect"`
	HTTPProxy string `toml:"http_proxy"`
}

// NewAppConfig will read file and return a app config object
func NewAppConfig(path string) (*AppConfig, error) {
	var config AppConfig
	if _, err := toml.DecodeFile(path, &config); err != nil {
		return nil, err
	}
	appConfig = &config
	return &config, nil
}

// GetAppConfig return config object for application
func GetAppConfig() *AppConfig {
	if appConfig == nil {
		return nil
	}
	return appConfig
}

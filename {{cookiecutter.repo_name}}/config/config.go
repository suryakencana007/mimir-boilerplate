package config

import "time"

// Config struct generate
type Config struct {
	App struct {
		Name         string         `mapstructure:"name"`
		Version      string         `mapstructure:"version"`
		Port         int            `mapstructure:"port"`
		ReadTimeout  int            `mapstructure:"read_timeout"`
		WriteTimeout int            `mapstructure:"write_timeout"`
		Timezone     string         `mapstructure:"timezone"`
		Debug        bool           `mapstructure:"debug"`
		Env          string         `mapstructure:"env"`
		SecretKey    string         `mapstructure:"secret_key"`
		ExpireIn     *time.Duration `mapstructure:"expire_in"`
	}
	CB struct {
		Retry      int `mapstructure:"retry_count"`
		Timeout    int `mapstructure:"db_timeout"`
		Concurrent int `mapstructure:"max_concurrent"`
	}
	DB struct {
		DsnMain           string `mapstructure:"dsn_main" toml:"dsn_main,omitempty"`
		DsnSecondary      string `mapstructure:"dsn_secondary" toml:"dsn_secondary,omitempty"`
		MaxLifeTime       int    `mapstructure:"max_life_time"`
		MaxIdleConnection int    `mapstructure:"max_idle_connection"`
		MaxOpenConnection int    `mapstructure:"max_open_connection"`
	}
	Rest struct {
		Version string `mapstructure:"version"`
		Prefix  string `mapstructure:"prefix"`
	}
}

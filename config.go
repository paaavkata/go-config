package goconfig

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/spf13/viper"
)

// Config provides configuration management
type Config struct {
	viper *viper.Viper
}

// NewConfig creates a new configuration instance
func NewConfig() *Config {
	v := viper.New()
	v.AutomaticEnv()
	v.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	return &Config{viper: v}
}

// LoadFromFile loads configuration from a file
func (c *Config) LoadFromFile(path string) error {
	c.viper.SetConfigFile(path)
	return c.viper.ReadInConfig()
}

// GetString gets a string configuration value
func (c *Config) GetString(key string) string {
	return c.viper.GetString(key)
}

// GetInt gets an integer configuration value
func (c *Config) GetInt(key string) int {
	return c.viper.GetInt(key)
}

// GetBool gets a boolean configuration value
func (c *Config) GetBool(key string) bool {
	return c.viper.GetBool(key)
}

// GetStringSlice gets a string slice configuration value
func (c *Config) GetStringSlice(key string) []string {
	return c.viper.GetStringSlice(key)
}

// GetFloat64 gets a float64 configuration value
func (c *Config) GetFloat64(key string) float64 {
	return c.viper.GetFloat64(key)
}

// Set sets a configuration value
func (c *Config) Set(key string, value interface{}) {
	c.viper.Set(key, value)
}

// SetDefault sets a default configuration value
func (c *Config) SetDefault(key string, value interface{}) {
	c.viper.SetDefault(key, value)
}

// IsSet checks if a configuration value is set
func (c *Config) IsSet(key string) bool {
	return c.viper.IsSet(key)
}

// AllSettings gets all configuration values
func (c *Config) AllSettings() map[string]interface{} {
	return c.viper.AllSettings()
}

// BindEnv binds an environment variable to a configuration key
func (c *Config) BindEnv(input ...string) error {
	return c.viper.BindEnv(input...)
}

// Validate validates the configuration
func (c *Config) Validate(required []string) error {
	for _, key := range required {
		if !c.IsSet(key) {
			return fmt.Errorf("required configuration key %s is not set", key)
		}
	}
	return nil
}

// GetEnv gets an environment variable with a default value
func GetEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}

// GetEnvInt gets an integer environment variable with a default value
func GetEnvInt(key string, defaultValue int) int {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	intValue, err := strconv.Atoi(value)
	if err != nil {
		return defaultValue
	}
	return intValue
}

// GetEnvBool gets a boolean environment variable with a default value
func GetEnvBool(key string, defaultValue bool) bool {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	boolValue, err := strconv.ParseBool(value)
	if err != nil {
		return defaultValue
	}
	return boolValue
}

// GetEnvStringSlice gets a string slice environment variable with a default value
func GetEnvStringSlice(key string, defaultValue []string) []string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return strings.Split(value, ",")
}

// GetEnvFloat64 gets a float64 environment variable with a default value
func GetEnvFloat64(key string, defaultValue float64) float64 {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	floatValue, err := strconv.ParseFloat(value, 64)
	if err != nil {
		return defaultValue
	}
	return floatValue
}

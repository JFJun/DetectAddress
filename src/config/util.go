package config

import (
	"time"
	"strconv"
	"github.com/spf13/viper"
)

func getInt64(key string, defaultValue int64) int64 {
	var (
		value int64
	)
	if value = viper.GetInt64(key); value < 0 {
		return defaultValue
	}
	return value
}
func getInt(key string, defaultValue int) int {
	var (
		value int
	)
	if value = viper.GetInt(key); value < 0 {
		return defaultValue
	}
	return value
}

func getFloat64(key string, defaultValue float64) float64 {
	var (
		value float64
	)
	if value = viper.GetFloat64(key); value < 0 {
		return defaultValue
	}
	return value
}

func getString(key string, defaultValue string) string {
	var (
		value string
	)
	if value = viper.GetString(key); value == "" {
		return defaultValue
	}
	return value
}

func getStringSlice(key string, defaultValue []string) []string {
	var (
		value []string
	)
	if value = viper.GetStringSlice(key); len(value) == 0 {
		return defaultValue
	}
	return value
}

func getDuration(key string, defaultValue time.Duration) time.Duration {
	var (
		value string
	)
	if value = viper.GetString(key); value == "" {
		return defaultValue
	}
	if duration, err := time.ParseDuration(value); err == nil {
		return duration
	}
	return defaultValue
}

func getbool(key string, defaultValue bool) bool {
	var (
		value bool
	)
	if value = viper.GetBool(key); value == false {
		return defaultValue
	}
	return value
}

func Float64ToString(num float64) string{
	var(
		value string
	)
	value = strconv.FormatFloat(num, 'f', -1, 64)
	return value
}

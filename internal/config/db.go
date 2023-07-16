package config

import (
	"fmt"
	"strings"
)

type DatabaseConfig struct {
	TimeZone     string
	Host         string
	Port         int
	User         string
	Password     string
	DatabaseName string
}

func (c DatabaseConfig) ToDataSourceName() string {
	return c.toDataSourceName(true)
}

func (c DatabaseConfig) ToEmptyDataSourceName() string {
	return c.toDataSourceName(false)
}

func (c DatabaseConfig) toDataSourceName(useDBName bool) string {
	configs := []string{
		// TODO: Remove after refactoring is done
		// fmt.Sprintf("TimeZone=%s", c.TimeZone),
		// fmt.Sprintf("host=%s", c.Host),
		// fmt.Sprintf("port=%d", c.Port),
		// fmt.Sprintf("user=%s", c.User),
		// fmt.Sprintf("password=%s", c.Password),
		// fmt.Sprintf("dbname=%s", c.DatabaseName),
	}

	// TODO: Fall back to defaults if the configs are not set
	if c.TimeZone != "" {
		configs = append(configs, fmt.Sprintf("TimeZone=%s", c.TimeZone))
	}
	if c.Host != "" {
		configs = append(configs, fmt.Sprintf("host=%s", c.Host))
	}
	if c.Port != 0 {
		configs = append(configs, fmt.Sprintf("port=%d", c.Port))
	}
	if c.User != "" {
		configs = append(configs, fmt.Sprintf("user=%s", c.User))
	}
	if c.Password != "" {
		configs = append(configs, fmt.Sprintf("password=%s", c.Password))
	}
	if c.DatabaseName != "" && useDBName {
		configs = append(configs, fmt.Sprintf("dbname=%s", c.DatabaseName))
	}

	dsnBuilder := strings.Builder{}
	for _, config := range configs {
		_, err := dsnBuilder.WriteString(config + " ")
		if err != nil {
			panic(err)
		}
	}

	return dsnBuilder.String()
}

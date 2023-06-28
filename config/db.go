package config

import (
	"fmt"
	"strings"
)

type DatabaseConfig struct {
	TimeZone string
	Host     string
	Port     int
	User     string
	Password string
	Database string
}

func (c DatabaseConfig) ToDataSourceName() string {
	configs := []string{
		fmt.Sprintf("TimeZone=%s", c.TimeZone),
		fmt.Sprintf("host=%s", c.Host),
		fmt.Sprintf("port=%d", c.Port),
		fmt.Sprintf("user=%s", c.User),
		fmt.Sprintf("password=%s", c.Password),
		fmt.Sprintf("dbname=%s", c.Database),
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

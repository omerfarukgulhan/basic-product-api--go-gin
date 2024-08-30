package app

import "example.com/product-api/common/postgresql"

type ConfigurationManager struct {
	PostgreSqlConfig postgresql.Config
}

func NewConfigurationManager() *ConfigurationManager {
	postgreSqlConfig := getPostgreSqlConfig()
	return &ConfigurationManager{
		PostgreSqlConfig: postgreSqlConfig,
	}
}

func getPostgreSqlConfig() postgresql.Config {
	return postgresql.Config{
		Host:                  "localhost",
		Port:                  "5432",
		UserName:              "postgres",
		Password:              "153515",
		DbName:                "workshops",
		MaxConnections:        "10",
		MaxConnectionIdleTime: "10s",
	}
}

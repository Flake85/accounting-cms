package config

import "flag"

type Configuration struct {
	Host 	 string
	Port 	 string
	User 	 string
	Password string
	DbName   string
}

var baseConfig = Configuration{}
var isInitialized = false

func Parse() Configuration {
	if !isInitialized {
		flag.StringVar(&baseConfig.Host, "POSTGRES_HOST", "localhost", "specifies postgres host")
		flag.StringVar(&baseConfig.Port, "POSTGRES_PORT", "54321", "specifies postgres port")
		flag.StringVar(&baseConfig.User, "POSTGRES_USER", "cms", "specifies postgres user")
		flag.StringVar(&baseConfig.Password, "POSTGRES_PASSWORD", "cms", "specifies postgres password")
		flag.StringVar(&baseConfig.DbName, "POSTGRES_DB", "cms", "specifies postgres db name")
		flag.Parse()
		isInitialized = true
	}
	return baseConfig
}

package config

type Config struct {
	ProjectName  string
	GoModule     string
	Database     string
	WebFramework string
	Orm          string
	Cache        string
	Logging      string
}

var config = &Config{
	ProjectName:  "example_demo",
	GoModule:     "github.com/Whitea029/example_demo",
	Database:     "mysql",
	WebFramework: "gin",
	Orm:          "gorm",
	Cache:        "redis",
	Logging:      "zap",
}

func GetConfig() *Config {
	return config
}

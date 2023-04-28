package conf

import (
	"github.com/gookit/config/v2"
	"github.com/gookit/config/v2/yaml"
)

var ApplicationConfig *config.Config

func init() {
	ApplicationConfig = config.Default()
	ApplicationConfig.WithOptions(config.ParseEnv)
	// add driver for support yaml content
	ApplicationConfig.AddDriver(yaml.Driver)
	err := ApplicationConfig.LoadFiles("conf/application.yml")
	if err != nil {
		panic(err)
	}
}

func GetApiKey() string {
	return ApplicationConfig.String("openai.apikey")
}

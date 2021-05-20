package viperenv

import "github.com/spf13/viper"

func ViperEnvVariable(key string) string {
	viper.SetConfigFile("../.env")
	viper.ReadInConfig()
	return viper.Get(key).(string)

}

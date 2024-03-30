package config

import (
	"fmt"
	"log"
	"os"
	"strings"
	"sync"

	"github.com/spf13/viper"
)

var instance *viper.Viper
var once sync.Once

func Load() *viper.Viper {
	once.Do(func() {
		v := viper.New()
		v.AutomaticEnv()
		v.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

		conf := os.Getenv("CONF")
		files := strings.Split(conf, ",")

		for _, file := range files {
			if file == "" {
				continue
			}
			v.SetConfigFile(file)
			if err := v.MergeInConfig(); err != nil {
				log.Fatal(fmt.Sprintf("error to load configs: %s", err))
			}
		}

		instance = v
	})

	return instance
}

func ReadConfigPath(key string, i interface{}) error {
	err := instance.UnmarshalKey(key, i)
	if err != nil {
		return err
	}

	return nil
}

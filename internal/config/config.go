package config

import (
	"fmt"
	"os"
	"strings"
	"sync"

	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
)

var (
	configOnce sync.Once
	configMap  map[string]string
)

func Config(key string) string {
	configOnce.Do(func() {
		err := godotenv.Load()
		if err != nil {
			logrus.Warn(fmt.Errorf("failed to load environment variables: %v", err))
		}

		configMap = make(map[string]string)
		for _, env := range os.Environ() {
			kv := strings.SplitN(env, "=", 2)
			if len(kv) == 2 {
				configMap[kv[0]] = kv[1]
			}
		}
	})

	return configMap[key]
}

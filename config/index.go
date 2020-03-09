package configs

import (
	"fmt"
	"encoding/json"
	"os"
	"time"
)

func NewConfig() *Configuration {
	var ENV string

	if ENV = os.Getenv("ENV"); len(ENV) == 0  {
		ENV = "development"
	}

	configFileName := fmt.Sprintf("config/environments/%s.json", ENV)

	file, _ := os.Open(configFileName)
	defer file.Close()
	decoder := json.NewDecoder(file)
	conf := Configuration{
		env: ENV,
	}
	err := decoder.Decode(&conf)
	if err != nil {
		fmt.Println("Error:", err)
	}

	return &conf
}

type DatabaseConfig struct {
	Username string
	Password string
	Database string
	Host string
	Dialect string
	Logging string
}

type WebConfig struct {
	Port int "json:port"
	ReadTimeout time.Duration "json:read_timeout"
	WriteTimeout time.Duration "json:write_timeout"
}

type MessageQueueConfig struct {
	Host string
	Exchange string
	Queue string
}

type RedisConfig struct {
	Host string
	Port int
	KeyPrefix string
}

type WebServiceConfig struct {
	Host string
	Port int
	Timeout int
}

type Configuration struct {
	env string
	Debug string
	AuthSecret string
	Web WebConfig
	Db DatabaseConfig
	MessageQueue MessageQueueConfig
}
package bootstrap

import (
	"log"
	"github.com/spf13/viper"
)

type Env struct {
	AppEnv 				string  	`mapstructure:"APP_ENV"`
	ServerPort 			string  	`mapstructure:"SERVER_PORT"`

	DBHost 				string  	`mapstructure:"DB_HOST"`
	DBPort 				string  	`mapstructure:"DB_PORT"`
	DBUser 				string  	`mapstructure:"DB_USER"`
	DBPass	 			string  	`mapstructure:"DB_PASS"`
	DBName		 		string  	`mapstructure:"DB_NAME"`

	CacheHost 			string 		`mapstructure:"CACHE_HOST"`
	CachePort 			string  	`mapstructure:"CACHE_PORT"`
	CachePassword 		string 		`mapstructure:"CACHE_PASSWORD"`
	CacheDB				int 		`mapstructure:"CACHE_DB"`

	MQHost 				string 		`mapstructure:"MQ_HOST"`
	MQPort 				string 		`mapstructure:"MQ_PORT"`
	MQTopic				string 		`mapstructure:"MQ_TOPIC"`
	MQBrokers 			[]string 	`mapstructure:"MQ_BROKERS"`
	MQGroupID			string 		`mapstructure:"MQ_GROUP_ID"`
	MQGroupSize			int 		`mapstructure:"MQ_GROUP_SIZE"`
	MQPartition			int 		`mapstructure:"MQ_PARTITION"`
	MQReplicationFactor	int 		`mapstructure:"MQ_REPLICATION_FACTOR"`
	MQMinBytes			int 		`mapstructure:"MQ_MIN_BYTES"`
	MQMaxBytes			int 		`mapstructure:"MQ_MAX_BYTES"`
	MQStartOffset		int 		`mapstructure:"MQ_START_OFFSET"`
	MQMaxWait			int 		`mapstructure:"MQ_MAX_WAIT"`
	MQReadBackoffMin	int 		`mapstructure:"MQ_READ_BACKOFF_MIN"`
	MQReadBackoffMax	int 		`mapstructure:"MQ_READ_BACKOFF_MAX"`
	MQCommitInterval	int 		`mapstructure:"MQ_COMMIT_INTERVAL"`
	
	MQEnableAutoCommit	bool 		`mapstructure:"MQ_ENABLE_AUTO_COMMIT"`

	ChanSize			int 		`mapstructure:"CHAN_SIZE"`
	ChanBatchSize		int 		`mapstructure:"CHAN_BATCH_SIZE"`
	ChanDuration		int 		`mapstructure:"CHAN_DURATION"`
}

func NewEnv() *Env {
	env := Env{}

	viper.SetConfigFile(".env")
	viper.SetConfigType("env")

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal("Can't find .env file:", err)
	}

	err = viper.Unmarshal(&env)
	if err != nil {
		log.Fatal("Error unmarshalling .env file:", err)
	}

	if env.AppEnv == "development" {
		log.Println("The Application is running in development environment")
	}

	return &env
}

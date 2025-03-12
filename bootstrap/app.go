package bootstrap

import (
	interfaces "course_seckill_clean_architecture/interface"
)

type Application struct {
	Env 		*Env
	MySQL		interfaces.Database
	Redis		interfaces.Cache
	Kafka		interfaces.MsgQueue
	Channel		interfaces.Channel
	Description string
}

func App() *Application {
	app := &Application{}
	app.Env = NewEnv()
	app.MySQL = NewMySQL(app.Env)
	app.Redis = NewRedis(app.Env)
	app.Kafka = NewKafka(app.Env)
	app.Channel = NewChannel(app.Env)
	app.Description = "Course Seckill System"
	return app
}

func (app *Application) CloseConnections() {
	CloseMySQL(app.MySQL)
	CloseRedis(app.Redis)
	CloseKafka(app.Kafka)
	CloseChannel(app.Channel)
}


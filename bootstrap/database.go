package bootstrap

import (
	"context"
	"fmt"
	"log"

	"course_seckill_clean_architecture/interface"
	"course_seckill_clean_architecture/internal/repository/mysql"
)

func NewMySQL(env *Env) interfaces.Database {
	host := env.DBHost
	port := env.DBPort
	user := env.DBUser
	pass := env.DBPass
	name := env.DBName

	mysqlDNS := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		user,
		pass,
		host,
		port,
		name,
	)

	db, err := mysql.NewInstance(context.Background(), mysqlDNS)
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	return db
}

func CloseMySQL(db interfaces.Database) {
	err := db.Close()
	if err != nil {
		log.Fatal("Failed to close database:", err)
	}
}

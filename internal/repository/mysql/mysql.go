package mysql

import (
	"context"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	interfaces "course_seckill_clean_architecture/interface"
	"fmt"
)

type mysqlDatabase struct {
	db *gorm.DB
}

func NewInstance(ctx context.Context, dsn string) (interfaces.Database, error) {
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return &mysqlDatabase{db: db}, nil
}

func (m *mysqlDatabase) Create(ctx context.Context, model interface{}, value ...interface{}) error {
	// Auto migrate the structs, if already exists, ignore
	err := m.db.AutoMigrate(model)
		if err != nil {
			fmt.Println("Failed to auto migrate")
			return fmt.Errorf("failed to auto migrate: %v", err)
		}
	// If no value is provided, only migrate
	if len(value) == 0 {
		return nil
	}
	for _, v := range value {
		err := m.db.Model(model).Create(v).Error
		if err != nil {
			return err
		}
	}
	return nil
}

func (m *mysqlDatabase) Find(ctx context.Context, model interface{}, query string, args ...interface{}) error {
	// Support query with or without query. If query is nil, it will return all data
	if query != "all" {
		return m.db.Where(query, args...).First(model).Error
	}
	return m.db.Find(model).Error
}

func (m *mysqlDatabase) Update(ctx context.Context, model interface{}, updates map[string]interface{}, query string, args ...interface{}) error {
	for k, v := range updates {
		updates[k] = gorm.Expr(v.(string))
	}
	return m.db.Model(model).Where(query, args...).Updates(updates).Error
}

func (m *mysqlDatabase) Delete(ctx context.Context, model interface{}, query string, args ...interface{}) error {
	if query != "all" {	
		return m.db.Where(query, args...).Delete(model).Error
	}
	return m.db.Migrator().DropTable(model)
}

func (m *mysqlDatabase) Transaction(ctx context.Context, fn func(ctx context.Context, client interfaces.Database, args ...interface{}) error, rb func(ctx context.Context, client interfaces.Database, err error, args ...interface{}) error, args ...interface{}) error {
	// A simple implementation of transaction. 
	// Using fn as main operation and rb for extra logic when rollback
	tx := m.db.Begin() 
	if tx.Error != nil {
		return tx.Error
	}

	err := fn(ctx, m, args...)
	if err != nil {
		rb(ctx, m, err, args...)	
		tx.Rollback()
		return err
	}

	err = tx.Commit().Error
	if err != nil {
		rb(ctx, m, err, args...)
		return err
	}

	return nil
}

func (m *mysqlDatabase) NotFoundError() error {
	return gorm.ErrRecordNotFound
}

func (m *mysqlDatabase) Expr(ctx context.Context, expr string) interface{} {
	return gorm.Expr(expr)
}

func (m *mysqlDatabase) Close() error {
	sqlDB, err := m.db.DB()
	if err != nil {
		return err
	}
	return sqlDB.Close()
}

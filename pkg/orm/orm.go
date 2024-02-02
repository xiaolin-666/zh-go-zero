package orm

import (
	"context"
	"github.com/zeromicro/go-zero/core/logx"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/utils"
	"time"
)

type Config struct {
	DSN          string
	MaxOpenConns int
	MaxIdleConns int
	MaxLifeTime  int
}

type ormLog struct {
	LogLevel logger.LogLevel
}

func (l *ormLog) LogMode(level logger.LogLevel) logger.Interface {
	l.LogLevel = level
	return l
}

func (l *ormLog) Info(ctx context.Context, s string, v ...interface{}) {
	if l.LogLevel >= logger.Info {
		logx.WithContext(ctx).Infof(s, v...)
	}
}

func (l *ormLog) Warn(ctx context.Context, s string, v ...interface{}) {
	if l.LogLevel >= logger.Warn {
		logx.WithContext(ctx).Infof(s, v...)
	}
}

func (l *ormLog) Error(ctx context.Context, s string, v ...interface{}) {
	if l.LogLevel >= logger.Error {
		logx.WithContext(ctx).Errorf(s, v...)
	}
}

func (l *ormLog) Trace(ctx context.Context, begin time.Time, fc func() (sql string, rowsAffected int64), err error) {
	elapsed := time.Since(begin)
	sql, rows := fc()
	logx.WithContext(ctx).WithDuration(elapsed).Infof(utils.FileWithLineNum(), float64(elapsed.Nanoseconds())/1e6, rows, sql)
}

func NewMysqlDB(conf *Config) (*gorm.DB, error) {
	if conf.MaxIdleConns == 0 {
		conf.MaxIdleConns = 10
	}
	if conf.MaxOpenConns == 0 {
		conf.MaxOpenConns = 100
	}
	if conf.MaxLifeTime == 0 {
		conf.MaxLifeTime = 3600
	}
	db, err := gorm.Open(mysql.Open(conf.DSN), &gorm.Config{
		Logger: &ormLog{},
	})
	if err != nil {
		return nil, err
	}
	sdb, err := db.DB()
	if err != nil {
		return nil, err
	}
	sdb.SetMaxIdleConns(conf.MaxIdleConns)
	sdb.SetMaxOpenConns(conf.MaxOpenConns)
	sdb.SetConnMaxLifetime(time.Second * time.Duration(conf.MaxLifeTime))

	err = db.Use(NewTraceAndMetricPlugin())
	if err != nil {
		return nil, err
	}

	return db, nil
}

func MustMysqlDB(conf *Config) *gorm.DB {
	db, err := NewMysqlDB(conf)
	if err != nil {
		panic(err)
	}
	return db
}

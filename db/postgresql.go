package db

import (
	"database/sql"
	"log"
	"os"
	"strconv"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

// NewPsqlGormInfra initiate postgres database client connection
func NewPsqlGormInfra(url string) *gorm.DB {
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold: time.Second, // Slow SQL threshold
			LogLevel:      logger.Info, // Log level
			Colorful:      true,        // Disable color
		},
	)
	sqlConf, _ := sql.Open("pgx", url)
	// dsn := "host=localhost user=gorm password=gorm dbname=gorm port=9920 sslmode=disable TimeZone=Asia/Shanghai"
	client, err := gorm.Open(postgres.New(postgres.Config{
		Conn: sqlConf,
	}), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{ // table name prefix, table for `User` would be `t_users`
			SingularTable: true, // use singular table name, table for `User` would be `user` with this option enabled
		},
		Logger: newLogger,
	})
	if err != nil {
		panic("failed to connect database")
	}
	// Reference: https://www.alexedwards.net/blog/configuring-sqldb
	// SetMaxIdleConns sets the maximum number of connections in the idle connection pool.

	sqlDB, _ := client.DB()
	maxIdleConn, _ := strconv.Atoi(os.Getenv("DATABASE_MAX_IDLE_CONNECTION"))
	if maxIdleConn == 0 {
		maxIdleConn = 4
	}
	maxOpenConn, _ := strconv.Atoi(os.Getenv("DATABASE_MAX_OPEN_CONNECTION"))
	if maxOpenConn == 0 {
		maxOpenConn = 8
	}
	connMaxLifeTimeInMinutes, _ := strconv.Atoi(os.Getenv("DATABASE_CONNECTION_MAX_LIFETIME_IN_MINUTE"))
	if connMaxLifeTimeInMinutes == 0 {
		connMaxLifeTimeInMinutes = 5
	}
	sqlDB.SetMaxIdleConns(maxIdleConn)
	// SetMaxOpenConns sets the maximum number of open connections to the database.
	sqlDB.SetMaxOpenConns(maxOpenConn)
	// SetConnMaxLifetime sets the maximum amount of time a connection may be reused.
	sqlDB.SetConnMaxLifetime(time.Duration(connMaxLifeTimeInMinutes) * time.Minute)

	return client
}

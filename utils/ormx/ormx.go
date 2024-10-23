package ormx

import (
	"fmt"

	"github.com/BurntSushi/toml"
	"github.com/syzhang42/go-fire/auth"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type PostgresCfg struct {
	Config struct {
		Host     string `toml:"host" required:"true"`
		Port     string `toml:"port" required:"true"`
		User     string `toml:"user" required:"true"`
		DBName   string `toml:"dbname" required:"true"`
		Password string `toml:"password" required:"true"`
		Sslmode  string `toml:"sslmode" default:"disable"`
	} `toml:"postgres" required:"true"`
}
type PostgresClient struct {
	db *gorm.DB
}

var (
	postgresCfg        PostgresCfg
	defaultPostgresCli *PostgresClient
)

func Init(cfgStr string) {
	if defaultPostgresCli == nil {
		_, err := toml.DecodeFile(cfgStr, &postgresCfg)
		auth.Must(err)
		//dsn := "host=localhost user=gorm dbname=gorm password=gorm port=9920 sslmode=disable"
		dsn := fmt.Sprintf("host=%v user=%v dbname=%v password=%v port=%v sslmode=%v",
			postgresCfg.Config.Host,
			postgresCfg.Config.User,
			postgresCfg.Config.DBName,
			postgresCfg.Config.Password,
			postgresCfg.Config.Port,
			postgresCfg.Config.Sslmode,
		)
		db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
		auth.Must(err)
		if db == nil {
			panic("gorm.Open res is nil")
		}
		defaultPostgresCli = &PostgresClient{db: db}
	}
}

func GetPostgresCli() *gorm.DB {
	if defaultPostgresCli == nil {
		return nil
	}
	return defaultPostgresCli.db
}

package main

import (
	"colorist/config"
	"colorist/models"
	"colorist/routers"
	"fmt"
	"net/http"
	"time"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Mysql struct {
	Path         string `mapstructure:"path"`
	Config       string `mapstructure:"version"`
	DBName       string `mapstructure:"dbName"`
	Username     string `mapstructure:"username"`
	Password     string `mapstructure:"password"`
	MaxIdleConns int    `mapstructure:"maxIdleConns"`
	MaxOpenConns int    `mapstructure:"maxOpenConns"`
	LogMode      bool   `mapstructure:"logMode"`
}

type Config struct {
	Mysql Mysql `mapstructure:"mysql" json:"mysql" yaml:"mysql"`
}

var DB *gorm.DB
var Conf = new(Config)
var err error

func main() {
	v := viper.New()
	v.SetConfigFile("config.json")
	err := v.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("fatal error config file: %s", err))
	}

	if err := v.Unmarshal(Conf); err != nil {
		panic(fmt.Errorf("unmarshal mysql failed, err: %s", err))
	}

	// fmt.Printf("data: %v\n", Conf.Mysql.DBName)
	InitMySQL()

	routersInit := routers.InitRouter()
	s := &http.Server{
		Addr:           ":8080",
		Handler:        routersInit,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	s.ListenAndServe()

	v.WatchConfig()
	v.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("configCon file changed:", e.Name)
	})

}

func InitMySQL() {
	m := Conf.Mysql
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", m.Username, m.Password, m.Path, m.DBName)

	if config.DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{}); err != nil {
		panic(fmt.Errorf("Mysql start failed: %s", err))
	} else {
		config.DB.AutoMigrate(models.Color{})
		sqlDB, _ := config.DB.DB()
		sqlDB.SetMaxIdleConns(Conf.Mysql.MaxIdleConns)
		sqlDB.SetMaxOpenConns(Conf.Mysql.MaxOpenConns)
	}

}

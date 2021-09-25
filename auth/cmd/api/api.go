package api

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	"github.com/apascualco/apascualco-auth/internal/platform/bus/command"
	"github.com/apascualco/apascualco-auth/internal/platform/bus/event"
	"github.com/apascualco/apascualco-auth/internal/platform/bus/query"
	"github.com/apascualco/apascualco-auth/internal/platform/server"
	"github.com/apascualco/apascualco-auth/internal/platform/storage/mysql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/kelseyhightower/envconfig"
)

type config struct {
	Host string `default:"localhost"`
	Port uint   `default:"8080"`

	JWTSecret string `default:"secret"`

	DBUserReader string `default:"admin"`
	DBPassReader string `default:"admin"`
	DBHostReader string `default:"mysql_auth_read"`
	DBPortReader string `default:"3306"`
	DBNameReader string `default:"auth"`

	DBUserWriter string `default:"admin"`
	DBPassWriter string `default:"admin"`
	DBHostWriter string `default:"mysql_auth_write"`
	DBPortWriter string `default:"3306"`
	DBNameWriter string `default:"auth"`

	DBTimeout time.Duration `default:"5s"`

	RBUser     string `default:"rabbitmq"`
	RBPassword string `default:"admin"`
	RBPort     string `default:"5672"`
	RBVhost    string `default:"apascualco"`
}

func Run() error {
	cfg, err := env()
	if err != nil {
		return err
	}

	dbReader, err := databaseReader(cfg)
	if err != nil {
		return err
	}

	dbWriter, err := databaseWriter(cfg)
	if err != nil {
		return err
	}

	if err = dbReader.Ping(); err != nil {
		log.Printf("Connect reader: %s", err.Error())
	}
	if err = dbWriter.Ping(); err != nil {
		log.Printf("Connect writer: %s", err.Error())
	}

	var (
		cb = command.NewCommandBus()
		qb = query.NewQueryBus()
		eb = event.NewRabbitEventBus(cfg.RBUser, cfg.RBPassword, cfg.RBPort, cfg.RBVhost)
	)

	cr := mysql.NewMysqlUserRepository(dbReader, dbWriter, cfg.DBTimeout)

	srv := server.New(cfg.Host, cfg.Port, cb, qb, eb, cr, cfg.JWTSecret)
	srv.ConfigureSwagger()
	srv.RegisterRoutes()
	return srv.Run()
}

func env() (config, error) {
	var cfg config
	if err := envconfig.Process("AUTH", &cfg); err != nil {
		return config{}, err
	}
	return cfg, nil
}

func databaseReader(cfg config) (*sql.DB, error) {
	mysqlURI := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True", cfg.DBUserReader, cfg.DBPassReader, cfg.DBHostReader, cfg.DBPortReader, cfg.DBNameReader)
	return sql.Open("mysql", mysqlURI)
}

func databaseWriter(cfg config) (*sql.DB, error) {
	mysqlURI := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True", cfg.DBUserWriter, cfg.DBPassWriter, cfg.DBHostWriter, cfg.DBPortWriter, cfg.DBNameWriter)
	return sql.Open("mysql", mysqlURI)
}

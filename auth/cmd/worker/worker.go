package worker

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/apascualco/apascualco-auth/internal/platform/queue"
	"github.com/apascualco/apascualco-auth/internal/platform/storage/mysql"
	"github.com/kelseyhightower/envconfig"
)

type config struct {
	WaittingToStart int `default:"0"`

	DBUserWriter string `default:"admin"`
	DBPassWriter string `default:"admin"`
	DBHostWriter string `default:"mysql_auth_reader"`
	DBPortWriter string `default:"3306"`
	DBNameWriter string `default:"auth"`

	RBUser     string `default:"rabbitmq"`
	RBPassword string `default:"admin"`
	RBPort     string `default:"5672"`
	RBVhost    string `default:"apascualco"`

	DBTimeout time.Duration `default:"5s"`
}

func Start() error {
	cfg, err := env()
	if err != nil {
		return err
	}

	defer func() {
		if err := recover(); err != nil {
			fmt.Printf("[PANIC] %s panic recovered:\n%s\n",
				time.Now().Format(time.RFC3339), err)
		}
	}()

	dbWriter, err := databaseWriter(cfg)
	if err != nil {
		return err
	}
	ur := mysql.NewMysqlUserRepository(nil, dbWriter, cfg.DBTimeout)
	queue.NewQueueHandler(cfg.RBUser, cfg.RBPassword, cfg.RBPort, cfg.RBVhost, ur,
		cfg.WaittingToStart).InitializeHandlers()

	for {
		time.Sleep(time.Second)
	}
}

func env() (config, error) {
	var cfg config
	if err := envconfig.Process("AUTH", &cfg); err != nil {
		return config{}, err
	}
	return cfg, nil
}

func databaseWriter(cfg config) (*sql.DB, error) {
	mysqlURI := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True", cfg.DBUserWriter, cfg.DBPassWriter, cfg.DBHostWriter, cfg.DBPortWriter, cfg.DBNameWriter)
	return sql.Open("mysql", mysqlURI)
}

package config

import (
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
)

type Config interface {
	Get(key string) string
	SetArgs() Args
}

type configImpl struct {
}

type Args struct {
	UrlTarget string
	MysqlDsn  string
	DbTimeout time.Duration
	Port      string
}

func (config *configImpl) Get(key string) string {
	val, err := os.LookupEnv(key)
	if !err {
		log.Fatal("Env " + key + " Not Found")
	}
	return val
}

func (config *configImpl) SetArgs() Args {
	var args Args

	args.MysqlDsn = config.Get("DB_DSN")
	args.UrlTarget = config.Get("URL_TARGET")
	args.Port = config.Get("PORT")
	tmout, err := time.ParseDuration(config.Get("DB_TIMEOUT"))
	if err != nil {
		log.Fatal("failed inisiate db : " + err.Error())
	}
	args.DbTimeout = tmout

	return args
}

func New(filenames ...string) Config {
	_ = godotenv.Load(filenames...)
	return &configImpl{}
}

package config

import (
	"bytes"
	"fmt"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"io/ioutil"
	"strings"
)

var Conf ConfYaml

var defaultConf = []byte(`
	core:
	 mode: "release" # release, debug, test
	 user_work_pool_size: 1000
	 group_work_pool_size: 100
	postgres:
	 host: ""
	 port: 5432
	 db: ""
	 user: ""
	 pass: ""
	 batch_count: 5
	kafka:
	 bootstrap_servers: ""
	 group_id: "random"
	 auto_offset_reset: "earliest"
	 topic: "goshak"
	firebase:
	 url: "http://fcm.googleapis.com/fcm/send"
	 token: ""
	 timeout: 10s
	 max_connection: 10000
	 work_pool_size: 10000
	apple:
	 enable: false
	 key: 456367
	 bundle_id: "ai.nasim.bale"
	 path: "/home/amsjavan/PushCertificate.p12"
	 password: "Elenoon@91"
	prometheus:
	 port: 8080
	log:
	 level: debug
	endpoints:
	 grpc:
	   address: "127.0.0.1:5050"
	 http:
	   address: ":4040"
	   user: "test"
	   pass: "test"
`)

type ConfYaml struct {
	Postgres   SectionPostgres   `yaml:"postgres"`
	Log        SectionLog        `yaml:"log"`
}


// SectionPostgres is sub section of config.
type SectionPostgres struct {
	Host       string `yaml:"host"`
	Port       int    `yaml:"port"`
	DB         string `yaml:"db"`
	User       string `yaml:"user"`
	Pass       string `yaml:"pass"`
	BatchCount int    `yaml:"batch_count"`
}

type SectionLog struct {
	Level string `yaml:"level"`
}

// LoadConf load config from file and read in environment variables that match
func LoadConf(confPath string) (ConfYaml, error) {
	var conf ConfYaml

	viper.SetConfigType("yaml")
	viper.AutomaticEnv()         // read in environment variables that match
	viper.SetEnvPrefix("goshak") // will be uppercased automatically
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	if confPath != "" {
		content, err := ioutil.ReadFile(confPath)

		if err != nil {
			log.Errorf("File does not exist : %s", confPath)
			return conf, err
		}

		if err := viper.ReadConfig(bytes.NewBuffer(content)); err != nil {
			return conf, err
		}
	} else {
		// Search config in home directory with name ".pkg" (without extension).
		viper.AddConfigPath("/etc/goshak/")
		viper.AddConfigPath("$HOME/.goshak")
		viper.AddConfigPath(".")
		viper.SetConfigName("config")

		// If a config file is found, read it in.
		if err := viper.ReadInConfig(); err == nil {
			fmt.Println("Using config file:", viper.ConfigFileUsed())
		} else {
			// load default config
			if err := viper.ReadConfig(bytes.NewBuffer(defaultConf)); err != nil {
				return conf, err
			}
		}
	}

	// Postgres
	conf.Postgres.Host = viper.GetString("postgres.host")
	conf.Postgres.Port = viper.GetInt("postgres.port")
	conf.Postgres.DB = viper.GetString("postgres.db")
	conf.Postgres.User = viper.GetString("postgres.user")
	conf.Postgres.Pass = viper.GetString("postgres.pass")
	conf.Postgres.BatchCount = viper.GetInt("postgres.batch_count")

	//Log
	conf.Log.Level = viper.GetString("log.level")

	return conf, nil
}

func Initialize(path string) {
	var err error
	Conf, err = LoadConf(path)
	if err != nil {
		log.Fatalf("Load yaml config file error: '%v'", err)
		return
	}
}

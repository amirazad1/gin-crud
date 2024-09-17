package setting

import (
	"github.com/go-ini/ini"
	"log"
)

type Server struct {
	RunMode     string
	HTTPPort    int
	BaseURL     string
	TestBaseURL string
}

var ServerSetting = &Server{}

type Database struct {
	Type     string
	User     string
	Password string
	Host     string
	Name     string
}

var DatabaseSetting = &Database{}

var cfg *ini.File

func Setup(confPath string) {
	var err error
	cfg, err = ini.Load(confPath)
	if err != nil {
		log.Fatalf("Could not load config file %v", err.Error())
	}

	mapTo("server", ServerSetting)
	mapTo("database", DatabaseSetting)

}

func mapTo(section string, v interface{}) {
	err := cfg.Section(section).MapTo(v)
	if err != nil {
		log.Fatalf("Map Setting %s err: %v", section, err)
	}
}

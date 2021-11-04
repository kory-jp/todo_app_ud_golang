package config

import (
	"log"

	"todo_app_ud_golang/utils"

	"gopkg.in/go-ini/ini.v1"
)

type ConfigList struct {
	Port      string
	LogFile   string
	SQLDriver string
	UserName  string
	Password  string
	DBPort    string
	DBname    string
	Static    string
}

// valueは未設定でグローバルな変数Configを設定
var Config ConfigList

func init() {
	LoadConfig()
	utils.LoggingSettings(Config.LogFile)
}

func LoadConfig() {
	// ./iniファイルの読み込み
	cfg, err := ini.Load("config.ini")
	if err != nil {
		log.Fatalln(err)
	}
	Config = ConfigList{
		// MustStringはiniファイルに値が無かった場合、8080が初期値として設定される
		Port:      cfg.Section("web").Key("port").MustString("8080"),
		LogFile:   cfg.Section("web").Key("logfile").String(),
		SQLDriver: cfg.Section("db").Key("driver").String(),
		UserName:  cfg.Section("db").Key("user_name").String(),
		Password:  cfg.Section("db").Key("password").String(),
		DBPort:    cfg.Section("db").Key("port").String(),
		DBname:    cfg.Section("db").Key("db_name").String(),
		Static:    cfg.Section("web").Key("static").String(),
	}
}

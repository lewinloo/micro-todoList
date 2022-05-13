package conf

import (
	"fmt"
	"os"

	"gopkg.in/ini.v1"
)

var (
	ConsulAddr string

	MysqlIp       string
	MysqlPort     string
	MysqlUsername string
	MysqlPassword string
	MysqlDatabase string
)

func init() {
	if config, err := ini.Load("app.ini"); err != nil {
		fmt.Printf("Fail to read file: %v", err)
		os.Exit(1)
	} else {

		ConsulAddr = config.Section("consul").Key("addr").String()
		MysqlIp = config.Section("mysql").Key("ip").String()
		MysqlPort = config.Section("mysql").Key("port").String()
		MysqlUsername = config.Section("mysql").Key("user").String()
		MysqlPassword = config.Section("mysql").Key("password").String()
		MysqlDatabase = config.Section("mysql").Key("database").String()
	}
}

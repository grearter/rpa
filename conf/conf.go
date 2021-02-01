package conf

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"net"
	"os"
)

// Conf 配置
type Conf struct {
	Service struct {
		UdpPort      int    `yaml:"udpPort"`
		HttpPort     int    `yaml:"httpPort"`
		SqliteDBFile string `yaml:"sqliteDbFile"`
	} `yaml:"service"`

	Server struct {
		IP   string `yaml:"ip"`
		Port int    `yaml:"port"`
	} `yaml:"server"`
}

var (
	C = Conf{}
)

// InitConf 读取并检查配置文件
func InitConf() (err error) {
	pwd, _ := os.Getwd()
	confFile := pwd + string(os.PathSeparator) + "conf.yaml"

	data, err := ioutil.ReadFile(confFile)
	if err != nil {
		logrus.Errorf("read conf file err: %s, file: %s", confFile)
		os.Exit(1)
	}

	if err = yaml.Unmarshal(data, &C); err != nil {
		logrus.Errorf("parse conf file err: %s, file: %s", err.Error(), confFile)
		os.Exit(1)
	}

	return checkConf(&C)
}

func checkConf(c *Conf) error {
	if !(1 <= c.Service.UdpPort && c.Service.UdpPort <= 65535) {
		return fmt.Errorf("service.udpPort无效, port=%d", c.Service.UdpPort)
	}

	if !(1 <= c.Service.HttpPort && c.Service.HttpPort <= 65535) {
		return fmt.Errorf("service.httpPort无效, port=%d", c.Service.UdpPort)
	}

	if c.Service.SqliteDBFile == "" {
		return fmt.Errorf("service.sqliteDbFile为空")
	}

	if ip := net.ParseIP(c.Server.IP); ip == nil || ip.To4() == nil {
		return fmt.Errorf("server.ip无效, ip=%s", c.Server.IP)
	}

	if !(1 <= c.Server.Port && c.Server.Port <= 65535) {
		return fmt.Errorf("server.port无效, port=%d", c.Server.Port)
	}

	return nil
}

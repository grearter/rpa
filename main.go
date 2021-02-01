package main

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/grearter/rpa-agent/conf"
	"github.com/grearter/rpa-agent/controller/message"
	"github.com/grearter/rpa-agent/controller/metric"
	"github.com/grearter/rpa-agent/controller/robot"
	"github.com/grearter/rpa-agent/controller/robotmsg"
	robotmsgdao "github.com/grearter/rpa-agent/dao/robotmsg"
	"github.com/grearter/rpa-agent/util"
	"github.com/sirupsen/logrus"
	"io/ioutil"
	"net/http"
	"os"
	"os/signal"
	"time"
)

var (
	udpServer  *util.UdpServer
	httpServer *util.HttpServer
	ginEngine  *gin.Engine
)

func main() {
	if err := conf.InitConf(); err != nil {
		logrus.Errorf("init conf err: %s", err.Error())
		return
	}

	if err := checkRegister(); err != nil {
		return
	}

	if err := robotmsgdao.InitDB(conf.C.Service.SqliteDBFile); err != nil {
		logrus.Errorf("init db err: %s, dbFile: %s", err.Error(), conf.C.Service.SqliteDBFile)
		return
	}

	ginEngine = gin.Default()

	httpServer = util.NewHttpServer(conf.C.Service.HttpPort, ginEngine)
	if httpServer == nil {
		return
	}

	udpServer = util.NewUdpServer(conf.C.Service.UdpPort, robotmsg.HandlerRobotMessage)

	go func() {
		message.InitRoute(ginEngine)
		robot.InitRoute(ginEngine)
		metric.InitRoute(ginEngine)

		if err := httpServer.Serve(); err != nil {
			logrus.Errorf("httpServer serve err: %s", err.Error())
			os.Exit(2)
			return
		}
	}()

	go func() {
		if err := udpServer.Serve(); err != nil {
			logrus.Errorf("udpServer serve err: %s", err.Error())
			os.Exit(3)
		}
	}()

	signCh := make(chan os.Signal)
	signal.Notify(signCh, os.Interrupt, os.Interrupt)

	sign := <-signCh

	udpServer.Shutdown()
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*1)
	httpServer.Shutdown(ctx)
	cancel()

	logrus.Infof("receive signal:%s, EXITED", sign)
	return
}

var (
	serviceTokenFile = "service_token"
)

func checkRegister() (err error) {

	data, err := ioutil.ReadFile(serviceTokenFile)
	if err != nil && !os.IsNotExist(err) {
		logrus.Errorf("open serviceTokenFile err: %s, file: %s", err.Error(), serviceTokenFile)
		return
	}

	if len(data) > 0 { // 已经注册过, 跳过
		return
	}

	// 创建token文件
	f, err := os.Create(serviceTokenFile)
	if err != nil {
		logrus.Errorf("create serviceTokenFile err: %s, file: %s", err.Error(), serviceTokenFile)
		return err
	}

	go func() {
		for {
			ok := register2Server(f)
			if ok {
				break
			}

			logrus.Infof("register to server failed, waiting for retry...")
			time.Sleep(time.Second * 30)
		}
	}()

	return
}

// register2Server 注册自身到server
// Note: 注册成功后需要关闭f文件
func register2Server(f *os.File) (ok bool) {
	logrus.Infof("register to server %s:%d", conf.C.Server.IP, conf.C.Server.Port)

	hostname, err := os.Hostname()
	if err != nil {
		logrus.Errorf("get hostname err: %s", err.Error())
		hostname = "unknown"
	}

	url := fmt.Sprintf("http://%s:%d/agent/%s", conf.C.Server.IP, conf.C.Server.Port, hostname)

	resp, err := http.Post(url, "application/json", nil)
	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	if err != nil {
		logrus.Errorf("post err: %s, url: %s", err.Error(), url)
		return
	}

	if resp.StatusCode != http.StatusCreated {
		logrus.Errorf("resp.statusCode=%d", resp.StatusCode)
		err = fmt.Errorf("resp.statusCode=%d", resp.StatusCode)
		return
	}

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		logrus.Errorf("read resp.body err: %s, url: %s", err.Error(), url)
		return
	}

	if _, err := f.Write(data); err != nil {
		logrus.Errorf("write serviceTokenFile err: %s, file: %s, data: %s", err.Error(), serviceTokenFile, data)
		return
	}

	ok = true
	_ = f.Close()
	return
}

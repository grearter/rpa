package util

import (
	"github.com/grearter/rpa-agent/api"
	"github.com/mackerelio/go-osstat/cpu"
	"github.com/mackerelio/go-osstat/memory"
	"github.com/sirupsen/logrus"
)

func GetCpuInfo() *api.CpuInfo {

	cpuInfo, err := cpu.Get()
	if err != nil {
		logrus.Errorf("get os cpu info err: %s", err.Error())
		return nil
	}

	return &api.CpuInfo{
		Count: 2,
		Usage: float64(cpuInfo.Idle),
	}
}

func GetMemInfo() *api.MemInfo {
	memInfo, err := memory.Get()
	if err != nil {
		logrus.Errorf("get os mem info err: %s", err.Error())
		return nil
	}

	return &api.MemInfo{
		Total: int(memInfo.Total),
		Used:  int(memInfo.Used),
	}
}

package util

import (
	"github.com/grearter/rpa-agent/api"
	"github.com/mackerelio/go-osstat/memory"
	"github.com/shirou/gopsutil/cpu"
	"github.com/sirupsen/logrus"
	"time"
)

func GetCpuInfo() *api.CpuInfo {
	cpuInfos, err := cpu.Info()
	if err != nil {
		logrus.Errorf("get cpu info err: %s", err.Error())
		return nil
	}

	cpuPercent, err := cpu.Percent(time.Second*3, false)

	logrus.Errorf("cpuInfos count: %d", len(cpuInfos))

	count := 0

	for _, cpuInfo := range cpuInfos {
		count += int(cpuInfo.Cores)
	}

	return &api.CpuInfo{
		Count: count,
		Usage: cpuPercent[0],
	}
}

func GetMemInfo() *api.MemInfo {
	mem, err := memory.Get()
	if err != nil {
		logrus.Errorf("get os mem info err: %s", err.Error())
		return nil
	}

	return &api.MemInfo{
		Total: int(mem.Total),
		Used:  int(mem.Used),
	}
}

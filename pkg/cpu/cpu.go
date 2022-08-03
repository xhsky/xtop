package cpu

import (
	pCpu "github.com/shirou/gopsutil/v3/cpu"
	"github.com/xhsky/xtop/pkg/common"
)

type CPU struct {
	PhysicalCounts   int
	LogicalCounts    int
	UsedPercent      float64
	UserPercent      float64
	SystemPercent    float64
	IdlePercent      float64
	NicePercent      float64
	IowaitPercent    float64
	HardirqPercent   float64
	SoftirqPercent   float64
	StealPercent     float64
	GuestPercent     float64
	GuestNicePercent float64
}

var (
	cpuInfo                                                                                                                    CPU
	totalRate, userRate, systemRate, idleRate, niceRate, iowaitRate, irqRate, softirqRate, stealRate, guestRate, guestNiceRate common.Rate
)

func GetCpuInfo() CPU {
	physicalCounts, err := pCpu.Counts(false)
	if err != nil {
		physicalCounts = 0
	}

	logicalCounts, err := pCpu.Counts(true)
	if err != nil {
		logicalCounts = 0
	}

	userPercent, err := pCpu.Percent(0, false)
	if err != nil {
		userPercent = []float64{0}
	}

	cpuInfo.PhysicalCounts = physicalCounts
	cpuInfo.LogicalCounts = logicalCounts
	cpuInfo.UsedPercent = userPercent[0]

	pCpuTimesStat, err := pCpu.Times(false)
	if err == nil {
		totalTime := totalRate.GetRate(pCpuTimesStat[0].Total(), 1.0)
		cpuInfo.UserPercent = userRate.GetRate(pCpuTimesStat[0].User, totalTime) * 100
		cpuInfo.SystemPercent = systemRate.GetRate(pCpuTimesStat[0].System, totalTime) * 100
		cpuInfo.IdlePercent = idleRate.GetRate(pCpuTimesStat[0].Idle, totalTime) * 100
		cpuInfo.NicePercent = niceRate.GetRate(pCpuTimesStat[0].Nice, totalTime) * 100
		cpuInfo.IowaitPercent = iowaitRate.GetRate(pCpuTimesStat[0].Iowait, totalTime) * 100
		cpuInfo.HardirqPercent = irqRate.GetRate(pCpuTimesStat[0].Irq, totalTime) * 100
		cpuInfo.SoftirqPercent = softirqRate.GetRate(pCpuTimesStat[0].Softirq, totalTime) * 100
		cpuInfo.StealPercent = stealRate.GetRate(pCpuTimesStat[0].Steal, totalTime) * 100
		cpuInfo.GuestPercent = guestRate.GetRate(pCpuTimesStat[0].Guest, totalTime) * 100
		cpuInfo.GuestNicePercent = guestNiceRate.GetRate(pCpuTimesStat[0].GuestNice, totalTime) * 100
	}
	return cpuInfo
}

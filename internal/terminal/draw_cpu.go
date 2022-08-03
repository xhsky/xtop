package terminal

import (
	"fmt"

	ui "github.com/gizak/termui/v3"

	//log "github.com/sirupsen/logrus"
	"github.com/xhsky/xtop/pkg/common"
	"github.com/xhsky/xtop/pkg/cpu"
	wd "github.com/xhsky/xtop/pkg/widgets"
)

var (
	wCpu        = wd.NewWCpu()
	cpuUsedPipe common.Pipe
)

func cpuInit() {
	cpuInfo := cpu.GetCpuInfo()
	physicalCounts := cpuInfo.PhysicalCounts
	logicalCounts := cpuInfo.LogicalCounts
	wCpu.Title = "CPU"
	wCpu.RightTitle = fmt.Sprintf("%d/%d", physicalCounts, logicalCounts)
	wCpu.Border = true
	wCpu.BorderStyle = wd.CpuBorderStyle
	wCpu.TextStyle = wd.CpuTextStyle
	wCpu.MinWidth = 13

	cpuUsedPipe.NewPipe(150)
	usedPercent := cpuInfo.UsedPercent
	cpuUsedPipe.Push(usedPercent)

	wCpu.Data = make([][]float64, 1)
	wCpu.MaxVal = 100

	//wCpu.LineColors[0] = ui.ColorRed
	wCpu.LineColors = []ui.Color{ui.ColorRed}
}

func updateCpu() {
	cpuInfo := cpu.GetCpuInfo()
	usedPercent := cpuInfo.UsedPercent
	cpuUsedPipe.Push(usedPercent)
	wCpu.Data[0] = cpuUsedPipe.Show()

	wCpu.Text = []string{fmt.Sprintf("Used: %.2f", usedPercent), fmt.Sprintf("User: %.2f", cpuInfo.UserPercent),
		fmt.Sprintf("Sys: %.2f", cpuInfo.SystemPercent), fmt.Sprintf("Idle: %.2f", cpuInfo.IdlePercent),
		fmt.Sprintf("Steal: %.2f", cpuInfo.StealPercent), fmt.Sprintf("IOwait: %.2f", cpuInfo.IowaitPercent), fmt.Sprintf("Nice: %.2f", cpuInfo.NicePercent),
		fmt.Sprintf("Irq: %.2f", cpuInfo.HardirqPercent), fmt.Sprintf("Sirq: %.2f", cpuInfo.SoftirqPercent),
		fmt.Sprintf("Guest: %.2f", cpuInfo.GuestPercent), fmt.Sprintf("GNice: %.2f", cpuInfo.GuestNicePercent),
	}
	//log.Info("data: ", cpuUsedPipe.Data)
}

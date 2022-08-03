package terminal

import (
	"fmt"

	"github.com/xhsky/xtop/global"
	"github.com/xhsky/xtop/pkg/common"
	"github.com/xhsky/xtop/pkg/processes"
	wd "github.com/xhsky/xtop/pkg/widgets"
	//log "github.com/sirupsen/logrus"
)

var (
	wProcesses    = wd.NewWProcesses()
	processesInfo []processes.Process
	processesRows int
)

func setProcessSelectShow(selectedIndex int) {
	wProcesses.BottomRightTitle = fmt.Sprintf("(%v/%v)", selectedIndex, processesRows)
}

func setFilterTitle(title string) {
	wProcesses.BottomTitle = title
}

func processesInit() {
	wProcesses.ColumnsImportance = []string{"Cmdline", "Threads", "User", "Mem", "CPU%", "Pid", "Program"}
	wProcesses.ColumnsMinWidth = []int{30, 7, 6, 8, 8, 5, 10}
	wProcesses.Title = "Process"
	wProcesses.BottomTitle = "filter"
	wProcesses.RightTitle = global.SortKey
	wProcesses.ActiveRowIndex = 0
	wProcesses.IncreaseWidthColumnIndex = 2
	wProcesses.Border = true
	wProcesses.BorderStyle = wd.ProcessesBorderStyle
	wProcesses.TextStyle = wd.ProcessesTextStyle
}

func updateProcesses() {
	processesInfo = processes.GetProcesses()
	processesRows = len(processesInfo)
	setProcessSelectShow(wProcesses.ActiveRowIndex)
	wProcesses.Rows = make([][]string, processesRows+1)
	wProcesses.Rows[0] = []string{"Pid", "Program", "Cmdline", "User", "Threads", "CPU%", "Mem"}
	processesInfo = processes.ProcessSort(processesInfo, global.SortKeyMap[global.SortKey])

	if global.Filter == true {
		processesInfo = processes.ProcessFilter(processesInfo, global.FilterKey)
	}

	for i, processInfo := range processesInfo {
		wProcesses.Rows[i+1] = []string{
			fmt.Sprintf("%v", processInfo.Pid), processInfo.Command, processInfo.Cmdline,
			processInfo.User, fmt.Sprintf("%v", processInfo.NumThreads), fmt.Sprintf("%.2f", processInfo.CpuPercent), common.FormatSize(uint64(processInfo.Mem)),
		}
	}
}

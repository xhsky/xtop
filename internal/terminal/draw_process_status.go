package terminal

import (
	"fmt"
	"strconv"
	"time"

	"github.com/xhsky/xtop/pkg/common"
	"github.com/xhsky/xtop/pkg/processes"
	wd "github.com/xhsky/xtop/pkg/widgets"
	//log "github.com/sirupsen/logrus"
)

var (
	wProcess = wd.NewWProcess()
	rate     []*common.Rate
	myPid    int32
)

func processInit(pid int32) {
	var readRate, writeRate common.Rate
	myPid = pid
	rate = []*common.Rate{&readRate, &writeRate}
	processInfo := processes.GetProcess(pid)
	wProcess.ColumnsImportance = []string{"Status", "Elapsed", "R/S", "W/S"}
	wProcess.ColumnsMinWidth = []int{6, 8, 8, 8}
	wProcess.Title = processInfo.Command
	wProcess.RightTitle = strconv.Itoa(int(pid))
	wProcess.Border = true
	wProcess.BorderStyle = wd.ProcessBorderStyle

	wProcess.TextStyle = wd.ProcessTextStyle
	wProcess.Rows = make([][]string, 2)
	wProcess.Rows[0] = []string{"Status", "Elapsed", "R/S", "W/S"}

	wProcess.StrStyle = wd.ProcessStrStyle
	wProcess.Widths = []int{20, 10, 0}
	wProcess.Align = []string{"left", "right", "center"}

	wProcess.CMDStyle = wd.ProcessCMDStyle
	wProcess.CMDTextStyle = wd.ProcessCMDTextStyle
}

func updateProcess() {
	processInfo := processes.GetProcess(myPid, rate...)

	elapsed := time.Since(processInfo.Start).Round(time.Second)
	wProcess.Rows[1] = []string{
		processInfo.Status, elapsed.String(), common.FormatSize(uint64(processInfo.RBytesPerS)), common.FormatSize(uint64(processInfo.WBytesPerS)),
	}

	wProcess.Str = []string{
		fmt.Sprintf("Mem: %s(%.2f%%)", common.FormatSize(uint64(processInfo.Mem)), processInfo.MemPercent),
		fmt.Sprintf("CPU: %.2f%%", processInfo.CpuPercent),
	}
	wProcess.FormattedText.RowY = []int{3, 3}

	wProcess.CMDText = processInfo.Cmdline
	wProcess.CMD.RowY = 5
}

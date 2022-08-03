package terminal

import (
	"fmt"

	//log "github.com/sirupsen/logrus"
	"github.com/xhsky/xtop/pkg/common"
	"github.com/xhsky/xtop/pkg/mem"
	wd "github.com/xhsky/xtop/pkg/widgets"
)

var (
	wMem = wd.NewWMem()
)

func memInit() {
	wMem.Title = "Mem"
	wMem.Border = true
	wMem.BorderStyle = wd.MemBorderStyle
	wMem.TextStyle = wd.MemTextStyle

	wMem.MinWidth = 19

	wMem.LabelMinWidth = 6
	wMem.StrMinWidth = 5 // 8
	wMem.BarMinWidth = 3

	wMem.BarColors = wd.MemBarColor
	wMem.LabelStyles = wd.MemStrStyles
	wMem.StrStyles = wd.MemLabelStyles
}

func updateMem() {
	memInfo := mem.GetMemInfo()
	wMem.Text = []string{
		fmt.Sprintf("Total: %s/%s", common.FormatSize(memInfo.Total), common.FormatSize(memInfo.Used)),
		fmt.Sprintf("Swap: %s/%s", common.FormatSize(memInfo.SwapTotal), common.FormatSize(memInfo.SwapUsed)),
	}

	wMem.MaxVal = float64(memInfo.Total)

	//total := memInfo.Total
	used := memInfo.Used
	free := memInfo.Free
	buffers := memInfo.Buffers
	cached := memInfo.Cached
	shared := memInfo.Shared
	available := memInfo.Available

	wMem.Labels = []string{"Used", "Shared", "Cached", "Free", "Buff", "Avai"}
	wMem.Data = []float64{
		//float64(total),
		float64(used),
		float64(shared),
		float64(cached),
		float64(free),
		float64(buffers),
		float64(available),
	}
	wMem.Str = []string{
		//common.FormatSize(total),
		common.FormatSize(used),
		common.FormatSize(shared),
		common.FormatSize(cached),
		common.FormatSize(free),
		common.FormatSize(buffers),
		common.FormatSize(available),
	}
}

package terminal

import (
	"unicode"
	//log "github.com/sirupsen/logrus"

	"github.com/xhsky/xtop/pkg/common"
	"github.com/xhsky/xtop/pkg/net"
	wd "github.com/xhsky/xtop/pkg/widgets"
)

var (
	wNet      = wd.NewWNet()
	netsInfo  []net.Net
	ignoreNet = []string{}
)

func deleteNet(netsInfo []net.Net) []net.Net {
	ret := []net.Net{}
	for _, v := range netsInfo {
		if unicode.IsDigit(rune(v.Addr[0])) {
			ret = append(ret, v)
		}
	}
	return ret
}

func netInit() {
	netsInfo = net.GetNetAddrs()
	netsInfo = deleteNet(netsInfo)
	wNet.ColumnsImportance = []string{"Name", "Recv/S", "Send/S", "IP"}
	wNet.ColumnsMinWidth = []int{15, 8, 8, 15}
	wNet.Title = "Net"
	wNet.Border = true
	wNet.BorderStyle = wd.NetBorderStyle
	wNet.TextStyle = wd.NetTextStyle
	wNet.Rows = make([][]string, len(netsInfo)+1)
	wNet.Rows[0] = []string{"Name", "IP", "Recv/S", "Send/S"}
	netsInfo = net.GetNetsInfo(netsInfo)
}

func updateNet() {
	netsInfo = net.GetNetsInfo(netsInfo)

	for i, netInfo := range netsInfo {
		wNet.Rows[i+1] = []string{
			netInfo.Name, netInfo.Addr, common.FormatSize(uint64(netInfo.RecvBytesPerS)), common.FormatSize(uint64(netInfo.SentBytesPerS)),
		}
	}
}

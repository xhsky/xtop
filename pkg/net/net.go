package net

import (
	"strings"

	pNet "github.com/shirou/gopsutil/v3/net"
	"github.com/xhsky/xtop/global"
	"github.com/xhsky/xtop/pkg/common"
)

type Net struct {
	Name          string
	Addr          string
	RecvBytesPerS float64
	SentBytesPerS float64
}

var (
	netsInfo []Net
	netsRate [][2]common.Rate
)

func GetNetAddrs() []Net {
	netAddrs, err := pNet.Interfaces()
	netInfo := Net{}
	if err == nil {
		for _, netAddr := range netAddrs {
			addr := netAddr.Addrs
			if cap(addr) != 0 {
				netInfo.Name = netAddr.Name
				netInfo.Addr = strings.Split(netAddr.Addrs[0].Addr, "/")[0]
				netsInfo = append(netsInfo, netInfo)
			}
		}
	}
	return netsInfo
}

func GetNetsInfo(netsInfo []Net) []Net {
	if netsRate == nil {
		netsRate = make([][2]common.Rate, len(netsInfo))
	}

	netIOInfo, err := pNet.IOCounters(true)
	if err == nil {
		for i, v1 := range netsInfo {
			for _, v2 := range netIOInfo {
				if v1.Name == v2.Name {
					netsInfo[i].RecvBytesPerS = netsRate[i][0].GetRate(float64(v2.BytesRecv), float64(global.T))
					netsInfo[i].SentBytesPerS = netsRate[i][1].GetRate(float64(v2.BytesSent), float64(global.T))
				}
			}
		}
	}
	return netsInfo
}

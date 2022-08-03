package disk

import (
	"path"

	pDisk "github.com/shirou/gopsutil/v3/disk"
	"github.com/xhsky/xtop/global"
	"github.com/xhsky/xtop/pkg/common"
)

type Disk struct {
	Device           string  `label:"Device"`
	Mountpoint       string  `label:"Mount"`
	Fstype           string  `label:"Type"`
	Total            uint64  `label:"Total"`
	Free             uint64  `label:"Free"`
	Used             uint64  `label:"Used"`
	UsedPercent      float64 `label:"Used%"`
	InodeUsedPercent float64 `label:"Inode"`
	RBytePerS        float64 `label:"R/S"`
	WBytePerS        float64 `label:"W/S"`
	Util             float64 `label:"Util%"`
}

var (
	disksInfo []Disk
	//readRate, writeRate, utilRate common.Rate
	//disksRate = make(map[string]map[string]common.Rate)
	disksRate [][3]common.Rate
)

func GetDiskPartitons() []Disk {
	diskPartitions, _ := pDisk.Partitions(false)
	//var disksInfo []Disk
	var diskInfo Disk
	for _, diskPartition := range diskPartitions {
		diskInfo.Device = diskPartition.Device
		diskInfo.Mountpoint = diskPartition.Mountpoint
		diskInfo.Fstype = diskPartition.Fstype
		disksInfo = append(disksInfo, diskInfo)
	}
	return disksInfo
}

func GetDisksUsage(disksInfo []Disk) []Disk {
	for i, diskInfo := range disksInfo {
		usageInfo, _ := pDisk.Usage(diskInfo.Mountpoint)
		disksInfo[i].Total = usageInfo.Total
		disksInfo[i].Free = usageInfo.Free
		disksInfo[i].Used = usageInfo.Used
		disksInfo[i].UsedPercent = usageInfo.UsedPercent
		disksInfo[i].InodeUsedPercent = usageInfo.InodesUsedPercent
	}
	return disksInfo
}

func GetDisksIO(disksInfo []Disk) []Disk {
	if disksRate == nil {
		disksRate = make([][3]common.Rate, len(disksInfo))
	}

	for i, diskInfo := range disksInfo {
		device := diskInfo.Device
		diskIO, err := pDisk.IOCounters(device)
		if err == nil {
			deviceName := path.Base(device)
			rData := diskIO[deviceName].ReadBytes
			wData := diskIO[deviceName].WriteBytes
			uData := diskIO[deviceName].IoTime
			//log.Info("device: ", deviceName, " data:", rData, wData, uData)
			//log.Info("rate: ", deviceName, rData, wData, uData)
			disksInfo[i].RBytePerS = disksRate[i][0].GetRate(float64(rData), float64(global.T))
			disksInfo[i].WBytePerS = disksRate[i][1].GetRate(float64(wData), float64(global.T))
			disksInfo[i].Util = disksRate[i][2].GetRate(float64(uData), float64(global.T)*1000)
		}
	}
	return disksInfo
}

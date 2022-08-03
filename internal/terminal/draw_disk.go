package terminal

import (
	"fmt"

	"github.com/xhsky/xtop/pkg/common"
	"github.com/xhsky/xtop/pkg/disk"
	wd "github.com/xhsky/xtop/pkg/widgets"
	//log "github.com/sirupsen/logrus"
)

var (
	wDisk      = wd.NewWDisk()
	disksInfo  []disk.Disk
	ignoreDisk = []string{"/boot", "/boot/efi"}
)

func deleteDisk(disksInfo []disk.Disk, delNames []string) []disk.Disk {
	ret := make([]disk.Disk, 0, len(disksInfo))
	for _, v := range disksInfo {
		flag := true
		for _, delName := range delNames {
			if v.Mountpoint == delName {
				flag = false
			}
		}
		if flag == true {
			ret = append(ret, v)
		}
	}
	return ret
}

func diskInit() {
	disksInfo = disk.GetDiskPartitons()
	disksInfo = deleteDisk(disksInfo, ignoreDisk)
	wDisk.ColumnsImportance = []string{"Type", "Device", "Used", "Free", "Inode%", "Total", "Util%", "Used%", "R/S", "W/S", "Mount"}
	wDisk.ColumnsMinWidth = []int{4, 10, 8, 8, 5, 8, 6, 5, 8, 8, 10}
	wDisk.Title = "Disk"
	wDisk.Border = true
	wDisk.BorderStyle = wd.DiskBorderStyle
	wDisk.TextStyle = wd.DiskTextStyle
	wDisk.Rows = make([][]string, len(disksInfo)+1)
	wDisk.Rows[0] = []string{"Device", "Mount", "Type", "Total", "Free", "Used", "Used%", "Inode%", "R/S", "W/S", "Util%"}
	disksInfo = disk.GetDisksIO(disksInfo)
}

func updateDisk() {
	disksInfo = disk.GetDisksUsage(disksInfo)
	disksInfo = disk.GetDisksIO(disksInfo)

	for i, diskInfo := range disksInfo {
		wDisk.Rows[i+1] = []string{
			diskInfo.Device, diskInfo.Mountpoint, diskInfo.Fstype,
			common.FormatSize(diskInfo.Total), common.FormatSize(diskInfo.Free), common.FormatSize(diskInfo.Used),
			fmt.Sprintf("%.2f", diskInfo.UsedPercent), fmt.Sprintf("%.2f", diskInfo.InodeUsedPercent),
			common.FormatSize(uint64(diskInfo.RBytePerS)), common.FormatSize(uint64(diskInfo.WBytePerS)), fmt.Sprintf("%.2f", diskInfo.Util*100),
		}
	}
}

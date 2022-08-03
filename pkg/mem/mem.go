package mem

import p_mem "github.com/shirou/gopsutil/v3/mem"

type Mem struct {
	Total       uint64
	Used        uint64
	UsedPercent float64
	Free        uint64
	Buffers     uint64
	Cached      uint64
	Shared      uint64
	Available   uint64
	SwapTotal   uint64
	SwapUsed    uint64
}

var memInfo Mem

func GetMemInfo() Mem {
	p_memInfo, err := p_mem.VirtualMemory()
	if err == nil {
		memInfo.Total = p_memInfo.Total
		memInfo.Used = p_memInfo.Used
		memInfo.UsedPercent = p_memInfo.UsedPercent
		memInfo.Free = p_memInfo.Free
		memInfo.Buffers = p_memInfo.Buffers
		memInfo.Cached = p_memInfo.Cached
		memInfo.Shared = p_memInfo.Shared
		memInfo.Available = p_memInfo.Available
	}
	p_swap_info, err := p_mem.SwapMemory()
	if err == nil {
		memInfo.SwapTotal = p_swap_info.Total
		memInfo.SwapUsed = p_swap_info.Used
	}
	return memInfo
}

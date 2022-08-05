package processes

import (
	"bufio"
	"io"
	"io/ioutil"
	"os"
	"os/user"
	"path"
	"sort"
	"strconv"
	"strings"
	"syscall"
	"time"

	//log "github.com/sirupsen/logrus"
	log "github.com/sirupsen/logrus"

	"github.com/shirou/gopsutil/v3/cpu"
	"github.com/tklauser/go-sysconf"

	"github.com/xhsky/xtop/global"
	"github.com/xhsky/xtop/pkg/common"
	"github.com/xhsky/xtop/pkg/mem"
)

type Process struct {
	Pid            int32  `tag:"Pid"`
	Command        string `tag:"Command"`
	User           string `tag:"User"`
	oldProcessTime float64
	CpuPercent     float64 `tag:"CPU%"`
	MemPercent     float64 `tag:"Mem%"`
	Mem            float64 `tag:"Mem%"`
	Status         string  `tag:"Status"`
	Cmdline        string
	start          int64
	Start          time.Time
	NumThreads     int32
	Cwd            string
	Exe            string
	RBytesPerS     float64
	WBytesPerS     float64
	REBytesPerS    float64
	SEBytesPerS    float64
	Ppid           int32
	IsRunning      bool
	NoFile         int
}

var (
	processesInfo []Process
	processesMap  = make(map[int32]*Process, 500)
	//processesRate [][4]common.Rate
	//processesCPURate []common.Rate
	procPath         string = "/proc/"
	cpuTotalTimeRate common.Rate
	pageSize               = os.Getpagesize()
	clickTicks       int64 = 100 // 默认
	cpuLogicalCounts int
	memTotal         uint64
)

func init() {
	clkTck, err := sysconf.Sysconf(sysconf.SC_CLK_TCK)
	if err == nil {
		clickTicks = clkTck
	}
	counts, err := cpu.Counts(true)
	if err == nil {
		cpuLogicalCounts = counts
	}

	memTotal = mem.GetMemInfo().Total
}

func GetPids() []int32 {
	var ret []int32
	pids, _ := ioutil.ReadDir(procPath)
	for _, f := range pids {
		pid, err := strconv.ParseInt(f.Name(), 10, 32)
		if err != nil {
			continue
		}
		ret = append(ret, int32(pid))
	}
	return ret
}

func parseFloat(val string) float64 {
	floatVal, _ := strconv.ParseFloat(val, 64)
	return floatVal
}

func parseInt(val string) int32 {
	intVal, _ := strconv.ParseInt(val, 10, 32)
	return int32(intVal)
}
func parseInt64(val string) int64 {
	intVal, _ := strconv.ParseInt(val, 10, 64)
	return intVal
}

func getCPUTotalTime() float64 {
	f, _ := os.Open(path.Join(procPath, "stat"))
	defer f.Close()

	var total float64
	rd := bufio.NewReader(f)
	for {
		line, err := rd.ReadString('\n')
		if strings.HasPrefix(line, "cpu ") {
			cpuTime := strings.Split(line, " ")
			for _, v := range cpuTime {
				total += parseFloat(v)
			}
			break
		} else if err == io.EOF {
			break
		} else {
			log.Error(err)
			break
		}
	}
	return total
}

func getProcessBTime(pid int32) time.Time {
	sys := syscall.Sysinfo_t{}
	syscall.Sysinfo(&sys)
	Uptime := time.Now().Unix() - sys.Uptime
	start := processesMap[pid].start

	return time.Unix(Uptime+start, 0)
}

func getPidRW(pid int32) (float64, float64) {
	var readBytes, writeBytes float64
	ioFile := path.Join(procPath, strconv.Itoa(int(pid)), "io")
	f, err := os.Open(ioFile)
	defer f.Close()
	if err == nil {
		buf := bufio.NewReader(f)
		for {
			line, err := buf.ReadString('\n')
			if err == nil {
				if strings.HasPrefix(line, "read_bytes:") {
					readBytes = parseFloat(strings.TrimSpace(strings.Split(line, ":")[1]))
				}
				if strings.HasPrefix(line, "write_bytes:") {
					writeBytes = parseFloat(strings.TrimSpace(strings.Split(line, ":")[1]))
				}
			} else if err == io.EOF {
				break
			} else {
				log.Error(err)
				break
			}
		}
	} else {
		log.Error(err)
	}
	return readBytes, writeBytes
}

func GetPidIO(pid int32, readRate, writeRate *common.Rate) (float64, float64) {
	readBytes, writeBytes := getPidRW(pid)
	return readRate.GetRate(readBytes, float64(global.T)), writeRate.GetRate(writeBytes, float64(global.T))
}

func GetPidNetIO(pid int32, receRate, sendRate common.Rate) (float64, float64) {
	receBytes, sendBytes := getPidRW(pid)
	return receRate.GetRate(receBytes, float64(global.T)), sendRate.GetRate(sendBytes, float64(global.T))
}

func GetPidStat(pid int32) (string, string, int32, float64, int32, float64, int64) {
	statFile := path.Join(procPath, strconv.Itoa(int(pid)), "stat")
	pidStat, err := ioutil.ReadFile(statFile)
	if err == nil {
		stat := strings.Split(string(pidStat), ") ")
		comm := strings.Split(stat[0], " (")[1]
		pidStatSlice := strings.Split(stat[1], " ")

		state := pidStatSlice[0]
		ppid := parseInt(pidStatSlice[1])
		uTime := parseFloat(pidStatSlice[11])
		sTime := parseFloat(pidStatSlice[12])
		cuTime := parseFloat(pidStatSlice[13])
		csTime := parseFloat(pidStatSlice[14])
		threads := parseInt(pidStatSlice[17])
		startTime := parseInt64(pidStatSlice[19])
		rss := parseFloat(pidStatSlice[21]) * float64(pageSize)
		return comm, state, ppid, uTime + sTime + cuTime + csTime, threads, rss, startTime
	}
	return "", "", 0, 0, 0, 0, 0
}

func GetPidStatus(pid int32) string {
	statusFile := path.Join(procPath, strconv.Itoa(int(pid)), "status")
	f, err := os.Open(statusFile)
	defer f.Close()
	if err == nil {
		buf := bufio.NewReader(f)
		for {
			line, err := buf.ReadString('\n')
			if err == nil {
				if strings.HasPrefix(line, "Uid:\t") {
					uid := strings.Split(line, "\t")[1]
					return uid
				}
			} else if err == io.EOF {
				return ""
			} else {
				log.Error(err)
				return ""
			}
		}
	}
	return ""
}

func GetPidCmdline(pid int32) string {
	cmdlineFile := path.Join(procPath, strconv.Itoa(int(pid)), "cmdline")
	cmdline, err := ioutil.ReadFile(cmdlineFile)
	if err == nil {
		return string(cmdline)
	}
	return ""
}

func GetPidUsername(pid int32) string {
	uid := GetPidStatus(pid)
	if uid == "" {
		return ""
	} else {
		u, err := user.LookupId(uid)
		if err == nil {
			return u.Username
		} else {
			return ""
		}
	}
}

func getPidNoFile(pid int32) int {
	fdPath := path.Join(procPath, strconv.Itoa(int(pid)), "fd")
	var nofile int
	fds, err := ioutil.ReadDir(fdPath)
	if err == nil {
		nofile = len(fds)
	}
	return nofile
}

func GetProcess(pid int32, rate ...*common.Rate) *Process {
	// rate: readRate, writeRate, receRate, sendRate

	processesMap[pid].MemPercent = processesMap[pid].Mem / float64(memTotal) * 100

	processesMap[pid].Start = getProcessBTime(pid)

	processesMap[pid].NoFile = getPidNoFile(pid)

	num := len(rate)
	// disk IO
	if num >= 2 {
		processesMap[pid].RBytesPerS, processesMap[pid].WBytesPerS = GetPidIO(pid, rate[0], rate[1])
	}
	// net IO
	if num >= 4 {
		processesMap[pid].REBytesPerS, processesMap[pid].SEBytesPerS = GetPidIO(pid, rate[2], rate[3])
	}
	return processesMap[pid]
}

func GetProcesses() []Process {
	pids := GetPids()
	processesInfo = make([]Process, 0, len(pids))
	cpuTimeTotal := getCPUTotalTime()
	cpuTime := cpuTotalTimeRate.GetRate(cpuTimeTotal, 1.0)
	var newProcessTime float64
	var startTime int64
	for _, pid := range pids {
		if processesMap[pid] == nil {
			processesMap[pid] = &Process{}
		}
		processesMap[pid].Pid = pid
		processesMap[pid].Command, processesMap[pid].Status, processesMap[pid].Ppid, newProcessTime, processesMap[pid].NumThreads, processesMap[pid].Mem, startTime = GetPidStat(pid)
		processesMap[pid].start = startTime / clickTicks
		processesMap[pid].CpuPercent = (newProcessTime - processesMap[pid].oldProcessTime) / cpuTime * 100 * float64(cpuLogicalCounts)
		processesMap[pid].User = GetPidUsername(pid)
		processesMap[pid].Cmdline = GetPidCmdline(pid)
		processesMap[pid].oldProcessTime = newProcessTime
	}

	for _, pid := range pids {
		processesInfo = append(processesInfo, *processesMap[pid])
	}
	return processesInfo
}

func ProcessSort(processesInfo []Process, sortedKey string) []Process {
	if sortedKey == "cpu" {
		sort.Slice(processesInfo, func(i, j int) bool {
			return processesInfo[i].CpuPercent > processesInfo[j].CpuPercent
		})
	} else if sortedKey == "mem" {
		sort.Slice(processesInfo, func(i, j int) bool {
			return float64(processesInfo[i].Mem) > float64(processesInfo[j].Mem)
		})
	}
	return processesInfo
}
func ProcessFilter(processesInfo []Process, filterKey string) []Process {
	filterProcessesInfo := []Process{}
	for _, v := range processesInfo {
		if strings.Contains(v.Command, filterKey) == true {
			filterProcessesInfo = append(filterProcessesInfo, v)
		}
	}
	return filterProcessesInfo
}

package terminal

import (
	"fmt"
	"strconv"
	"strings"
	"time"
	"unicode"

	ui "github.com/gizak/termui/v3"
	log "github.com/sirupsen/logrus"
	"github.com/xhsky/xtop/global"
)

var (
	termWidth, termHeight int
)

var (
	initFunc   = []func(){memInit, cpuInit, diskInit, netInit, processesInit}
	updateFunc = []func(){updateMem, updateCpu, updateDisk, updateNet, updateProcesses}
)

func funInit() {
	for _, fun := range initFunc {
		fun()
	}
}

func currentTimeString() string {
	t := time.Now()
	return fmt.Sprintf(t.Format("15:04:05"))
}

func refresh() {
	wProcesses.Title = currentTimeString()
	for _, fun := range updateFunc {
		fun()
	}
}

func showProcess(pid string, grid *ui.Grid, previousPid string) string {
	grid.Items = nil
	updateFunc = updateFunc[:5]
	if pid == previousPid || pid == "Pid" || pid == "" {
		grid.Set(
			ui.NewCol(1.0/2,
				ui.NewRow(2.5/7, wCpu),
				ui.NewRow(2.0/7, wDisk),
				ui.NewRow(2.5/7, wMem)),
			ui.NewCol(1.0/2,
				ui.NewRow(4.5/7, wProcesses),
				ui.NewRow(2.5/7, wNet)),
		)
		previousPid = ""
	} else {
		pidInt, _ := strconv.ParseInt(pid, 10, 32)
		processInit(int32(pidInt))
		updateProcess()
		updateFunc = append(updateFunc, updateProcess)

		grid2 := ui.NewGrid()
		grid2.Set(
			//ui.NewCol(1.0/3, wPCpu),
			ui.NewCol(1.0/1, wProcess),
		)

		grid.Set(
			ui.NewCol(1.0/2,
				ui.NewRow(2.5/7, wCpu),
				ui.NewRow(2.0/7, wDisk),
				ui.NewRow(2.5/7, wMem)),
			ui.NewCol(1.0/2,
				ui.NewRow(2.5/7, wProcesses),
				ui.NewRow(2.0/7, grid2),
				ui.NewRow(2.5/7, wNet)),
		)
		previousPid = pid
	}
	return previousPid
}

func Show() {
	if err := ui.Init(); err != nil {
		log.Fatalf("failed to initialize termui: %v", err)
	}
	defer ui.Close()

	funInit()

	termWidth, termHeight = ui.TerminalDimensions()
	log.Info("init:", termWidth, termHeight)

	grid := ui.NewGrid()
	termWidth, termHeight = ui.TerminalDimensions()
	grid.SetRect(0, 0, termWidth, termHeight)
	grid.Set(
		ui.NewCol(1.0/2,
			ui.NewRow(2.5/7, wCpu),
			ui.NewRow(2.0/7, wDisk),
			ui.NewRow(2.5/7, wMem)),
		ui.NewCol(1.0/2,
			ui.NewRow(4.5/7, wProcesses),
			ui.NewRow(2.5/7, wNet)),
	)

	refresh()
	ui.Render(grid)

	uiEvents := ui.PollEvents()
	previousKey := ""
	previousPid := ""
	filterTitle := ""
	filterStr := make([]string, 0, 10)
	filterComplete := false
	ticker := time.NewTicker(time.Second * time.Duration(global.T)).C
	for {
		select {
		case e := <-uiEvents:
			if previousKey == "f" {
				global.Filter = true
				byteEvent := []rune(e.ID)[0]
				if unicode.IsLetter(byteEvent) == true || unicode.IsNumber(byteEvent) == true || byteEvent == '_' {
					filterStr = append(filterStr, e.ID)
					e.ID = "f"
				} else if e.ID == "<Backspace>" {
					if len(filterStr) > 0 {
						filterStr = filterStr[:len(filterStr)-1]
						e.ID = "f"
					}
				} else if e.ID == "<Enter>" {
					filterLen := len(filterStr)
					if filterLen == 0 {
						setFilterTitle("filter")
						global.Filter = false
						e.ID = ""
					} else {
						filterComplete = true
						e.ID = "f"
					}
				}
			}
			switch e.ID {
			case "q", "<C-c>":
				return
			case "j", "<Down>":
				wProcesses.ScrolDown()
				setProcessSelectShow(wProcesses.ActiveRowIndex)
			case "k", "<Up>":
				wProcesses.ScrolUp()
				setProcessSelectShow(wProcesses.ActiveRowIndex)
			case "g":
				if previousKey == "g" {
					wProcesses.ScrollTop()
					previousKey = ""
				}
			case "f":
				global.FilterKey = strings.Join(filterStr, "")
				if filterComplete == true {
					filterTitle = fmt.Sprintf("f: %s Del", global.FilterKey)
					e.ID = ""
				} else {
					filterTitle = fmt.Sprintf("f: %s ↲", global.FilterKey)
				}
				setFilterTitle(filterTitle)
			case "<Delete>":
				setFilterTitle("filter")
				global.Filter = false
				filterStr = nil
				filterComplete = false
				e.ID = ""
			case "<Home>":
				wProcesses.ScrollTop()
				setProcessSelectShow(wProcesses.ActiveRowIndex)
			case "G", "<End>":
				wProcesses.ScrollBottom()
				setProcessSelectShow(wProcesses.ActiveRowIndex)
			case "P":
				global.SortKey = "< CPU >"
				wProcesses.RightTitle = global.SortKey
			case "M":
				global.SortKey = "< MEM >"
				wProcesses.RightTitle = global.SortKey
			case "<Enter>":
				global.Resize = true
				row := wProcesses.GetRow()
				if len(row) != 0 {
					previousPid = showProcess(row[0], grid, previousPid)
					ui.Clear()
					ui.Render(grid)
					global.Resize = false
				}
			case "<Resize>":
				global.Resize = true
				payload := e.Payload.(ui.Resize)
				grid.SetRect(0, 0, payload.Width, payload.Height)
				ui.Clear()
				ui.Render(grid)
				global.Resize = false
			}

			previousKey = e.ID
			/*
				if previousKey == "g" {
					previousKey = ""
				} else {
					previousKey = e.ID
				}
			*/

			ui.Render(grid)
		case <-ticker:
			refresh()
			ui.Render(grid)
			global.Resize = false
		}
	}
}

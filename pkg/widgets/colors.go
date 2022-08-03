package widgets

import ui "github.com/gizak/termui/v3"

// 对termui颜色的额外补充
const (
	ColorGreen1     ui.Color = 46
	ColorGreen2     ui.Color = 42
	ColorGreen3     ui.Color = 34
	DarkRed         ui.Color = 52
	ColorRed1       ui.Color = 195
	ColorRed3       ui.Color = 99
	RoyalBlue1      ui.Color = 63
	NavyBlue        ui.Color = 17
	MediumPurple3   ui.Color = 141
	DeepPink4       ui.Color = 125
	LightSteelBlue  ui.Color = 147
	Gold3           ui.Color = 178
	Aquamarine3     ui.Color = 79
	Orange4         ui.Color = 94
	LightGreen      ui.Color = 119
	SteelBlue1      ui.Color = 81
	CornflowerBlue  ui.Color = 69
	LightGoldenrod3 ui.Color = 179
	IndianRed       ui.Color = 173
)

var MyBorderStyle = ui.NewStyle(ui.ColorCyan)

var TextStyle = ui.NewStyle(ui.ColorYellow)

//var DiskStyle = ui.NewStyle(ui.ColorGreen)

var TitleStyle = ui.Theme.Block.Title

// CPU
var CpuBorderStyle = ui.NewStyle(ui.ColorCyan)
var CpuTextStyle = ui.NewStyle(Gold3)

// Mem
var MemBorderStyle = ui.NewStyle(ui.ColorCyan)
var MemTextStyle = ui.NewStyle(MediumPurple3)
var MemStrStyles = []ui.Style{
	ui.NewStyle(LightGoldenrod3),
	ui.NewStyle(LightGoldenrod3),
	ui.NewStyle(LightGoldenrod3),
	ui.NewStyle(LightGoldenrod3),
	ui.NewStyle(LightGoldenrod3),
	ui.NewStyle(LightGoldenrod3),
}
var MemLabelStyles = MemStrStyles
var MemBarColor = []ui.Color{
	//ui.ColorCyan,
	ui.ColorRed,
	ColorRed1,
	ColorRed3,
	ui.ColorGreen,
	ColorGreen1,
	ColorGreen2,
}

// Disk
var DiskBorderStyle = ui.NewStyle(ui.ColorCyan)
var DiskTextStyle = ui.NewStyle(Aquamarine3)

// Net
var NetBorderStyle = ui.NewStyle(ui.ColorCyan)
var NetTextStyle = ui.NewStyle(IndianRed)

// Processes
var ProcessesBorderStyle = ui.NewStyle(ui.ColorCyan)
var ProcessesTextStyle = ui.NewStyle(ui.ColorGreen)

var ShadowStyle = ui.NewStyle(ui.ColorYellow, ui.ColorBlack, ui.ModifierBold)

// Process
var ProcessBorderStyle = ui.NewStyle(SteelBlue1)
var ProcessTextStyle = ui.NewStyle(LightGreen)
var ProcessStrStyle = []ui.Style{ui.NewStyle(LightGreen), ui.NewStyle(LightGreen), ui.NewStyle(LightGreen)}

var ProcessCMDStyle = ui.NewStyle(LightGreen, ui.ColorClear, ui.ModifierBold)
var ProcessCMDTextStyle = ui.NewStyle(LightGreen)

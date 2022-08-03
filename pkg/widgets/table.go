package widgets

import (
	"image"

	. "github.com/gizak/termui/v3"
	"github.com/xhsky/xtop/global"
	//log "github.com/sirupsen/logrus"
)

type Table struct {
	Rows                     [][]string
	TitleStyle               Style
	ColumnsImportance        []string
	ColumnsMinWidth          []int
	IncreaseWidthColumnIndex int
	TextStyle                Style
	ColumnsStyles            map[int]Style
	FillRow                  bool

	adapTitleIndex  []int
	adapColumnWidth []int

	ActiveRowIndex int
	ActiveRowStyle Style
}

func NewTable() *Table {
	return &Table{
		TextStyle:       Theme.Table.Text,
		ColumnsStyles:   make(map[int]Style),
		adapTitleIndex:  []int{},
		adapColumnWidth: []int{},
		ActiveRowIndex:  0,
		FillRow:         true,
	}
}

func judeWidth(terminalWidth int, columnsSum int, ColumnsImportance []string, ColumnsMinWidth []int, delColumns []string, i int) []string {
	// 根据termianlWidth返回要删除的字段名
	if terminalWidth >= columnsSum {
		return delColumns
	} else {
		columnsSum -= ColumnsMinWidth[i]
		terminalWidth += 1 // 加上删去字段后的空格
		delColumns = append(delColumns, ColumnsImportance[i])
		i++
		return judeWidth(terminalWidth, columnsSum, ColumnsImportance, ColumnsMinWidth, delColumns, i)
	}
}

func getSomeData(title []string, widgetWidth int, ColumnsMinWidth []int) (columnsNum, avaiWidth, columnsSum int) {
	columnsNum = len(title)
	avaiWidth = widgetWidth - (columnsNum - 1) // 字段以空格为间隔

	columnsSum = 0
	for _, v := range ColumnsMinWidth {
		columnsSum += v
	}
	return
}

func (self *Table) jude(dx int) ([]int, []int) {
	// 删减title
	_, avaiWidth, columnsSum := getSomeData(self.Rows[0], dx, self.ColumnsMinWidth)
	i := 0 // 次数
	delTitle := judeWidth(avaiWidth, columnsSum, self.ColumnsImportance, self.ColumnsMinWidth, []string{}, i)
	//log.Info("delCol: ", self.Title, delTitle, avaiWidth, columnsSum)

	// adapTitle
	var adapTitle []string
	var adapTitleIndex []int
	for index, i := range self.Rows[0] {
		flag := true
		for _, j := range delTitle {
			if i == j {
				flag = false
				break
			}
		}
		if flag == true {
			adapTitle = append(adapTitle, i)
			adapTitleIndex = append(adapTitleIndex, index)
		}
	}
	//log.Info("adap: ", adapTitle, adapTitleIndex)

	// 调整title中列的最小长度
	var adapColumnsMinWidth []int
	var columnsMinWidthMap = make(map[string]int)
	for i, v := range self.ColumnsImportance {
		columnsMinWidthMap[v] = self.ColumnsMinWidth[i]
	}
	for _, v := range adapTitle {
		adapColumnsMinWidth = append(adapColumnsMinWidth, columnsMinWidthMap[v])
	}

	var columnsNum int
	columnsNum, avaiWidth, columnsSum = getSomeData(adapTitle, dx, adapColumnsMinWidth)
	rowMults := avaiWidth / columnsSum
	rowExtra := avaiWidth % columnsSum
	colMults := rowExtra / columnsNum
	colExtra := rowExtra % columnsNum

	adapColumnWidth := make([]int, len(adapColumnsMinWidth))
	for i, v := range adapColumnsMinWidth {
		if i == self.IncreaseWidthColumnIndex {
			adapColumnWidth[i] = v*rowMults + colMults + colExtra
		} else {
			adapColumnWidth[i] = v*rowMults + colMults
		}
	}
	return adapTitleIndex, adapColumnWidth
}

func (self *Table) Render(buf *Buffer, block *ExtraBlock) {
	if global.Resize == true {
		self.adapTitleIndex, self.adapColumnWidth = self.jude(block.Inner.Dx())
	}

	startI := 0 // scroll
	if self.ActiveRowIndex+1 > block.Inner.Dy() {
		startI = self.ActiveRowIndex + 2 - block.Inner.Dy()
	}

	yCoordinate := block.Inner.Min.Y
	for i := 0; i < len(self.Rows) && yCoordinate < block.Inner.Max.Y; i++ {
		if i != 0 && i < startI { // 保持title位置不变
			continue
		}
		row := self.Rows[i]
		//rowStyle = self.RowStyles[i]

		rowStyle := self.TextStyle
		if self.ActiveRowIndex > 0 && self.ActiveRowIndex == i {
			rowStyle = ShadowStyle
		}
		if self.FillRow {
			blackCell := NewCell(' ', rowStyle)
			buf.Fill(blackCell, image.Rect(block.Inner.Min.X, yCoordinate, block.Inner.Max.X, yCoordinate+1))
		}

		// draw row
		xCoordinate := block.Inner.Min.X
		for j := 0; j < len(row); j++ {
			for m, titleIndex := range self.adapTitleIndex {
				if j == titleIndex {
					col := ParseStyles(row[j], rowStyle)
					xLimit := xCoordinate + self.adapColumnWidth[m]
					for k, cell := range col {
						if xCoordinate+k < xLimit {
							buf.SetCell(cell, image.Pt(xCoordinate+k, yCoordinate))
						}
					}
					xCoordinate += (self.adapColumnWidth[m] + 1)
					break
				}
			}
		}
		yCoordinate += 1
	}
}

func (self *Table) ScrolDown() {
	if self.ActiveRowIndex < len(self.Rows)-1 {
		self.ActiveRowIndex++
	}
}

func (self *Table) ScrolUp() {
	if self.ActiveRowIndex > 0 {
		self.ActiveRowIndex--
	}
}

func (self *Table) ScrollBottom() {
	self.ActiveRowIndex = len(self.Rows) - 1
}

func (self *Table) ScrollTop() {
	self.ActiveRowIndex = 1
}

func (self *Table) GetRow() []string {
	return self.Rows[self.ActiveRowIndex]
}

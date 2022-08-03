package widgets

import (
	"image"

	//log "github.com/sirupsen/logrus"
	. "github.com/gizak/termui/v3"
)

type FormattedText struct {
	Str      []string
	StrStyle []Style
	RowY     []int
	Align    []string
	Widths   []int
}

func getCenterX(dx, width, xCoord int) int {
	var x int
	if dx >= width {
		x = (dx - width) / 2
	}
	return xCoord + x
}

func (self *FormattedText) render(buf *Buffer, block *ExtraBlock) {
	var strWidth int
	for i, str := range self.Str {
		var xCoord, dx int
		xCoord = block.Inner.Min.X
		switch self.Align[i] {
		case "left":
			dx = block.Inner.Dx() / 2
			xCoord = block.Inner.Min.X
		case "right":
			dx = block.Inner.Dx() / 2
			if block.Inner.Dx()%2 == 0 {
				xCoord = xCoord + dx
			} else {
				xCoord = xCoord + dx + 1
			}
		default:
			dx = block.Inner.Dx()
			xCoord = block.Inner.Min.X
		}
		if self.Widths[i] == 0 {
			strWidth = len(str)
		} else {
			strWidth = self.Widths[i]
		}

		for j, char := range str {
			cell := NewCell(char, self.StrStyle[i])
			x := getCenterX(dx, strWidth, xCoord) + j
			y := block.Inner.Min.Y + self.RowY[i]
			if x < block.Inner.Max.X && y < block.Inner.Max.Y {
				buf.SetCell(cell, image.Pt(x, y))
			}
		}
	}
}

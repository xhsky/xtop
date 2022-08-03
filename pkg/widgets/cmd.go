package widgets

import (
	"image"

	//log "github.com/sirupsen/logrus"
	. "github.com/gizak/termui/v3"
)

type CMD struct {
	CMDText      string
	CMDTextStyle Style
	CMDStyle     Style
	RowY         int
}

func (self *CMD) render(buf *Buffer, block *ExtraBlock) {
	rowStr := make([]string, 3)

	for i, char := range "CMD" {
		cell := NewCell(char, self.CMDStyle)
		x := block.Inner.Min.X
		y := block.Inner.Min.Y + self.RowY + i
		if x < block.Inner.Max.X && y < block.Inner.Max.Y {
			buf.SetCell(cell, image.Pt(x, y))
		}
	}

	cmdLen := len(self.CMDText)
	width := block.Inner.Dx() - 2
	rows := cmdLen / width

	if rows == 0 {
		rowStr[1] = self.CMDText
	} else if rows == 1 {
		rowStr[0] = self.CMDText[0:width]
		rowStr[1] = self.CMDText[width:cmdLen]
	} else if rows == 2 {
		rowStr[0] = self.CMDText[0:width]
		rowStr[1] = self.CMDText[width : width*2]
		rowStr[2] = self.CMDText[width*2 : cmdLen]
	} else {
		rowStr[0] = self.CMDText[0:width]
		rowStr[1] = self.CMDText[width : width*2]
		rowStr[2] = self.CMDText[width*2 : width*3]
	}

	for i, str := range rowStr {
		x := getCenterX(width, len(str), block.Inner.Min.X+2)
		y := block.Inner.Min.Y + self.RowY + i
		for j, char := range str {
			cell := NewCell(char, self.CMDTextStyle)
			if x+j < block.Inner.Max.X && y < block.Inner.Max.Y {
				buf.SetCell(cell, image.Pt(x+j, y))
			}
		}
	}
}

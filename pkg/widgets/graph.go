package widgets

import (
	"image"

	. "github.com/gizak/termui/v3"
	//log "github.com/sirupsen/logrus"
)

type Graph struct {
	Text      []string
	TextStyle Style
	MinWidth  int
}

func (self *Graph) render(buf *Buffer, block *ExtraBlock) {
	width := block.Inner.Dx() / len(self.Text)
	var tempWidth int
	if width >= self.MinWidth {
		tempWidth = width
	} else {
		tempWidth = self.MinWidth
	}
	rowNums := block.Inner.Dx() / tempWidth
	if rowNums == 0 {
		return
	}
	nums := len(self.Text)
	rows := nums / rowNums
	extra := nums - rows*rowNums
	for row := 0; row <= rows; row++ {
		for num := 0; num < rowNums; num++ {
			if row == rows && num > extra-1 {
				break
			}
			cells := ParseStyles(self.Text[(row*rowNums)+num], self.TextStyle)
			for x, cell := range cells {
				buf.SetCell(cell, image.Pt((num*tempWidth)+x, row).Add(block.Inner.Min))
			}
		}
	}
}

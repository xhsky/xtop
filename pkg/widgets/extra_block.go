package widgets

import (
	"image"

	. "github.com/gizak/termui/v3"
)

type ExtraBlock struct {
	Block

	RightTitle      string
	RightTitleStyle Style

	BottomTitle      string
	BottomTitleStyle Style

	BottomRightTitle      string
	BottomRightTitleStyle Style
}

func NewExtraBlock() *ExtraBlock {
	return &ExtraBlock{
		Block:                 *NewBlock(),
		RightTitleStyle:       TitleStyle,
		BottomTitleStyle:      TitleStyle,
		BottomRightTitleStyle: TitleStyle,
	}
}

func (self *ExtraBlock) Draw(buf *Buffer) {
	self.Block.Draw(buf)

	buf.SetString(
		self.BottomTitle,
		self.BottomTitleStyle,
		image.Pt(self.Min.X+2, self.Max.Y-1),
	)

	buf.SetString(
		self.RightTitle,
		self.RightTitleStyle,
		image.Pt(self.Max.X-len(self.RightTitle)-2, self.Min.Y),
	)

	buf.SetString(
		self.BottomRightTitle,
		self.BottomRightTitleStyle,
		image.Pt(self.Max.X-len(self.BottomRightTitle)-2, self.Max.Y-1),
	)
}

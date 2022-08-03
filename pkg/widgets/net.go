package widgets

import (
	. "github.com/gizak/termui/v3"
	//log "github.com/sirupsen/logrus"
)

type WNet struct {
	ExtraBlock
	Table
}

func NewWNet() *WNet {
	return &WNet{
		ExtraBlock: *NewExtraBlock(),
		Table:      *NewTable(),
	}
}

func (self *WNet) Draw(buf *Buffer) {
	self.ExtraBlock.Draw(buf)

	self.Table.Render(buf, &self.ExtraBlock)
}

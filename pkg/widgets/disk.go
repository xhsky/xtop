package widgets

import (
	. "github.com/gizak/termui/v3"
	//log "github.com/sirupsen/logrus"
)

type WDisk struct {
	ExtraBlock
	Table
}

func NewWDisk() *WDisk {
	return &WDisk{
		ExtraBlock: *NewExtraBlock(),
		Table:      *NewTable(),
	}
}

func (self *WDisk) Draw(buf *Buffer) {
	self.ExtraBlock.Draw(buf)
	self.Table.Render(buf, &self.ExtraBlock)
}

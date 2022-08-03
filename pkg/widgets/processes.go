package widgets

import (
	. "github.com/gizak/termui/v3"
	//log "github.com/sirupsen/logrus"
)

type WProcesses struct {
	ExtraBlock
	Table
}

func NewWProcesses() *WProcesses {
	return &WProcesses{
		ExtraBlock: *NewExtraBlock(),
		Table:      *NewTable(),
	}
}

func (self *WProcesses) Draw(buf *Buffer) {
	self.ExtraBlock.Draw(buf)
	self.Table.Render(buf, &self.ExtraBlock)
}

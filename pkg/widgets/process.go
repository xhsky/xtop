package widgets

import (
	. "github.com/gizak/termui/v3"
	//log "github.com/sirupsen/logrus"
)

type WProcess struct {
	ExtraBlock
	Table
	FormattedText
	CMD
}

func NewWProcess() *WProcess {
	return &WProcess{
		ExtraBlock: *NewExtraBlock(),
		Table:      *NewTable(),
	}
}

func (self *WProcess) Draw(buf *Buffer) {
	self.ExtraBlock.Draw(buf)

	self.Table.Render(buf, &self.ExtraBlock)
	self.FormattedText.render(buf, &self.ExtraBlock)
	self.CMD.render(buf, &self.ExtraBlock)
}

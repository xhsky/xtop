package widgets

import (
	. "github.com/gizak/termui/v3"
	//log "github.com/sirupsen/logrus"
)

type WCpu struct {
	ExtraBlock
	Graph
	Line
}

func NewWCpu() *WCpu {
	return &WCpu{
		ExtraBlock: *NewExtraBlock(),
	}
}

func (self *WCpu) Draw(buf *Buffer) {
	self.ExtraBlock.Draw(buf)

	drawArea := self.Inner
	if self.Inner.Dy() > 6 {
		self.Line.render(buf, drawArea, &self.ExtraBlock)
	}
	self.Graph.render(buf, &self.ExtraBlock)
}

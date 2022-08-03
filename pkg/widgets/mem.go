package widgets

import (
	. "github.com/gizak/termui/v3"
	//log "github.com/sirupsen/logrus"
)

type WMem struct {
	ExtraBlock
	Graph
	BarChart
}

func NewWMem() *WMem {
	return &WMem{
		ExtraBlock: *NewExtraBlock(),
	}
}

func (self *WMem) Draw(buf *Buffer) {
	self.ExtraBlock.Draw(buf)

	self.BarChart.render(buf, &self.ExtraBlock)
	self.Graph.render(buf, &self.ExtraBlock)
}

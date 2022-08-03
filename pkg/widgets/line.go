// Copyright 2017 Zack Guo <zack.y.guo@gmail.com>. All rights reserved.
// Use of this source code is governed by a MIT license that can
// be found in the LICENSE file.

package widgets

import (
	"image"

	. "github.com/gizak/termui/v3"
	//log "github.com/sirupsen/logrus"
)

type Line struct {
	Data   [][]float64
	MaxVal float64

	LineColors []Color

	//HorizontalScale int
}

func (self *Line) render(buf *Buffer, drawArea image.Rectangle, block *ExtraBlock) {
	maxVal := self.MaxVal

	canvas := NewCanvas()
	canvas.Rectangle = drawArea

	for i, line := range self.Data {
		previousHeight := int((line[1] / maxVal) * float64(drawArea.Dy()-1))
		for j, val := range line[1:] {
			height := int((val / maxVal) * float64(drawArea.Dy()-1))
			canvas.SetLine(
				image.Pt(
					//(drawArea.Min.X+(j*self.HorizontalScale))*2,
					(drawArea.Min.X+j)*2,
					(drawArea.Max.Y-previousHeight-1)*4,
				),
				image.Pt(
					//(drawArea.Min.X+((j+1)*self.HorizontalScale))*2,
					(drawArea.Min.X+(j+1))*2,
					(drawArea.Max.Y-height-1)*4,
				),
				SelectColor(self.LineColors, i),
			)
			previousHeight = height
		}
	}
	canvas.Draw(buf)
}

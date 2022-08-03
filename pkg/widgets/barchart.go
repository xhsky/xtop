package widgets

import (
	"image"

	. "github.com/gizak/termui/v3"
	//log "github.com/sirupsen/logrus"
)

var (
	maxWidth = 9
	gap      int
)

type BarChart struct {
	BarColors     []Color
	Data          []float64
	Str           []string
	StrStyles     []Style
	StrMinWidth   int
	Labels        []string
	LabelStyles   []Style
	LabelMinWidth int
	BarMinWidth   int // 7
	MaxVal        float64
}

func (self BarChart) render(buf *Buffer, block *ExtraBlock) {
	nums := len(self.Data)
	width := block.Inner.Dx() / nums

	XBar := block.Inner.Min.X
	if width > maxWidth {
		gap = width - maxWidth
		width = maxWidth
	} else {
		gap = 1
		width = width - gap
	}

	if width >= 1 {
		for i, data := range self.Data {
			// draw bar
			height := int((data / self.MaxVal) * float64(block.Inner.Dy()))
			YBar := block.Inner.Max.Y - height

			for x := XBar; x < XBar+width; x++ {
				for y := block.Inner.Max.Y - 1; y > YBar; y-- {
					c := NewCell(' ', NewStyle(ColorClear, SelectColor(self.BarColors, i)))
					buf.SetCell(c, image.Pt(x, y))
				}
			}

			// draw label
			if width >= 1 && width <= self.LabelMinWidth {
				for i := range self.Labels {
					self.Labels[i] = self.Labels[i][0:1]
				}
			}
			buf.SetString(
				self.Labels[i],
				SelectStyle(self.LabelStyles, i),
				image.Pt(XBar, block.Inner.Max.Y),
			)

			// draw str
			if width >= self.StrMinWidth {
				if height == 0 {
					YBar = YBar - 1 // 避免覆盖label
				}
				buf.SetString(
					self.Str[i],
					NewStyle(
						SelectStyle(self.StrStyles, i).Fg,
					),
					image.Pt(XBar, YBar),
				)
			}
			XBar += (width + gap)
		}
	}
}

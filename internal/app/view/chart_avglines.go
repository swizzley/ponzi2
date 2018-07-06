package view

import (
	"image"
	"math"

	"github.com/btmura/ponzi2/internal/app/model"
	"github.com/btmura/ponzi2/internal/gfx"
)

// ChartAvgLines renders a moving average line for a single stock.
type ChartAvgLines struct {
	renderable bool
	lines      *gfx.VAO
}

// NewChartAvgLines creates a new ChartAvgLines.
func NewChartAvgLines() *ChartAvgLines {
	return &ChartAvgLines{}
}

// SetStock sets the ChartAvgLine's stock.
func (ch *ChartAvgLines) SetStock(st *model.Stock) {
	// Reset everything.
	ch.Close()

	// Bail out if there is no data yet.
	if st.DailyTradingSessionSeries == nil || st.DailyMovingAverageSeries25 == nil || st.DailyMovingAverageSeries50 == nil || st.DailyMovingAverageSeries200 == nil {
		return // Stock has no data yet.
	}

	ch.lines = createChartAvgLinesVAO(st)
	ch.renderable = true
}

func createChartAvgLinesVAO(st *model.Stock) *gfx.VAO {
	minPrice := float32(math.MaxFloat32)
	maxPrice := float32(0)
	for _, s := range st.DailyTradingSessionSeries.TradingSessions {
		if minPrice > s.Low {
			minPrice = s.Low
		}
		if maxPrice < s.High {
			maxPrice = s.High
		}
	}

	data := &gfx.VAOVertexData{}
	var v uint16 // vertex index

	add := func(ma *model.MovingAverageSeries, color [3]float32) {
		dx := 2.0 / float32(len(ma.MovingAverages)) // (-1 to 1) on X-axis
		calcX := func(i int) float32 {
			return -1.0 + dx*float32(i) + dx*0.5
		}
		calcY := func(v float32) float32 {
			return 2.0*(v-minPrice)/(maxPrice-minPrice) - 1.0
		}

		first := true
		for i, mv := range ma.MovingAverages {
			if mv.Value < minPrice || mv.Value > maxPrice {
				continue
			}
			data.Vertices = append(data.Vertices, calcX(i), calcY(mv.Value), 0)
			data.Colors = append(data.Colors, color[0], color[1], color[2])
			if !first {
				data.Indices = append(data.Indices, v, v-1)
			}
			v++
			first = false
		}
	}

	add(st.DailyMovingAverageSeries25, purple)
	add(st.DailyMovingAverageSeries50, yellow)
	add(st.DailyMovingAverageSeries200, white)

	return gfx.NewVAO(gfx.Lines, data)
}

// Render renders the ChartAvgLine.
func (ch *ChartAvgLines) Render(r image.Rectangle) {
	if !ch.renderable {
		return
	}

	gfx.SetModelMatrixRect(r)
	ch.lines.Render()
}

// Close frees the resources backing the ChartAvgLine.
func (ch *ChartAvgLines) Close() {
	ch.renderable = false
	if ch.lines != nil {
		ch.lines.Delete()
	}
}
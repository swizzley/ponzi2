package app

import (
	"bytes"
	"image"
	math "math"

	"github.com/btmura/ponzi2/internal/gfx"
)

// White horizontal and vertical line VAOs that can be reused anywhere.
var (
	horizLine = gfx.HorizColoredLineVAO(white, white)
	vertLine  = gfx.VertColoredLineVAO(white, white)
)

var (
	roundedCornerNWFaces = gfx.ReadPLYVAO(bytes.NewReader(MustAsset("roundedCornerNWFaces.ply")))
	roundedCornerNWEdges = gfx.ReadPLYVAO(bytes.NewReader(MustAsset("roundedCornerNWEdges.ply")))
)

// fillRoundedRect renders a filled rounded rectangle within r.
func fillRoundedRect(r image.Rectangle, rounding int) {
	// NORTHWEST Corner
	gfx.SetModelMatrixRect(image.Rect(r.Min.X, r.Max.Y-rounding, r.Min.X+rounding, r.Max.Y))
	roundedCornerNWFaces.Render()

	// NORTHEAST Corner
	gfx.SetModelMatrixRotatedRect(image.Rect(r.Max.X-rounding, r.Max.Y-rounding, r.Max.X, r.Max.Y), -math.Pi/2)
	roundedCornerNWFaces.Render()

	// SOUTHEAST Corner
	gfx.SetModelMatrixRotatedRect(image.Rect(r.Max.X-rounding, r.Min.Y, r.Max.X, r.Min.Y+rounding), -math.Pi)
	roundedCornerNWFaces.Render()

	// SOUTHWEST Corner
	gfx.SetModelMatrixRotatedRect(image.Rect(r.Min.X, r.Min.Y, r.Min.X+rounding, r.Min.Y+rounding), -3*math.Pi/2)
	roundedCornerNWFaces.Render()
}

// strokeRoundedRect renders a stroked rounded rectangle within r.
func strokeRoundedRect(r image.Rectangle, rounding int) {
	// NORTHWEST Corner
	gfx.SetModelMatrixRect(image.Rect(r.Min.X, r.Max.Y-rounding, r.Min.X+rounding, r.Max.Y))
	roundedCornerNWEdges.Render()

	// NORTHEAST Corner
	gfx.SetModelMatrixRotatedRect(image.Rect(r.Max.X-rounding, r.Max.Y-rounding, r.Max.X, r.Max.Y), -math.Pi/2)
	roundedCornerNWEdges.Render()

	// SOUTHEAST Corner
	gfx.SetModelMatrixRotatedRect(image.Rect(r.Max.X-rounding, r.Min.Y, r.Max.X, r.Min.Y+rounding), -math.Pi)
	roundedCornerNWEdges.Render()

	// SOUTHWEST Corner
	gfx.SetModelMatrixRotatedRect(image.Rect(r.Min.X, r.Min.Y, r.Min.X+rounding, r.Min.Y+rounding), -3*math.Pi/2)
	roundedCornerNWEdges.Render()

	// fudge is how much to extend the borders to close gaps in OpenGL rendering.
	const fudge = 2

	// TOP border
	hMinX, hMaxX := r.Min.X+rounding-fudge, r.Max.X-rounding+fudge
	gfx.SetModelMatrixRect(image.Rect(hMinX, r.Max.Y, hMaxX, r.Max.Y))
	horizLine.Render()

	// BOTTOM border
	gfx.SetModelMatrixRect(image.Rect(hMinX, r.Min.Y, hMaxX, r.Min.Y))
	horizLine.Render()

	// LEFT border
	vMinX, vMaxX := r.Min.Y+rounding-fudge, r.Max.Y-rounding+fudge
	gfx.SetModelMatrixRect(image.Rect(r.Min.X, vMinX, r.Min.X, vMaxX))
	vertLine.Render()

	// RIGHT border
	gfx.SetModelMatrixRect(image.Rect(r.Max.X, vMinX, r.Max.X, vMaxX))
	vertLine.Render()
}

// sliceRenderHorizDividers horizontally cuts a rectangle from the bottom at
// the percentage amounts and draws the VAO at those percentages.
func sliceRenderHorizDividers(r image.Rectangle, dividerVAO *gfx.VAO, percentages ...float32) []image.Rectangle {
	rects := sliceRect(r, percentages...)
	for _, r := range rects {
		renderHorizDivider(r, dividerVAO)
	}
	return rects
}

// renderHorizDivider renders a VAO in a single pixel horizontal
// rectangle at the top edge of the rectangle.
func renderHorizDivider(r image.Rectangle, dividerVAO *gfx.VAO) {
	gfx.SetModelMatrixRect(image.Rect(r.Min.X, r.Max.Y, r.Max.X, r.Max.Y))
	dividerVAO.Render()
}

// sliceRect horizontally cuts a rectangle from the bottom at the percentage
// amounts. It returns n rectangles given n percentages.
func sliceRect(r image.Rectangle, percentages ...float32) []image.Rectangle {
	var rs []image.Rectangle
	addRect := func(minY, maxY int) {
		rs = append(rs, image.Rect(r.Min.X, minY, r.Max.X, maxY))
	}

	cy := r.Min.Y // Start at the bottom and cut horizontally up.
	for _, p := range percentages {
		dy := int(float32(r.Dy()) * p)
		addRect(cy, cy+dy)
		cy += dy // Bump upwards.
	}

	return rs
}

// transRect returns a rectangle translated by the dx and dy amounts.
func transRect(r image.Rectangle, dx, dy int) image.Rectangle {
	return r.Add(image.Pt(dx, dy))
}
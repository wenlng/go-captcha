/**
 * @Author Awen
 * @Date 2024/05/01
 * @Email wengaolng@gmail.com
 **/

package canvas

import (
	"math"
)

// RotatePoint is the point of rotation
func RotatePoint(x, y, sin, cos float64) (float64, float64) {
	return x*cos - y*sin, x*sin + y*cos
}

// RotatedSize is the size of rotation
func RotatedSize(w, h int, angle float64) (int, int) {
	if w <= 0 || h <= 0 {
		return 0, 0
	}

	sin, cos := math.Sincos(math.Pi * angle / 180)
	x1, y1 := RotatePoint(float64(w-1), 0, sin, cos)
	x2, y2 := RotatePoint(float64(w-1), float64(h-1), sin, cos)
	x3, y3 := RotatePoint(0, float64(h-1), sin, cos)

	minX := math.Min(x1, math.Min(x2, math.Min(x3, 0)))
	maxX := math.Max(x1, math.Max(x2, math.Max(x3, 0)))
	minY := math.Min(y1, math.Min(y2, math.Min(y3, 0)))
	maxY := math.Max(y1, math.Max(y2, math.Max(y3, 0)))

	width := maxX - minX + 1
	if width-math.Floor(width) > 0.1 {
		width++
	}
	height := maxY - minY + 1
	if height-math.Floor(height) > 0.1 {
		height++
	}

	return int(width), int(height)
}

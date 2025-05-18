/**
 * @Author Awen
 * @Date 2024/06/01
 * @Email wengaolng@gmail.com
 **/

package click

import (
	"math"
)

// Validate checks if a click point is within the specified area
// params:
//   - sx, sy: Coordinates of the click point
//   - dx, dy: Top-left coordinates of the target area
//   - width, height: Width and height of the target area
//   - padding: Padding of the area
//
// return: Whether the point is within the area
func Validate(sx, sy, dx, dy, width, height, padding int) bool {
	newWidth := width + (padding * 2)
	newHeight := height + (padding * 2)
	newDx := int(math.Max(float64(dx), float64(dx-padding)))
	newDy := int(math.Max(float64(dy), float64(dy-padding)))

	return sx >= newDx &&
		sx <= newDx+newWidth &&
		sy >= newDy &&
		sy <= newDy+newHeight
}

// Deprecated: As of 2.1.0, it will be removed, please use [click.Validate]
func CheckPoint(sx, sy, dx, dy, width, height, padding int64) bool {
	newWidth := width + (padding * 2)
	newHeight := height + (padding * 2)
	newDx := int64(math.Max(float64(dx), float64(dx-padding)))
	newDy := int64(math.Max(float64(dy), float64(dy-padding)))

	return sx >= newDx &&
		sx <= newDx+newWidth &&
		sy >= newDy &&
		sy <= newDy+newHeight
}

/**
 * @Author Awen
 * @Date 2024/06/01
 * @Email wengaolng@gmail.com
 **/

package click

import (
	"math"
)

// CheckPoint is to the position of the detection point
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

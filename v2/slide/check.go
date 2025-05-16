/**
 * @Author Awen
 * @Date 2024/06/01
 * @Email wengaolng@gmail.com
 **/

package slide

// CheckPoint checks if the point position is within the specified range
// params:
//   - sx: Source X coordinate
//   - sy: Source Y coordinate
//   - dx: Target X coordinate
//   - dy: Target Y coordinate
//   - padding: Padding
//
// return: Whether within range
func CheckPoint(sx, sy, dx, dy, padding int64) bool {
	newX := padding * 2
	newY := padding * 2
	newDx := dx - padding
	newDy := dy - padding

	return sx >= newDx &&
		sx <= newDx+newX &&
		sy >= newDy &&
		sy <= newDy+newY
}

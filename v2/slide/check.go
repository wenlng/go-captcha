/**
 * @Author Awen
 * @Date 2024/06/01
 * @Email wengaolng@gmail.com
 **/

package slide

// CheckData checks if the point position is within the specified range
// params:
//   - sx: Source X coordinate
//   - sy: Source Y coordinate
//   - dx: Target X coordinate
//   - dy: Target Y coordinate
//   - padding: Padding
//
// return: Whether within range
func CheckData(sx, sy, dx, dy, padding int) bool {
	newX := padding * 2
	newY := padding * 2
	newDx := dx - padding
	newDy := dy - padding

	return sx >= newDx &&
		sx <= newDx+newX &&
		sy >= newDy &&
		sy <= newDy+newY
}

// Deprecated: As of 2.1.0, it will be removed, please use [CheckData]
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

/**
 * @Author Awen
 * @Date 2024/06/01
 * @Email wengaolng@gmail.com
 **/

package rotate

// CheckData checks if the rotation angle is within the specified range
// params:
//   - angle: Current angle
//   - dAngle: Target angle
//   - padding: Angle padding
//
// return: Whether within range
func CheckData(angle, dAngle, padding int) bool {
	minAngle := 360 - padding
	maxAngle := 360 + padding
	angle += dAngle

	return angle >= minAngle && angle <= maxAngle
}

// Deprecated: As of 2.1.0, it will be removed, please use [CheckData]
func CheckAngle(angle, dAngle, padding int64) bool {
	minAngle := 360 - padding
	maxAngle := 360 + padding
	angle += dAngle

	return angle >= minAngle && angle <= maxAngle
}

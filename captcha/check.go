/**
 * @Author Awen
 * @Description Captcha
 * @Date 2021/12/20
 * @Email wengaolng@gmail.com
 **/

package captcha

// CheckPointDist is a function
/**
 * @Description: Calculate the distance between two points
 * @param sx
 * @param sy
 * @param dx
 * @param dy
 * @param width
 * @param height
 * @return bool
 */
func CheckPointDist(sx, sy, dx, dy, width, height int64) bool {
	return sx >= dx &&
		sx <= dx + width &&
		sy <= dy &&
		sy >= dy - height
}
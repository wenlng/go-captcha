/**
 * @Author Awen
 * @Description Captcha
 * @Date 2021/12/20
 * @Email wengaolng@gmail.com
 **/

package captcha

import "math"

// CheckPointDist is a function
/**
 * @Description: 计算点的位置在区域是否命中
 * @param sx		用户点击的x轴
 * @param sy		用户点击的y轴
 * @param dx		校验文本的x轴
 * @param dy		校验文本的y轴
 * @param width		校验文本的宽度
 * @param height	校验文本的高度
 * @return bool
 */
func CheckPointDist(sx, sy, dx, dy, width, height int64) bool {
	return sx >= dx &&
		sx <= dx + width &&
		sy <= dy &&
		sy >= dy - height
}

// CheckPointDistWithPadding is a function
/**
 * @Description: 计算点的位置在扩张区域(原区域+外边距)是否命中
 * @param sx		用户点击的x轴
 * @param sy		用户点击的y轴
 * @param dx		校验文本的x轴
 * @param dy		校验文本的y轴
 * @param width		校验文本的宽度
 * @param height	校验文本的高度
 * @param padding 	在原有的区域上添加额外边距进行扩张计算区域，不推荐设置padding
 * @return bool
 */
func CheckPointDistWithPadding(sx, sy, dx, dy, width, height, padding int64) bool {
	newWidth := width + (padding * 2)
	newHeight := height + (padding * 2)
	newDx := int64(math.Max(float64(dx), float64(dx - padding)))
	newDy := dy + padding

	return sx >= newDx &&
		sx <= newDx + newWidth &&
		sy <= newDy &&
		sy >= newDy - newHeight
}
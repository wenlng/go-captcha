/**
 * @Author Awen
 * @Description
 * @Date 2021/7/20
 **/

package main

import (
	"github.com/wenlng/go-captcha/captcha"
	"os"
)

/**
 * @Description: 获取当前目录
 * @return string
 */
func getPWD() string {
	path, err := os.Getwd()
	if err != nil {
		return ""
	}
	return path + "/.."
}

/**
 * @Description: 获取Captcha
 * @return *captcha.Captcha
 */
func getCaptcha() *captcha.Captcha {
	capt := captcha.GetCaptcha()

	//capt.SetFont([]string{
	//	getPWD() + "/__example/resources/fonts/fzshengsksjw_cu.ttf",
	//	getPWD() + "/__example/resources/fonts/hyrunyuan.ttf",
	//})
	//
	//capt.SetBackground([]string{
	//	getPWD() + "/__example/resources/images/1.jpg",
	//	getPWD() + "/__example/resources/images/2.jpg",
	//	getPWD() + "/__example/resources/images/3.jpg",
	//	getPWD() + "/__example/resources/images/4.jpg",
	//	getPWD() + "/__example/resources/images/5.jpg",
	//})

	return capt
}

// GetDraw is a function
/**
 * @Description: 获取Draw
 * @return *captcha.Draw
 */
func GetDraw() *captcha.Draw {
	return &captcha.Draw{}
}

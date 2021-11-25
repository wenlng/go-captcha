/**
 * @Author Awen
 * @Description
 * @Date 2021/7/19
 **/

package captcha

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"image"
	"image/png"
)

/**
 * @Description: 二进制编码
 * @param img
 * @return []byte
 */
func binaryEncoding(img image.Image) []byte {
	var buf bytes.Buffer
	if err := png.Encode(&buf, img); err != nil {
		panic(err.Error())
	}
	return buf.Bytes()
}

// EncodeB64string is a function
/**
 * @Description: base64编码
 * @param img
 * @return string
 */
func EncodeB64string(img image.Image) string {
	return fmt.Sprintf("data:%s;base64,%s", "image/png", base64.StdEncoding.EncodeToString(binaryEncoding(img)))
}

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
	"image/jpeg"
	"image/png"
)

/**
 * @Description: 图片编码二进制
 * @param img
 * @return []byte
 */
func encodingImageToBinary(img image.Image) (ret []byte) {
	var buf bytes.Buffer
	if err := png.Encode(&buf, img); err != nil {
		panic(err.Error())
	}
	ret = buf.Bytes()
	buf.Reset()
	return
}


/**
 * @Description: 二进制编码图片
 * @param []byte
 * @return image.Image
 * @return error
 */
func decodingBinaryToImage(b []byte) (img image.Image, err error) {
	var buf bytes.Buffer
	buf.Write(b)
	img, err = jpeg.Decode(&buf)
	buf.Reset()
	return
}

// EncodeB64string is a function
/**
 * @Description: base64编码
 * @param img
 * @return string
 */
func EncodeB64string(img image.Image) string {
	return fmt.Sprintf("data:%s;base64,%s", "image/png", base64.StdEncoding.EncodeToString(encodingImageToBinary(img)))
}

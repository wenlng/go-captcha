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
 * @Description: 图片编码二进制，PNG格式
 * @param img
 * @return []byte
 */
func encodingImageToBinaryWithPng(img image.Image) (ret []byte) {
	var buf bytes.Buffer
	if err := png.Encode(&buf, img); err != nil {
		panic(err.Error())
	}
	ret = buf.Bytes()
	buf.Reset()
	return
}


/**
 * @Description: 图片编码二进制，IMAGE格式
 * @param img
 * @param quality
 * @return []byte
 */
func encodingImageToBinaryWithJpeg(img image.Image, quality int) (ret []byte) {
	var buf bytes.Buffer
	if err := jpeg.Encode(&buf, img, &jpeg.Options{Quality: quality}); err != nil {
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
func EncodeB64stringWithPng(img image.Image) string {
	return fmt.Sprintf("data:%s;base64,%s", "image/png", base64.StdEncoding.EncodeToString(encodingImageToBinaryWithPng(img)))
}

// EncodeB64string is a function
/**
 * @Description: 	base64编码
 * @param img		图片
 * @param quality 	清晰度
 * @return string
 */
func EncodeB64stringWithJpeg(img image.Image, quality int) string {
	return fmt.Sprintf("data:%s;base64,%s", "image/jpeg", base64.StdEncoding.EncodeToString(encodingImageToBinaryWithJpeg(img, quality)))
}

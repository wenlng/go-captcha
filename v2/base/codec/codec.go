/**
 * @Author Awen
 * @Date 2024/06/01
 * @Email wengaolng@gmail.com
 **/

package codec

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"image"
	"image/jpeg"
	"image/png"
)

// EncodePNGToByte is to encode the png into a byte array
func EncodePNGToByte(img image.Image) (ret []byte, err error) {
	var buf bytes.Buffer
	if err = png.Encode(&buf, img); err != nil {
		return
	}
	ret = buf.Bytes()
	buf.Reset()
	return
}

// EncodeJPEGToByte is to encode the image into a byte array
func EncodeJPEGToByte(img image.Image, quality int) (ret []byte, err error) {
	var buf bytes.Buffer
	if err = jpeg.Encode(&buf, img, &jpeg.Options{Quality: quality}); err != nil {
		return
	}
	ret = buf.Bytes()
	buf.Reset()
	return
}

// DecodeByteToJpeg is to decode the byte array into an image
func DecodeByteToJpeg(b []byte) (img image.Image, err error) {
	var buf bytes.Buffer
	buf.Write(b)
	img, err = jpeg.Decode(&buf)
	buf.Reset()
	return
}

// DecodeByteToPng is to decode the byte array into a png
func DecodeByteToPng(b []byte) (img image.Image, err error) {
	var buf bytes.Buffer
	buf.Write(b)
	img, err = png.Decode(&buf)
	buf.Reset()
	return
}

// EncodePNGToBase64 is to encode the png into string
func EncodePNGToBase64(img image.Image) (string, error) {
	b, err := EncodePNGToByte(img)
	if err != nil {
		return ``, err
	}
	return fmt.Sprintf("data:%s;base64,%s", "image/png", base64.StdEncoding.EncodeToString(b)), err
}

// EncodeJPEGToBase64 is to encode the image into string
func EncodeJPEGToBase64(img image.Image, quality int) (string, error) {
	b, err := EncodeJPEGToByte(img, quality)
	if err != nil {
		return ``, err
	}
	return fmt.Sprintf("data:%s;base64,%s", "image/jpeg", base64.StdEncoding.EncodeToString(b)), err
}

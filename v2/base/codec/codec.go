/**
 * @Author Awen
 * @Date 2024/06/01
 * @Email wengaolng@gmail.com
 **/

package codec

import (
	"bytes"
	"encoding/base64"
	"image"
	"image/jpeg"
	"image/png"
	"sync"
)

const (
	pngBasePrefix  = "data:image/png;base64,"
	jpegBasePrefix = "data:image/jpeg;base64,"
	maxBufferSize  = 10 * 1024 * 1024 // 10MB
)

var (
	// bufferPool is a pool of bytes.Buffer to reduce memory allocations
	bufferPool = sync.Pool{
		New: func() interface{} {
			return &bytes.Buffer{}
		},
	}
)

// getBuffer returns a buffer from the pool
func getBuffer() *bytes.Buffer {
	return bufferPool.Get().(*bytes.Buffer)
}

// putBuffer returns a buffer to the pool after resetting it
func putBuffer(buf *bytes.Buffer) {
	if buf != nil {
		buf.Reset()
		if buf.Cap() < maxBufferSize {
			bufferPool.Put(buf)
		}
	}
}

// EncodePNGToByte encodes a PNG image to a byte array
func EncodePNGToByte(img image.Image) ([]byte, error) {
	if img == nil {
		return nil, nil
	}

	buf := getBuffer()
	defer putBuffer(buf)

	bounds := img.Bounds()
	estimatedSize := bounds.Dx() * bounds.Dy() * 4
	if estimatedSize > 0 && estimatedSize < maxBufferSize {
		buf.Grow(estimatedSize)
	}

	if err := png.Encode(buf, img); err != nil {
		return nil, err
	}

	result := make([]byte, buf.Len())
	copy(result, buf.Bytes())
	return result, nil
}

// EncodeJPEGToByte encodes a JPEG image to a byte array
func EncodeJPEGToByte(img image.Image, quality int) ([]byte, error) {
	if img == nil {
		return nil, nil
	}

	buf := getBuffer()
	defer putBuffer(buf)

	bounds := img.Bounds()
	estimatedSize := bounds.Dx() * bounds.Dy() * 3 / 2
	if estimatedSize > 0 && estimatedSize < maxBufferSize {
		buf.Grow(estimatedSize)
	}

	if err := jpeg.Encode(buf, img, &jpeg.Options{Quality: quality}); err != nil {
		return nil, err
	}

	result := make([]byte, buf.Len())
	copy(result, buf.Bytes())
	return result, nil
}

// DecodeByteToJpeg decodes a byte array to a JPEG image
func DecodeByteToJpeg(b []byte) (img image.Image, err error) {
	var buf bytes.Buffer
	buf.Write(b)
	img, err = jpeg.Decode(&buf)
	buf.Reset()
	return
}

// DecodeByteToPng decodes a byte array to a PNG image
func DecodeByteToPng(b []byte) (img image.Image, err error) {
	var buf bytes.Buffer
	buf.Write(b)
	img, err = png.Decode(&buf)
	buf.Reset()
	return
}

// EncodePNGToBase64 encodes a PNG image to a Base64 string
func EncodePNGToBase64(img image.Image) (string, error) {
	base64Str, err := EncodePNGToBase64Data(img)
	if err != nil {
		return "", err
	}

	return pngBasePrefix + base64Str, nil
}

// EncodeJPEGToBase64 encodes a JPEG image to a Base64 string
func EncodeJPEGToBase64(img image.Image, quality int) (string, error) {
	base64Str, err := EncodeJPEGToBase64Data(img, quality)
	if err != nil {
		return "", err
	}

	return jpegBasePrefix + base64Str, nil
}

// EncodePNGToBase64Data encodes a PNG image to Base64 data (without prefix)
func EncodePNGToBase64Data(img image.Image) (string, error) {
	byteCode, err := EncodePNGToByte(img)
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(byteCode), nil
}

// EncodeJPEGToBase64Data encodes a JPEG image to Base64 data (without prefix)
func EncodeJPEGToBase64Data(img image.Image, quality int) (string, error) {
	byteCode, err := EncodeJPEGToByte(img, quality)
	if err != nil {
		return "", err
	}

	return base64.StdEncoding.EncodeToString(byteCode), nil
}

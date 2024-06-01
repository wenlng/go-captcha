/**
 * @Author Awen
 * @Date 2024/05/01
 * @Email wengaolng@gmail.com
 **/

package imagedata

import (
	"fmt"
	"image"

	"github.com/wenlng/go-captcha/capts/base/codec"
	"github.com/wenlng/go-captcha/capts/base/option"
)

// JPEGImageData .
type JPEGImageData interface {
	Get() image.Image
	ToBytes() []byte
	ToBytesWithQuality(imageQuality int) []byte
	ToBase64() string
	ToBase64WithQuality(imageQuality int) string
	SaveToFile(filepath string, quality int) error
}

var _ JPEGImageData = (*jpegImageDta)(nil)

// jpegImageDta .
type jpegImageDta struct {
	image image.Image
}

// NewJPEGImageData .
func NewJPEGImageData(img image.Image) JPEGImageData {
	return &jpegImageDta{
		image: img,
	}
}

// Get is get the original picture
func (c *jpegImageDta) Get() image.Image {
	return c.image
}

// SaveToFile is to save JPEG as a file
func (c *jpegImageDta) SaveToFile(filepath string, quality int) error {
	if c.image == nil {
		return fmt.Errorf("missing image data")
	}

	return saveToFile(c.image, filepath, false, quality)
}

// ToBytes is to convert JPEG into byte array
func (c *jpegImageDta) ToBytes() []byte {
	if c.image == nil {
		return []byte{}
	}

	return codec.EncodeJPEGToByte(c.image, option.QualityNone)
}

// ToBytesWithQuality is to convert JPEG into byte array with quality
func (c *jpegImageDta) ToBytesWithQuality(imageQuality int) []byte {
	if c.image == nil {
		return []byte{}
	}

	if imageQuality <= option.QualityNone && imageQuality >= option.QualityLevel5 {
		return codec.EncodeJPEGToByte(c.image, imageQuality)
	}
	return codec.EncodeJPEGToByte(c.image, option.QualityNone)
}

// ToBase64 is to convert JPEG into base64
func (c *jpegImageDta) ToBase64() string {
	if c.image == nil {
		return ""
	}

	return codec.EncodeJPEGToBase64(c.image, option.QualityNone)
}

// ToBase64WithQuality is to convert JPEG into base64 with quality
func (c *jpegImageDta) ToBase64WithQuality(imageQuality int) string {
	if c.image == nil {
		return ""
	}

	if imageQuality <= option.QualityNone && imageQuality >= option.QualityLevel5 {
		return codec.EncodeJPEGToBase64(c.image, imageQuality)
	}
	return codec.EncodeJPEGToBase64(c.image, option.QualityNone)
}

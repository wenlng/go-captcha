/**
 * @Author Awen
 * @Date 2024/06/01
 * @Email wengaolng@gmail.com
 **/

package imagedata

import (
	"image"

	"github.com/wenlng/go-captcha/v2/base/codec"
	"github.com/wenlng/go-captcha/v2/base/option"
)

// JPEGImageData interface for JPEG image data
type JPEGImageData interface {
	Get() image.Image
	ToBytes() ([]byte, error)
	ToBytesWithQuality(imageQuality int) ([]byte, error)
	ToBase64() (string, error)
	ToBase64WithQuality(imageQuality int) (string, error)
	ToBase64Data() (string, error)
	ToBase64DataWithQuality(imageQuality int) (string, error)
	SaveToFile(filepath string, quality int) error
}

var _ JPEGImageData = (*jpegImageDta)(nil)

// jpegImageDta struct for JPEG image data
type jpegImageDta struct {
	image image.Image
}

// NewJPEGImageData creates a new JPEG image data instance
func NewJPEGImageData(img image.Image) JPEGImageData {
	return &jpegImageDta{
		image: img,
	}
}

// Get retrieves the original image
func (c *jpegImageDta) Get() image.Image {
	return c.image
}

// SaveToFile saves the JPEG image to a file
func (c *jpegImageDta) SaveToFile(filepath string, quality int) error {
	if c.image == nil {
		return ImageMissingDataErr
	}

	return saveToFile(c.image, filepath, false, quality)
}

// ToBytes converts the JPEG image to a byte array
func (c *jpegImageDta) ToBytes() ([]byte, error) {
	if c.image == nil {
		return []byte{}, ImageEmptyErr
	}

	return codec.EncodeJPEGToByte(c.image, option.QualityNone)
}

// ToBytesWithQuality converts the JPEG image to a byte array with specified quality
func (c *jpegImageDta) ToBytesWithQuality(imageQuality int) ([]byte, error) {
	if c.image == nil {
		return []byte{}, ImageEmptyErr
	}

	if imageQuality <= option.QualityNone && imageQuality >= option.QualityLevel5 {
		return codec.EncodeJPEGToByte(c.image, imageQuality)
	}
	return codec.EncodeJPEGToByte(c.image, option.QualityNone)
}

// ToBase64Data converts the JPEG image to Base64 data (without prefix)
func (c *jpegImageDta) ToBase64Data() (string, error) {
	if c.image == nil {
		return "", ImageEmptyErr
	}

	return codec.EncodeJPEGToBase64Data(c.image, option.QualityNone)
}

// ToBase64DataWithQuality converts the JPEG image to Base64 data with specified quality (without prefix)
func (c *jpegImageDta) ToBase64DataWithQuality(imageQuality int) (string, error) {
	if c.image == nil {
		return "", ImageEmptyErr
	}

	if imageQuality <= option.QualityNone && imageQuality >= option.QualityLevel5 {
		return codec.EncodeJPEGToBase64Data(c.image, imageQuality)
	}
	return codec.EncodeJPEGToBase64Data(c.image, option.QualityNone)
}

// ToBase64 converts the JPEG image to a Base64 string
func (c *jpegImageDta) ToBase64() (string, error) {
	if c.image == nil {
		return "", ImageEmptyErr
	}

	return codec.EncodeJPEGToBase64(c.image, option.QualityNone)
}

// ToBase64WithQuality converts the JPEG image to a Base64 string with specified quality
func (c *jpegImageDta) ToBase64WithQuality(imageQuality int) (string, error) {
	if c.image == nil {
		return "", ImageEmptyErr
	}

	if imageQuality <= option.QualityNone && imageQuality >= option.QualityLevel5 {
		return codec.EncodeJPEGToBase64(c.image, imageQuality)
	}
	return codec.EncodeJPEGToBase64(c.image, option.QualityNone)
}

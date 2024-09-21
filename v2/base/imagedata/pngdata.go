/**
 * @Author Awen
 * @Date 2024/06/01
 * @Email wengaolng@gmail.com
 **/

package imagedata

import (
	"fmt"
	"image"

	"github.com/wenlng/go-captcha/v2/base/codec"
	"github.com/wenlng/go-captcha/v2/base/option"
)

// PNGImageData .
type PNGImageData interface {
	Get() image.Image
	ToBytes() []byte
	ToBase64() string
	SaveToFile(filepath string) error
}

var _ PNGImageData = (*pngImageDta)(nil)

// pngImageDta .
type pngImageDta struct {
	image image.Image
}

// NewPNGImageData .
func NewPNGImageData(img image.Image) PNGImageData {
	return &pngImageDta{
		image: img,
	}
}

// Get is to get the original picture
func (c *pngImageDta) Get() image.Image {
	return c.image
}

// SaveToFile is to save PNG as a file
func (c *pngImageDta) SaveToFile(filepath string) error {
	if c.image == nil {
		return fmt.Errorf("missing image data")
	}

	return saveToFile(c.image, filepath, true, option.QualityNone)
}

// ToBytes is to convert PNG into byte array
func (c *pngImageDta) ToBytes() []byte {
	if c.image == nil {
		return []byte{}
	}

	return codec.EncodePNGToByte(c.image)
}

// ToBase64 is to convert PNG into base64
func (c *pngImageDta) ToBase64() string {
	if c.image == nil {
		return ""
	}
	return codec.EncodePNGToBase64(c.image)
}

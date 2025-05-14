/**
 * @Author Awen
 * @Date 2024/06/01
 * @Email wengaolng@gmail.com
 **/

package click

import "github.com/wenlng/go-captcha/v2/base/imagedata"

// CaptchaData defines the interface for captcha data
type CaptchaData interface {
	GetData() map[int]*Dot
	GetMasterImage() imagedata.JPEGImageData
	GetThumbImage() imagedata.PNGImageData
}

// CaptData is the concrete implementation of the CaptchaData interface
type CaptData struct {
	dots        map[int]*Dot
	masterImage imagedata.JPEGImageData
	thumbImage  imagedata.PNGImageData
}

var _ CaptchaData = (*CaptData)(nil)

// GetData gets the dot data of the captcha
// return: Map of dot data
func (c CaptData) GetData() map[int]*Dot {
	return c.dots
}

// GetMasterImage gets the main captcha image
// return: Main image in JPEG format
func (c CaptData) GetMasterImage() imagedata.JPEGImageData {
	return c.masterImage
}

// GetThumbImage gets the thumbnail image
// return: Thumbnail image in PNG format
func (c CaptData) GetThumbImage() imagedata.PNGImageData {
	return c.thumbImage
}

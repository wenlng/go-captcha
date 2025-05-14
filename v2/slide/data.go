/**
 * @Author Awen
 * @Date 2024/06/01
 * @Email wengaolng@gmail.com
 **/

package slide

import "github.com/wenlng/go-captcha/v2/base/imagedata"

// CaptchaData defines the interface for slide CAPTCHA data
type CaptchaData interface {
	GetData() *Block
	GetMasterImage() imagedata.JPEGImageData
	GetTileImage() imagedata.PNGImageData
}

// CaptData is the concrete implementation of the CaptchaData interface
type CaptData struct {
	block       *Block
	masterImage imagedata.JPEGImageData
	tileImage   imagedata.PNGImageData
}

var _ CaptchaData = (*CaptData)(nil)

// GetData gets the block data of the CAPTCHA
// return: Pointer to block data
func (c CaptData) GetData() *Block {
	return c.block
}

// GetMasterImage gets the main CAPTCHA image
// return: Main image in JPEG format
func (c CaptData) GetMasterImage() imagedata.JPEGImageData {
	return c.masterImage
}

// GetTileImage gets the tile image
// return: Tile image in PNG format
func (c CaptData) GetTileImage() imagedata.PNGImageData {
	return c.tileImage
}

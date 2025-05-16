/**
 * @Author Awen
 * @Date 2024/06/01
 * @Email wengaolng@gmail.com
 **/

package rotate

import (
	"image"

	"github.com/wenlng/go-captcha/v2/base/canvas"
	"github.com/wenlng/go-captcha/v2/base/randgen"
	"golang.org/x/image/draw"
)

// DrawImageParams defines the parameters for drawing the main image
type DrawImageParams struct {
	Rotate     int
	SquareSize int
	Background image.Image
	Alpha      float32
}

// DrawCropCircleImageParams defines the parameters for drawing a cropped circle image
type DrawCropCircleImageParams struct {
	ScaleRatioSize int
	Rotate         int
	SquareSize     int
	Background     image.Image
	Alpha          float32
}

// DrawImage defines the interface for drawing images
type DrawImage interface {
	DrawWithNRGBA(params *DrawImageParams) (img image.Image, err error)
	DrawWithCropCircle(params *DrawCropCircleImageParams) (image.Image, error)
}

var _ DrawImage = (*drawImage)(nil)

// drawImage is the concrete implementation of the DrawImage interface
type drawImage struct {
}

// NewDrawImage creates a new DrawImage instance
// return: DrawImage interface instance
func NewDrawImage() DrawImage {
	return &drawImage{}
}

// DrawWithCropCircle draws a cropped circle image (thumbnail)
// params:
//   - params: Drawing parameters
//
// returns:
//   - image.Image: Drawn thumbnail image
//   - error: Error information
func (d *drawImage) DrawWithCropCircle(params *DrawCropCircleImageParams) (image.Image, error) {
	bgImage := params.Background

	bgBounds := bgImage.Bounds()
	cvs := canvas.CreateNRGBACanvas(bgImage.Bounds().Dx(), bgImage.Bounds().Dy(), true)
	draw.Draw(cvs.Get(), bgImage.Bounds(), bgImage, image.Point{}, draw.Over)
	cvs.CropScaleCircle(bgImage.Bounds().Dx()/2, bgImage.Bounds().Dy()/2, bgImage.Bounds().Dy()/2, params.ScaleRatioSize)
	cvs.Rotate(params.Rotate, true)

	cvBounds := cvs.Bounds()
	if cvBounds.Dy() > bgBounds.Dy() || cvBounds.Dx() > bgBounds.Dx() {
		newCvs := canvas.CreateNRGBACanvas(bgImage.Bounds().Dx(), bgImage.Bounds().Dy(), true)
		draw.Draw(newCvs.Get(), newCvs.Bounds(), cvs.Get(), image.Point{X: (cvBounds.Dy() - bgBounds.Dy()) / 2, Y: (cvBounds.Dx() - bgBounds.Dx()) / 2}, draw.Over)
		cvs = newCvs
	}

	return cvs.Get(), nil
}

// DrawWithNRGBA draws the main CAPTCHA image using NRGBA format
// params:
//   - params: Drawing parameters
//
// return:
//   - image.Image: Drawn image
//   - error: Error information
func (d *drawImage) DrawWithNRGBA(params *DrawImageParams) (img image.Image, err error) {
	var rcm = canvas.CreateNRGBACanvas(params.SquareSize, params.SquareSize, true)
	if params.Background != nil {
		bgImage := params.Background
		b := bgImage.Bounds()
		rc := canvas.CreateNRGBACanvas(b.Dx(), b.Dy(), true)
		point := randgen.RangCutImagePos(params.SquareSize, params.SquareSize, bgImage)
		draw.Draw(rc.Get(), b, bgImage, point, draw.Over)
		rc.SubImage(image.Rect(0, 0, params.SquareSize, params.SquareSize))
		draw.Draw(rcm.Get(), rcm.Bounds(), rc.Get(), image.Point{}, draw.Over)
	}

	rcm.CropCircle(rcm.Bounds().Dx()/2, rcm.Bounds().Dy()/2, rcm.Bounds().Dy()/2)
	return rcm.Get(), nil
}

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

// DrawImageParams .
type DrawImageParams struct {
	Rotate     int
	SquareSize int
	Background image.Image
	Alpha      float32
}

// DrawCropCircleImageParams .
type DrawCropCircleImageParams struct {
	ScaleRatioSize int
	Rotate         int
	SquareSize     int
	Background     image.Image
	Alpha          float32
}

// DrawImage .
type DrawImage interface {
	DrawWithNRGBA(params *DrawImageParams) (img image.Image, err error)
	DrawWithCropCircle(params *DrawCropCircleImageParams) (image.Image, error)
}

var _ DrawImage = (*drawImage)(nil)

// NewDrawImage .
func NewDrawImage() DrawImage {
	return &drawImage{}
}

// drawImage .
type drawImage struct {
}

// DrawWithCropCircle is to draw a crop circle
func (d *drawImage) DrawWithCropCircle(params *DrawCropCircleImageParams) (image.Image, error) {
	bgImage := params.Background
	bgBounds := bgImage.Bounds()
	cvs := canvas.CreateNRGBACanvas(bgImage.Bounds().Dx(), bgImage.Bounds().Dy(), true)
	draw.Draw(cvs.Get(), bgImage.Bounds(), bgImage, image.Point{}, draw.Over)
	cvs.CropCircle(bgImage.Bounds().Dx()/2, bgImage.Bounds().Dy()/2, bgImage.Bounds().Dy()/2, params.ScaleRatioSize)
	cvs.Rotate(params.Rotate)

	cvBounds := cvs.Bounds()
	if cvBounds.Dy() > bgBounds.Dy() || cvBounds.Dx() > bgBounds.Dx() {
		newCvs := canvas.CreateNRGBACanvas(bgImage.Bounds().Dx(), bgImage.Bounds().Dy(), true)
		draw.Draw(newCvs.Get(), newCvs.Bounds(), cvs.Get(), image.Point{X: (cvBounds.Dy() - bgBounds.Dy()) / 2, Y: (cvBounds.Dx() - bgBounds.Dx()) / 2}, draw.Over)
		cvs = newCvs
	}

	return cvs, nil
}

// DrawWithNRGBA is to draw with a NRGBA
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

	rcm.CropCircle(rcm.Bounds().Dx()/2, rcm.Bounds().Dy()/2, rcm.Bounds().Dy()/2, 0)
	return rcm, nil
}

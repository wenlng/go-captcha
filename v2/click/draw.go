/**
 * @Author Awen
 * @Date 2024/06/01
 * @Email wengaolng@gmail.com
 **/

package click

import (
	"image"
	"image/color"
	"math"
	mRand "math/rand"

	"github.com/golang/freetype"
	"github.com/wenlng/go-captcha/v2/base/canvas"
	"github.com/wenlng/go-captcha/v2/base/helper"
	"github.com/wenlng/go-captcha/v2/base/option"
	"github.com/wenlng/go-captcha/v2/base/randgen"
	"github.com/wenlng/go-captcha/v2/base/random"
	"golang.org/x/image/draw"
	"golang.org/x/image/font"
)

type DrawType int

const (
	DrawTypeString DrawType = iota
	DrawTypeImage
)

// DrawImageParams .
type DrawImageParams struct {
	Width                 int
	Height                int
	Background            image.Image
	BackgroundDistort     int
	BackgroundCirclesNum  int
	BackgroundSlimLineNum int
	Alpha                 float32
	FontHinting           font.Hinting
	CaptchaDrawDot        []*DrawDot
	ShowShadow            bool
	ShadowColor           string
	ShadowPoint           *option.Point
	ThumbDisturbAlpha     float32
}

type DrawImage interface {
	DrawWithNRGBA(params *DrawImageParams) (image.Image, error)
	DrawWithPalette(params *DrawImageParams, textColors []color.Color, bgColors []color.Color) (image.Image, error)
	DrawWithNRGBA2(params *DrawImageParams, textColors []color.Color, bgColors []color.Color) (image.Image, error)
}

var _ DrawImage = (*drawImage)(nil)

// drawImage .
type drawImage struct {
}

// NewDrawImage .
func NewDrawImage() DrawImage {
	return &drawImage{}
}

// DrawWithNRGBA is to draw with a NRGBA
func (d *drawImage) DrawWithNRGBA(params *DrawImageParams) (image.Image, error) {
	dots := params.CaptchaDrawDot
	cvs := canvas.CreateNRGBACanvas(params.Width, params.Height, true)

	for i := 0; i < len(dots); i++ {
		dot := dots[i]

		dotImage, areaPoint, err := d.DrawDotImage(dot, params)
		if err != nil {
			return nil, err
		}
		minX := areaPoint.MinX
		maxX := areaPoint.MaxX
		minY := areaPoint.MinY
		maxY := areaPoint.MaxY

		width := maxX - minX
		height := maxY - minY

		draw.Draw(cvs.Get(), image.Rect(dot.X, dot.Y, dot.X+width, dot.Y+height), dotImage, image.Point{X: minX, Y: minY}, draw.Over)

		dot.Height = height
		dot.Width = width
		dot.Dot.Height = height
		dot.Dot.Width = width
	}

	img := params.Background
	b := cvs.Bounds()
	m := canvas.CreateNRGBACanvas(b.Dx(), b.Dx(), true)
	point := randgen.RangCutImagePos(params.Width, params.Height, img)
	draw.Draw(m.Get(), b, img, point, draw.Src)
	draw.Draw(m.Get(), cvs.Bounds(), cvs, image.Point{}, draw.Over)
	m.SubImage(image.Rect(0, 0, params.Width, params.Height))
	return m, nil
}

// DrawWithPalette is to draw with a palette
func (d *drawImage) DrawWithPalette(params *DrawImageParams, tColors []color.Color, bgColors []color.Color) (image.Image, error) {
	dots := params.CaptchaDrawDot

	nBgColors := make([]color.Color, 0, len(bgColors))
	for _, bgColor := range bgColors {
		r, g, b, _ := bgColor.RGBA()
		aa := helper.FormatAlpha(params.ThumbDisturbAlpha)
		nBgColors = append(nBgColors, color.RGBA{R: uint8(r), G: uint8(g), B: uint8(b), A: aa})
	}

	var p = make([]color.Color, 0, len(tColors)+len(nBgColors)+1)
	p = append(p, color.RGBA{R: 0xFF, G: 0xFF, B: 0xFF, A: 0x00})
	p = append(p, tColors...)
	p = append(p, nBgColors...)

	cvs := canvas.NewPalette(image.Rect(0, 0, params.Width, params.Height), p)
	if params.BackgroundCirclesNum > 0 {
		d.randomFillWithCircles(cvs, params.BackgroundCirclesNum, 1, nBgColors)
	}
	if params.BackgroundSlimLineNum > 0 {
		d.randomDrawSlimLine(cvs, params.BackgroundSlimLineNum, nBgColors)
	}

	for i := 0; i < len(dots); i++ {
		dot := dots[i]
		cColor, _ := helper.ParseHexColor(dot.Color)
		var err error
		if dot.DrawType == DrawTypeImage {
			var dotImage canvas.NRGBA
			dotImage, err = d.DrawShapeImage(dot, cColor)
			if err != nil {
				return nil, err
			}

			dotImage.Rotate(dot.Angle)
			bgBounds := cvs.Get().Bounds()
			dotBounds := dotImage.Bounds()
			drawAt := image.Point{X: bgBounds.Dx() - dotBounds.Dx(), Y: bgBounds.Dy() - dotBounds.Dy()}
			draw.Draw(cvs.Get(), image.Rect(dot.X, drawAt.Y, drawAt.X+dotBounds.Dx(), drawAt.Y+dotBounds.Dy()), dotImage.Get(), image.Point{}, draw.Over)
		} else {
			pt := freetype.Pt(dot.X, dot.Y)
			err = cvs.DrawString(&canvas.DrawStringParams{
				Color:   cColor,
				Size:    dot.Size,
				Width:   dot.Width,
				Height:  dot.Height,
				FontDPI: dot.FontDPI,
				Text:    dot.Text,
				Font:    dot.Font,
			}, pt)
		}

		if err != nil {
			return cvs, err
		}
	}

	if params.Background != nil {
		img := params.Background
		b := img.Bounds()
		m := canvas.CreateNRGBACanvas(b.Dx(), b.Dy(), true)
		point := randgen.RangCutImagePos(params.Width, params.Height, img)
		draw.Draw(m.Get(), b, img, point, draw.Src)
		cvs.Distort(float64(random.RandInt(5, 10)), float64(random.RandInt(120, 200)))
		draw.Draw(m.Get(), cvs.Bounds(), cvs, image.Point{}, draw.Over)
		rc := m.Get().SubImage(image.Rect(0, 0, params.Width, params.Height)).(*image.NRGBA)
		return rc, nil
	}

	if params.BackgroundDistort > 0 {
		cvs.Distort(float64(random.RandInt(5, 10)), float64(params.BackgroundDistort))
	}

	return cvs, nil

}

// DrawWithNRGBA2 is to draw with a NRGBA
func (d *drawImage) DrawWithNRGBA2(params *DrawImageParams, tColors []color.Color, bgColors []color.Color) (image.Image, error) {
	dots := params.CaptchaDrawDot

	nBgColors := make([]color.Color, 0, len(bgColors))
	for _, bgColor := range bgColors {
		r, g, b, _ := bgColor.RGBA()
		aa := helper.FormatAlpha(params.ThumbDisturbAlpha)
		nBgColors = append(nBgColors, color.RGBA{R: uint8(r), G: uint8(g), B: uint8(b), A: aa})
	}

	var p = make([]color.Color, 0, len(tColors)+len(nBgColors)+1)
	p = append(p, color.RGBA{R: 0xFF, G: 0xFF, B: 0xFF, A: 0x00})
	p = append(p, tColors...)
	p = append(p, nBgColors...)

	ccvs := canvas.NewNRGBA(image.Rect(0, 0, params.Width, params.Height), true)
	if params.Background != nil {
		img := params.Background
		b := img.Bounds()
		m := canvas.CreateNRGBACanvas(b.Dx(), b.Dy(), true)
		point := randgen.RangCutImagePos(params.Width, params.Height, img)
		draw.Draw(m.Get(), b, img, point, draw.Src)
		rc := m.Get().SubImage(image.Rect(0, 0, params.Width, params.Height)).(*image.NRGBA)
		draw.Draw(ccvs.Get(), rc.Bounds(), rc, image.Point{}, draw.Over)
	}

	cvs := canvas.NewPalette(image.Rect(0, 0, params.Width, params.Height), p)
	if params.BackgroundCirclesNum > 0 {
		d.randomFillWithCircles(cvs, params.BackgroundCirclesNum, 1, nBgColors)
	}
	if params.BackgroundSlimLineNum > 0 {
		d.randomDrawSlimLine(cvs, params.BackgroundSlimLineNum, nBgColors)
	}
	if params.BackgroundDistort > 0 {
		cvs.Distort(float64(random.RandInt(5, 10)), float64(params.BackgroundDistort))
	}

	cvsBounds := cvs.Bounds()
	width := cvsBounds.Dx() / len(dots)
	for i := 0; i < len(dots); i++ {
		dot := dots[i]
		cColor, _ := helper.ParseHexColor(dot.Color)
		var err error
		if dot.DrawType == DrawTypeImage {
			var dotImage canvas.NRGBA
			dotImage, err = d.DrawShapeImage(dot, cColor)
			if err != nil {
				return nil, err
			}

			dotImage.Rotate(dot.Angle)

			bgBounds := ccvs.Get().Bounds()
			dotBounds := dotImage.Bounds()
			drawAt := image.Point{X: bgBounds.Dx() - dotBounds.Dx(), Y: bgBounds.Dy() - dotBounds.Dy()}
			draw.Draw(ccvs.Get(), image.Rect(dot.X, drawAt.Y, drawAt.X+dotBounds.Dx(), drawAt.Y+dotBounds.Dy()), dotImage.Get(), image.Point{}, draw.Over)
		} else {
			var cImage canvas.NRGBA
			cImage, err = d.DrawStringImage(dot, cColor)
			if err != nil {
				return nil, err
			}
			cImage.Rotate(dot.Angle)

			areaPoint := cImage.CalcMarginBlankArea()
			minX := areaPoint.MinX
			maxX := areaPoint.MaxX
			minY := areaPoint.MinY
			maxY := areaPoint.MaxY
			cImage.SubImage(image.Rect(minX, minY, maxX, maxY))
			bounds := cImage.Bounds()

			dx := int(math.Max(float64(width*i+width/bounds.Dx()), 8))
			dy := random.RandInt(1, cvsBounds.Dy()-bounds.Dy()-4)

			draw.Draw(ccvs.Get(), image.Rect(dx, dy, dx+bounds.Dx(), dy+bounds.Dy()), cImage, image.Point{X: bounds.Min.X, Y: bounds.Min.Y}, draw.Over)
		}

		if err != nil {
			return ccvs, err
		}
	}

	ncvs := canvas.NewNRGBA(image.Rect(0, 0, params.Width, params.Height), true)
	draw.Draw(ncvs.Get(), cvs.Bounds(), cvs, image.Point{}, draw.Over)
	draw.Draw(ccvs.Get(), ncvs.Bounds(), ncvs, image.Point{}, draw.Over)
	return ccvs, nil
}

// randomFillWithCircles is to draw circle randomly
func (d *drawImage) randomFillWithCircles(m canvas.Palette, n, maxRadius int, colorB []color.Color) {
	maxx := m.Bounds().Max.X
	maxy := m.Bounds().Max.Y
	for i := 0; i < n; i++ {
		co := randgen.RandColor(colorB)
		//co.A = uint8(0xee)
		r := random.RandInt(1, maxRadius)
		m.DrawCircle(random.RandInt(r, maxx-r), random.RandInt(r, maxy-r), r, co)
	}
}

// randomDrawSlimLine is to draw slim line randomly
func (d *drawImage) randomDrawSlimLine(m canvas.Palette, num int, colorB []color.Color) {
	first := m.Bounds().Max.X / 10
	end := first * 9
	y := m.Bounds().Max.Y / 3
	for i := 0; i < num; i++ {
		point1 := image.Point{X: mRand.Intn(first), Y: mRand.Intn(y)}
		point2 := image.Point{X: mRand.Intn(first) + end, Y: mRand.Intn(y)}

		if i%2 == 0 {
			point1.Y = mRand.Intn(y) + y*2
			point2.Y = mRand.Intn(y)
		} else {
			point1.Y = mRand.Intn(y) + y*(i%2)
			point2.Y = mRand.Intn(y) + y*2
		}

		co := randgen.RandColor(colorB)
		//co.A = uint8(0xee)
		m.DrawBeeline(point1, point2, co)
	}
}

// DrawDotImage is to draw dot image
func (d *drawImage) DrawDotImage(dot *DrawDot, params *DrawImageParams) (canvas.NRGBA, *canvas.AreaRect, error) {
	cColor, err := helper.ParseHexColor(dot.Color)
	if err != nil {
		return nil, nil, err
	}
	cColor.A = helper.FormatAlpha(params.Alpha)

	var cImage image.Image
	if dot.DrawType == DrawTypeImage {
		cImage, err = d.DrawShapeImage(dot, cColor)
	} else {
		cImage, err = d.DrawStringImage(dot, cColor)
	}
	if err != nil {
		return nil, nil, err
	}

	shadowColorHex := shadowColor
	if params.ShadowColor != "" {
		shadowColorHex = params.ShadowColor
	}

	sColor, err := helper.ParseHexColor(shadowColorHex)
	if err != nil {
		return nil, nil, err
	}

	cvs := canvas.CreateNRGBACanvas(dot.Width+10, dot.Height+10, true)
	if params.ShowShadow {
		var shadowImg canvas.NRGBA
		if dot.DrawType == DrawTypeImage {
			shadowImg, err = d.DrawShapeImage(dot, sColor)
		} else {
			shadowImg, err = d.DrawStringImage(dot, sColor)
		}
		if err != nil {
			return nil, nil, err
		}

		pointX := params.ShadowPoint.X
		pointY := params.ShadowPoint.Y
		draw.Draw(cvs.Get(), shadowImg.Bounds(), shadowImg, image.Point{X: pointX, Y: pointY}, draw.Over)
	}
	draw.Draw(cvs.Get(), cImage.Bounds(), cImage, image.Point{}, draw.Over)
	cvs.Rotate(dot.Angle)

	ap := cvs.CalcMarginBlankArea()

	return cvs, ap, nil
}

// DrawStringImage is to draw string image
func (d *drawImage) DrawStringImage(dot *DrawDot, textColor color.Color) (canvas.NRGBA, error) {
	cvs := canvas.CreateNRGBACanvas(dot.Width+10, dot.Height+10, true)

	pt := freetype.Pt(12, dot.Height-5)
	if helper.IsChineseChar(dot.Text) {
		pt = freetype.Pt(10, dot.Height)
	}

	err := cvs.DrawString(&canvas.DrawStringParams{
		Color:   textColor,
		Size:    dot.Size,
		Width:   dot.Width,
		Height:  dot.Height,
		FontDPI: dot.FontDPI,
		Text:    dot.Text,
		Font:    dot.Font,
	}, pt)

	if err != nil {
		return nil, err
	}

	return cvs, nil
}

// DrawShapeImage is to draw shape image
func (d *drawImage) DrawShapeImage(dot *DrawDot, cColor color.Color) (canvas.NRGBA, error) {
	cr, cg, cb, ca := cColor.RGBA()

	var colorArr = []color.RGBA{
		{R: uint8(cr), G: uint8(cg), B: uint8(cb), A: uint8(ca)},
	}

	ncvs := canvas.CreateNRGBACanvas(dot.Width+10, dot.Height+10, true)
	var bounds image.Rectangle
	var img image.Image
	if dot.UseOriginalColor {
		cvs := canvas.CreateNRGBACanvas(dot.Width+10, dot.Height+10, true)
		draw.BiLinear.Scale(cvs.Get(), cvs.Bounds(), dot.Image, dot.Image.Bounds(), draw.Over, nil)
		bounds = cvs.Bounds()
		img = cvs.Get()
	} else {
		cvs := canvas.CreatePaletteCanvas(dot.Width+10, dot.Height+10, colorArr)
		draw.BiLinear.Scale(cvs.Get(), cvs.Bounds(), dot.Image, dot.Image.Bounds(), draw.Over, nil)
		bounds = cvs.Bounds()
		img = cvs.Get()
	}

	draw.Draw(ncvs.Get(), bounds, img, image.Point{}, draw.Over)

	return ncvs, nil
}

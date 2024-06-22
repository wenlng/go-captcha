/**
 * @Author Awen
 * @Date 2024/06/01
 * @Email wengaolng@gmail.com
 **/

package click

import (
	"fmt"
	"image"
	"image/color"
	"math"
	"math/rand"

	"github.com/wenlng/go-captcha/v2/base/helper"
	"github.com/wenlng/go-captcha/v2/base/imagedata"
	"github.com/wenlng/go-captcha/v2/base/logger"
	"github.com/wenlng/go-captcha/v2/base/option"
	"github.com/wenlng/go-captcha/v2/base/randgen"
	"github.com/wenlng/go-captcha/v2/base/random"
)

// Version # of captcha
const Version = "2.0.0"

// Captcha .
type Captcha interface {
	setOptions(opts ...Option)
	setResources(resources ...Resource)
	GetOptions() *Options
	Generate() (CaptchaData, error)
}

type Mode int

const (
	ModeText Mode = iota
	ModeShape
)

var _ Captcha = (*captcha)(nil)

// captcha .
type captcha struct {
	version   string
	logger    logger.Logger
	drawImage DrawImage
	opts      *Options
	resources *Resources
	mode      Mode
}

// newWithMode is to create a captcha
func newWithMode(mode Mode, opts ...Option) *captcha {
	capt := &captcha{
		version:   Version,
		logger:    logger.New(),
		drawImage: NewDrawImage(),
		mode:      mode,
		opts:      NewOptions(),
		resources: NewResources(),
	}

	defaultOptions()(capt.opts)
	defaultResource()(capt.resources)

	if mode == ModeShape {
		capt.opts.thumbBgDistort = option.DistortLevel1
		capt.opts.rangeSize = &option.RangeVal{Min: 24, Max: 30}
		capt.opts.rangeThumbSize = &option.RangeVal{Min: 14, Max: 20}
	}

	capt.setOptions(opts...)

	return capt
}

// setOptions is the set option
func (c *captcha) setOptions(opts ...Option) {
	for _, opt := range opts {
		opt(c.opts)
	}
}

// setResources is the set resource
func (c *captcha) setResources(resources ...Resource) {
	for _, resource := range resources {
		resource(c.resources)
	}
}

// GetOptions is to get options
func (c *captcha) GetOptions() *Options {
	return c.opts
}

// Generate is to generate the captcha data
func (c *captcha) Generate() (CaptchaData, error) {
	if c.mode == ModeShape {
		return c.generateWithShape()
	}

	return c.generateWithText()
}

// generateWithShape is to generate the graphical captcha data
func (c *captcha) generateWithShape() (CaptchaData, error) {
	if err := c.check(); err != nil {
		return nil, err
	}

	shapes, err := c.genShapes()
	if err != nil {
		return nil, err
	}

	var dots, thumbDots, verifyDots map[int]*Dot
	var verifyShapes []string
	var masterImage, thumbImage image.Image

	dots = c.genDots(c.opts.imageSize, c.opts.rangeSize, shapes, 10)
	verifyDots, verifyShapes = c.rangeCheckDots(dots)
	thumbDots = c.genDots(c.opts.thumbImageSize, c.opts.rangeThumbSize, verifyShapes, 0)

	masterImage, err = c.genMasterImage(c.opts.imageSize, dots)
	if err != nil {
		return nil, err
	}

	thumbImage, err = c.genThumbImage(c.opts.thumbImageSize, thumbDots)
	if err != nil {
		return nil, err
	}

	return &CaptData{
		dots:        verifyDots,
		masterImage: imagedata.NewJPEGImageData(masterImage),
		thumbImage:  imagedata.NewPNGImageData(thumbImage),
	}, nil
}

// generateWithText is to generate the text captcha data
func (c *captcha) generateWithText() (CaptchaData, error) {
	if err := c.check(); err != nil {
		return nil, err
	}

	chars, err := c.genChars()
	if err != nil {
		return nil, err
	}

	var dots, thumbDots, verifyDots map[int]*Dot
	var verifyShapes []string
	var masterImage, thumbImage image.Image

	dots = c.genDots(c.opts.imageSize, c.opts.rangeSize, chars, 10)
	verifyDots, verifyShapes = c.rangeCheckDots(dots)
	thumbDots = c.genDots(c.opts.thumbImageSize, c.opts.rangeThumbSize, verifyShapes, 0)

	masterImage, err = c.genMasterImage(c.opts.imageSize, dots)
	if err != nil {
		return nil, err
	}

	thumbImage, err = c.genThumbImage(c.opts.thumbImageSize, thumbDots)
	if err != nil {
		return nil, err
	}

	return &CaptData{
		dots:        verifyDots,
		masterImage: imagedata.NewJPEGImageData(masterImage),
		thumbImage:  imagedata.NewPNGImageData(thumbImage),
	}, nil
}

// genShapes is to generate an orderly shapes list
func (c *captcha) genShapes() ([]string, error) {
	length := random.RandInt(c.opts.rangeLen.Min, c.opts.rangeLen.Max)
	shapeNames := c.genRandShape(length)
	if len(shapeNames) == 0 {
		return []string{}, fmt.Errorf("click captcha err: %v", "no shapes generation")
	}
	return shapeNames, nil
}

// genChars is to generate an orderly character list
func (c *captcha) genChars() ([]string, error) {
	length := random.RandInt(c.opts.rangeLen.Min, c.opts.rangeLen.Max)
	chars := c.genRandChar(length)
	if len(chars) == 0 {
		return []string{}, fmt.Errorf("click captcha err: %v", "no character generation")
	}
	return chars, nil
}

// genDots is to generate an orderly dot list
func (c *captcha) genDots(imageSize *option.Size, size *option.RangeVal, values []string, padding int) map[int]*Dot {
	var dots = make(map[int]*Dot)
	width := imageSize.Width
	height := imageSize.Height
	if padding > 0 {
		width -= padding
		height -= padding
	}

	length := len(values)
	for i := 0; i < length; i++ {
		value := values[i]
		randAngle := c.randAngle()

		randColor := randgen.RandHexColor(c.opts.rangeColors)
		randColor2 := randgen.RandHexColor(c.opts.rangeThumbColors)

		randSize := random.RandInt(size.Min, size.Max)
		cHeight := randSize
		cWidth := randSize

		if c.mode == ModeText && helper.LenChineseChar(value) > 1 {
			cWidth = randSize * helper.LenChineseChar(value)

			if randAngle > 0 {
				surplus := cWidth - randSize
				ra := randAngle % 90
				pr := float64(surplus) / 90
				r := math.Max(float64(ra)*pr, 1)
				cHeight = cHeight + int(r)
				cWidth = cWidth + int(r)
			}
		}

		dy := 10
		w := width / length
		rd := math.Abs(float64(w) - float64(cWidth))
		xx := (i * w) + random.RandInt(0, int(math.Max(rd, 1)))
		yy := random.RandInt(dy, height+cHeight)

		x := int(math.Min(math.Max(float64(xx), float64(dy)), float64(width-dy-(padding*2))))
		y := int(math.Min(math.Max(float64(yy), float64(cHeight+dy)), float64(height+(cHeight/2)-(padding*2))))

		dots[i] = &Dot{
			Index:  i,
			X:      x,
			Y:      y - cHeight,
			Size:   randSize,
			Width:  cWidth,
			Height: cHeight,
			Angle:  randAngle,
			Color:  randColor,
			Color2: randColor2,
		}

		if c.mode == ModeShape {
			dots[i].Shape = value
		} else {
			dots[i].Text = fmt.Sprintf("%s", value)
		}
	}

	return dots
}

// check is to check the captcha parameter
func (c *captcha) check() error {
	if c.mode == ModeText {
		if len(c.resources.chars) < c.opts.rangeLen.Max {
			return fmt.Errorf("click captcha err: chars must be large than to %d", c.opts.rangeLen.Max)
		}
		return nil
	} else if c.mode == ModeShape {
		if len(c.resources.shapes) < c.opts.rangeLen.Max {
			return fmt.Errorf("click captcha err: shapes must be large than to %d", c.opts.rangeLen.Max)
		}

		for name, img := range c.resources.shapeMaps {
			if img == nil {
				return fmt.Errorf("click captcha err: [%s] shape must be is image type", name)
			}
		}

		return nil
	}

	if len(c.resources.rangBackgrounds) == 0 {
		return fmt.Errorf("click captcha err: no rang backgroun image")
	}

	return fmt.Errorf("click captcha err: %v", "mode not supported")
}

// rangeCheckDots is to generate random detection points
func (c *captcha) rangeCheckDots(dots map[int]*Dot) (map[int]*Dot, []string) {
	rs := random.Perm(len(dots))
	chkDots := make(map[int]*Dot)
	count := random.RandInt(c.opts.rangeVerifyLen.Min, c.opts.rangeVerifyLen.Max)
	var values []string
	for i, value := range rs {
		if i >= count {
			continue
		}
		dot := dots[value]
		dot.Index = i
		chkDots[i] = dot
		if c.mode == ModeShape {
			values = append(values, chkDots[i].Shape)
		} else {
			values = append(values, chkDots[i].Text)
		}
	}
	return chkDots, values
}

// genMasterImage is the master image of drawing captcha
func (c *captcha) genMasterImage(size *option.Size, dots map[int]*Dot) (image.Image, error) {
	var drawDots = make([]*DrawDot, 0, len(dots))

	for i := 0; i < len(dots); i++ {
		dot := dots[i]
		drawDot := &DrawDot{
			Dot:    dot,
			X:      dot.X,
			Y:      dot.Y,
			Width:  dot.Width,
			Height: dot.Height,
			Angle:  dot.Angle,
			Color:  dot.Color,
			Size:   dot.Size,
		}

		if c.mode == ModeShape {
			drawDot.DrawType = DrawTypeImage
			if img, ok := c.resources.shapeMaps[dot.Shape]; ok {
				drawDot.Image = img
			}
			drawDot.UseOriginalColor = c.opts.useShapeOriginalColor
		} else {
			drawDot.DrawType = DrawTypeString
			drawDot.Text = dot.Text
			drawDot.FontDPI = c.opts.fontDPI
			drawDot.Font = randgen.RandFont(c.resources.rangFonts)
		}

		drawDots = append(drawDots, drawDot)
	}

	return c.drawImage.DrawWithNRGBA(&DrawImageParams{
		Width:          size.Width,
		Height:         size.Height,
		Background:     randgen.RandImage(c.resources.rangBackgrounds),
		Alpha:          c.opts.imageAlpha,
		FontHinting:    c.opts.fontHinting,
		CaptchaDrawDot: drawDots,

		ShowShadow:  c.opts.displayShadow,
		ShadowColor: c.opts.shadowColor,
		ShadowPoint: c.opts.shadowPoint,
	})
}

// genThumbImage is the thumbnail image of drawing captcha
func (c *captcha) genThumbImage(size *option.Size, dots map[int]*Dot) (image.Image, error) {
	var drawDots = make([]*DrawDot, 0, len(dots))

	width := size.Width / len(dots)
	for i := 0; i < len(dots); i++ {
		dot := dots[i]
		length := 1
		if c.mode == ModeText {
			length = len(dot.Text)
		}

		dx := int(math.Max(float64(width*i+width/dot.Width), 8))
		dy := size.Height/2 + dot.Size/2 - rand.Intn(size.Height/16*length)

		drawDot := &DrawDot{
			Dot:    dot,
			X:      dx,
			Y:      dy,
			Angle:  dot.Angle,
			Color:  dot.Color2,
			Size:   dot.Size,
			Width:  dot.Width,
			Height: dot.Height,
		}

		if c.mode == ModeShape {
			drawDot.DrawType = DrawTypeImage
			if img, ok := c.resources.shapeMaps[dot.Shape]; ok {
				drawDot.Image = img
			}
			drawDot.UseOriginalColor = c.opts.useShapeOriginalColor
		} else {
			drawDot.DrawType = DrawTypeString
			drawDot.Text = dot.Text
			drawDot.FontDPI = c.opts.fontDPI
			drawDot.Font = randgen.RandFont(c.resources.rangFonts)
		}

		drawDots = append(drawDots, drawDot)
	}

	params := &DrawImageParams{
		Width:                 size.Width,
		Height:                size.Height,
		CaptchaDrawDot:        drawDots,
		BackgroundDistort:     c.randDistortWithLevel(c.opts.thumbBgDistort),
		BackgroundCirclesNum:  c.opts.thumbBgCirclesNum,
		BackgroundSlimLineNum: c.opts.thumbBgSlimLineNum,
		ThumbDisturbAlpha:     c.opts.thumbDisturbAlpha,
	}

	if len(c.resources.rangThumbBackgrounds) > 0 {
		params.Background = randgen.RandImage(c.resources.rangThumbBackgrounds)
	}

	var mTextColors []color.Color
	for _, cStr := range c.opts.rangeThumbColors {
		co, _ := helper.ParseHexColor(cStr)
		mTextColors = append(mTextColors, co)
	}

	var bgColors []color.Color
	for _, co := range c.opts.rangeThumbBgColors {
		rc, _ := helper.ParseHexColor(co)
		bgColors = append(bgColors, rc)
	}

	if c.opts.useShapeOriginalColor || c.opts.isThumbNonDeformAbility {
		return c.drawImage.DrawWithNRGBA2(params, mTextColors, bgColors)
	}
	return c.drawImage.DrawWithPalette(params, mTextColors, bgColors)
}

// genRandShape is to generate random shape array
func (c *captcha) genRandShape(length int) []string {
	var nameA []string
	for len(nameA) < length {
		img := randgen.RandString(c.resources.shapes)
		if !helper.InArrayWithStr(nameA, img) {
			nameA = append(nameA, img)
		}
	}

	return nameA
}

// genRandCharis is to generate random character array
func (c *captcha) genRandChar(length int) []string {
	var strA []string
	for len(strA) < length {
		char := randgen.RandString(c.resources.chars)
		if !helper.InArrayWithStr(strA, char) {
			strA = append(strA, char)
		}
	}

	return strA
}

// randDistortWithLevel is to generate random distort
func (c *captcha) randDistortWithLevel(level int) int {
	if level == 1 {
		return random.RandInt(240, 320)
	} else if level == 2 {
		return random.RandInt(180, 240)
	} else if level == 3 {
		return random.RandInt(120, 180)
	} else if level == 4 {
		return random.RandInt(100, 160)
	} else if level == 5 {
		return random.RandInt(80, 140)
	}
	return 0
}

// randAngle is to generate random angle
func (c *captcha) randAngle() int {
	angles := c.opts.rangeAnglePos

	index := helper.RandIndex(len(angles))
	if index < 0 {
		return 0
	}

	angle := angles[index]
	res := random.RandInt(angle.Min, angle.Max)

	return res
}

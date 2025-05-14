/**
 * @Author Awen
 * @Date 2024/06/01
 * @Email wengaolng@gmail.com
 **/

package click

import (
	"errors"
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

// Captcha defines the interface for captcha
type Captcha interface {
	setOptions(opts ...Option)
	setResources(resources ...Resource)
	GetOptions() *Options
	Generate() (CaptchaData, error)
}

// Mode defines the mode of the captcha
type Mode int

const (
	ModeText  Mode = iota // Text mode
	ModeShape             // Shape mode
)

var _ Captcha = (*captcha)(nil)

var (
	EmptyShapesErr          = errors.New("empty shapes")
	EmptyCharacterErr       = errors.New("empty character")
	CharRangeLenErr         = errors.New("character length must be large than to 'rangeLen.Max'")
	ShapesRangeLenErr       = errors.New("total number of shapes must be large than to 'rangeLen.Max'")
	ShapesTypeErr           = errors.New("shape must be is image type")
	EmptyBackgroundImageErr = errors.New("no background image")
	ModeSupportErr          = errors.New("mode not supported")
)

// captcha is the concrete implementation of the Captcha interface
type captcha struct {
	version   string
	logger    logger.Logger
	drawImage DrawImage
	opts      *Options
	resources *Resources
	mode      Mode
}

// newWithMode creates a captcha with the specified mode
// params:
//   - mode: Captcha mode
//   - opts: Optional options
//
// return: captcha instance
func newWithMode(mode Mode, opts ...Option) *captcha {
	capt := &captcha{
		logger:    logger.New(),
		drawImage: NewDrawImage(),
		mode:      mode,
		opts:      NewOptions(),
		resources: NewResources(),
	}

	defaultOptions()(capt.opts)
	defaultResource()(capt.resources)

	if mode == ModeShape {
		// Specific settings for shape mode
		capt.opts.thumbBgDistort = option.DistortLevel1
		capt.opts.rangeSize = &option.RangeVal{Min: 24, Max: 30}
		capt.opts.rangeThumbSize = &option.RangeVal{Min: 14, Max: 20}
	}

	capt.setOptions(opts...)

	return capt
}

// setOptions sets the captcha options
// opts: Options to set
func (c *captcha) setOptions(opts ...Option) {
	for _, opt := range opts {
		opt(c.opts)
	}
}

// setResources sets the captcha resources
// res: Resources to set
func (c *captcha) setResources(res ...Resource) {
	for _, resource := range res {
		resource(c.resources)
	}
}

// GetOptions gets the captcha options
// return: Captcha options
func (c *captcha) GetOptions() *Options {
	return c.opts
}

// Generate generates captcha data
// returns:
//   - CaptchaData: Generated captcha data
//   - error: Error information
func (c *captcha) Generate() (CaptchaData, error) {
	if c.mode == ModeShape {
		return c.generateWithShape()
	}

	return c.generateWithText()
}

// generateWithShape generates captcha data for shape mode
// returns:
//   - CaptchaData: Generated captcha data
//   - error: Error information
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

// generateWithText generates captcha data for text mode
// returns:
//   - CaptchaData: Generated captcha data
//   - error: Error information
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

// genShapes generates an ordered list of shapes
// returns:
//   - []string: List of shape names
//   - error: Error information
func (c *captcha) genShapes() ([]string, error) {
	length := random.RandInt(c.opts.rangeLen.Min, c.opts.rangeLen.Max)
	shapeNames := c.genRandShape(length)
	if len(shapeNames) == 0 {
		return []string{}, EmptyShapesErr
	}
	return shapeNames, nil
}

// genChars generates an ordered list of characters
// returns:
//   - []string: List of characters
//   - error: Error information
func (c *captcha) genChars() ([]string, error) {
	length := random.RandInt(c.opts.rangeLen.Min, c.opts.rangeLen.Max)
	chars := c.genRandChar(length)
	if len(chars) == 0 {
		return []string{}, EmptyCharacterErr
	}
	return chars, nil
}

// genDots generates an ordered list of dots
// params:
//   - imageSize: Image size
//   - size: Dot size range
//   - values: List of values (characters or shapes)
//   - padding: Padding
//
// return: Map of dot data
func (c *captcha) genDots(imageSize *option.Size, size *option.RangeVal, values []string, padding int) map[int]*Dot {
	var dots = make(map[int]*Dot, len(values))
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
			dots[i].Text = value
		}
	}

	return dots
}

// check checks the captcha parameters
// returns: Error information
func (c *captcha) check() error {
	if c.mode == ModeText {
		if len(c.resources.chars) < c.opts.rangeLen.Max {
			return CharRangeLenErr
		}
		return nil
	} else if c.mode == ModeShape {
		if len(c.resources.shapes) < c.opts.rangeLen.Max {
			return ShapesRangeLenErr
		}

		for _, img := range c.resources.shapeMaps {
			if img == nil {
				return ShapesTypeErr
			}
		}

		return nil
	}

	if len(c.resources.rangBackgrounds) == 0 {
		return EmptyBackgroundImageErr
	}

	return ModeSupportErr
}

// rangeCheckDots generates random verification dots
// params:
//   - dots: Map of dot data
//
// return:
//   - map[int]*Dot: Verification dot data
//   - []string: List of verification values
func (c *captcha) rangeCheckDots(dots map[int]*Dot) (map[int]*Dot, []string) {
	rs := random.Perm(len(dots))
	chkDots := make(map[int]*Dot)
	count := random.RandInt(c.opts.rangeVerifyLen.Min, c.opts.rangeVerifyLen.Max)
	var values []string
	for i, value := range rs {
		if !c.opts.disabledRangeVerifyLen && i >= count {
			break
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

// genMasterImage generates the main captcha image
// params:
//   - size: Image size
//   - dots: Map of dot data
//
// returns:
//   - image.Image: Generated image
//   - error: Error information
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

// genThumbImage generates the thumbnail image
// params:
//   - size: Image size
//   - dots: Map of dot data
//
// returns:
//   - image.Image: Generated thumbnail
//   - error: Error information
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

// genRandShape generates a random shape array
// params:
//   - length: Number of shapes
//
// returns: List of shape names
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

// genRandChar generates a random character array
// params:
//   - length: Number of characters
//
// return: List of characters
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

// randDistortWithLevel generates a random distortion level
// params:
//   - level: Distortion level
//
// return: Distortion value
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

// randAngle generates a random angle
// return: Angle value
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

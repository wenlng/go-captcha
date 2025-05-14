/**
 * @Author Awen
 * @Date 2024/06/01
 * @Email wengaolng@gmail.com
 **/

package rotate

import (
	"errors"
	"image"

	"github.com/wenlng/go-captcha/v2/base/helper"
	"github.com/wenlng/go-captcha/v2/base/imagedata"
	"github.com/wenlng/go-captcha/v2/base/logger"
	"github.com/wenlng/go-captcha/v2/base/randgen"
	"github.com/wenlng/go-captcha/v2/base/random"
)

// Captcha defines the interface for rotate CAPTCHA
type Captcha interface {
	setOptions(opts ...Option)
	setResources(resources ...Resource)
	GetOptions() *Options
	Generate() (CaptchaData, error)
}

var _ Captcha = (*captcha)(nil)

var (
	EmptyImageErr = errors.New("no image")
	ImageTypeErr  = errors.New("image must be is image.Image type")
)

// captcha is the concrete implementation of the Captcha interface
type captcha struct {
	version   string
	logger    logger.Logger
	drawImage DrawImage
	opts      *Options
	resources *Resources
}

// newRotate creates a new rotate CAPTCHA instance
// params:
//   - opts: Optional initial options
//
// return: Captcha interface instance
func newRotate(opts ...Option) Captcha {
	capt := &captcha{
		logger:    logger.New(),
		drawImage: NewDrawImage(),
		opts:      NewOptions(),
		resources: NewResources(),
	}

	defaultOptions()(capt.opts)
	defaultResource()(capt.resources)

	capt.setOptions(opts...)

	return capt
}

// setOptions sets the CAPTCHA options
// params:
//   - opts: Options to set
func (c *captcha) setOptions(opts ...Option) {
	for _, opt := range opts {
		opt(c.opts)
	}
}

// setResources sets the CAPTCHA resources
// params:
//   - resources: Resources to set
func (c *captcha) setResources(resources ...Resource) {
	for _, resource := range resources {
		resource(c.resources)
	}
}

// GetOptions gets the CAPTCHA options
// return: Pointer to options
func (c *captcha) GetOptions() *Options {
	return c.opts
}

// Generate generates rotate CAPTCHA data
// returns:
//   - CaptchaData: Generated CAPTCHA data
//   - error: Error information
func (c *captcha) Generate() (CaptchaData, error) {
	if err := c.check(); err != nil {
		return nil, err
	}

	thumbImageSquareSize := c.randThumbImageSquareSize()
	block := c.genBlock(c.opts.imageSquareSize, thumbImageSquareSize)
	var masterImage, tileImage image.Image
	var err error

	masterImage, err = c.genMasterImage(c.opts.imageSquareSize, block)
	if err != nil {
		return nil, err
	}

	tileImage, err = c.genThumbImage(masterImage, block, thumbImageSquareSize)
	if err != nil {
		return nil, err
	}

	return &CaptData{
		block:       block,
		masterImage: imagedata.NewPNGImageData(masterImage),
		thumbImage:  imagedata.NewPNGImageData(tileImage),
	}, nil
}

// genMasterImage generates the master CAPTCHA image
// params:
//   - size: Image size
//   - block: Block data
//
// returns:
//   - image.Image: Generated master image
//   - error: Error information
func (c *captcha) genMasterImage(size int, block *Block) (image.Image, error) {
	return c.drawImage.DrawWithNRGBA(&DrawImageParams{
		Rotate:     block.Angle,
		SquareSize: size,
		Background: randgen.RandImage(c.resources.rangImages),
	})
}

// genThumbImage generates a thumbnail image
// params:
//   - bgImage: Background image
//   - block: Block data
//   - thumbImageSquareSize: Thumbnail size
//
// returns:
//   - image.Image: Generated thumbnail image
//   - error: Error information
func (c *captcha) genThumbImage(bgImage image.Image, block *Block, thumbImageSquareSize int) (image.Image, error) {
	return c.drawImage.DrawWithCropCircle(&DrawCropCircleImageParams{
		Background:     bgImage,
		Alpha:          c.opts.thumbImageAlpha,
		SquareSize:     thumbImageSquareSize,
		Rotate:         block.Angle,
		ScaleRatioSize: (c.opts.imageSquareSize - thumbImageSquareSize) / 2,
	})
}

// randAngle generates a random angle
// returns: Random angle value
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

// randThumbImageSquareSize generates a random thumbnail size
// return: Random thumbnail size
func (c *captcha) randThumbImageSquareSize() int {
	size := c.opts.rangeThumbImageSquareSize

	index := helper.RandIndex(len(size))
	if index < 0 {
		return 0
	}

	return size[index]
}

// genBlock generates CAPTCHA block data
// params:
//   - imageSize: Main image size
//   - thumbImageSquareSize: Thumbnail size
//
// return: Pointer to block data
func (c *captcha) genBlock(imageSize int, thumbImageSquareSize int) *Block {
	var block = &Block{}
	thumbWidth := thumbImageSquareSize
	thumbHeight := thumbImageSquareSize

	block.Angle = c.randAngle()
	block.Width = thumbWidth
	block.Height = thumbHeight

	block.ParentWidth = imageSize
	block.ParentHeight = imageSize

	return block
}

// check checks the CAPTCHA parameters
// return: Error information
func (c *captcha) check() error {
	if len(c.resources.rangImages) == 0 {
		return EmptyImageErr
	}
	for _, img := range c.resources.rangImages {
		if img == nil {
			return ImageTypeErr
		}
	}
	return nil
}

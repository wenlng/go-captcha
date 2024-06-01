/**
 * @Author Awen
 * @Date 2024/05/01
 * @Email wengaolng@gmail.com
 **/

package rotate

import (
	"fmt"
	"image"

	"github.com/wenlng/go-captcha/v2/base/helper"
	"github.com/wenlng/go-captcha/v2/base/imagedata"
	"github.com/wenlng/go-captcha/v2/base/logger"
	"github.com/wenlng/go-captcha/v2/base/randgen"
	"github.com/wenlng/go-captcha/v2/base/random"
)

// Version # of captcha
const Version = "2.0.0"

// Captcha .
type Captcha interface {
	SetOptions(opts ...Option)
	GetOptions() *Options
	SetResources(resources ...Resource)
	Generate() (CaptchaData, error)
}

var _ Captcha = (*captcha)(nil)

// captcha .
type captcha struct {
	version   string
	logger    logger.Logger
	drawImage DrawImage
	opts      *Options
	resources *Resources
}

// GetOptions is to get options
func (c *captcha) GetOptions() *Options {
	return c.opts
}

// New .
func New(opts ...Option) Captcha {
	capt := &captcha{
		version:   Version,
		logger:    logger.New(),
		drawImage: NewDrawImage(),
		opts:      NewOptions(),
		resources: NewResources(),
	}

	defaultOptions()(capt.opts)
	defaultResource()(capt.resources)

	capt.SetOptions(opts...)

	return capt
}

// SetOptions is the set option
func (c *captcha) SetOptions(opts ...Option) {
	for _, opt := range opts {
		opt(c.opts)
	}
}

// SetResources is the set resource
func (c *captcha) SetResources(resources ...Resource) {
	for _, resource := range resources {
		resource(c.resources)
	}
}

// Generate is to generate the captcha data
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

// generateWithShape is to generate the master image
func (c *captcha) genMasterImage(size int, block *Block) (image.Image, error) {
	return c.drawImage.DrawWithNRGBA(&DrawImageParams{
		Rotate:     block.Angle,
		SquareSize: size,
		Background: randgen.RandImage(c.resources.rangImages),
	})
}

// genThumbImage is to generate a thumbnail image
func (c *captcha) genThumbImage(bgImage image.Image, block *Block, thumbImageSquareSize int) (image.Image, error) {
	return c.drawImage.DrawWithCropCircle(&DrawCropCircleImageParams{
		Background:     bgImage,
		Alpha:          c.opts.thumbImageAlpha,
		SquareSize:     thumbImageSquareSize,
		Rotate:         block.Angle,
		ScaleRatioSize: (c.opts.imageSquareSize - thumbImageSquareSize) / 2,
	})
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

// randAngle is to generate random size
func (c *captcha) randThumbImageSquareSize() int {
	size := c.opts.rangeThumbImageSquareSize

	index := helper.RandIndex(len(size))
	if index < 0 {
		return 0
	}

	return size[index]
}

// genBlock is to generate block
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

// check is to check the captcha parameter
func (c *captcha) check() error {
	if len(c.resources.rangImages) == 0 {
		return fmt.Errorf("rotate captcha err: no rang image")
	}
	for _, img := range c.resources.rangImages {
		if img == nil {
			return fmt.Errorf("rotate captcha err: image must be is image.Image type")
		}
	}
	return nil
}

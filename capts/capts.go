package capts

import (
	"github.com/wenlng/go-captcha/capts/click"
	"github.com/wenlng/go-captcha/capts/rotate"
	"github.com/wenlng/go-captcha/capts/slide"
)

// NewClick .
func NewClick(opts ...click.Option) click.Captcha {
	return click.New(opts...)
}

// NewClickWithShape .
func NewClickWithShape(opts ...click.Option) click.Captcha {
	return click.NewWithShape(opts...)
}

// NewSlide .
func NewSlide(opts ...slide.Option) slide.Captcha {
	return slide.New(opts...)
}

// NewSlideWithRegion .
func NewSlideWithRegion(opts ...slide.Option) slide.Captcha {
	return slide.NewWithRegion(opts...)
}

// NewRotate .
func NewRotate(opts ...rotate.Option) rotate.Captcha {
	return rotate.New(opts...)
}

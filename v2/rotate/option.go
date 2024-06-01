/**
 * @Author Awen
 * @Date 2024/05/01
 * @Email wengaolng@gmail.com
 **/

package rotate

import (
	"github.com/wenlng/go-captcha/v2/base/option"
)

type Options struct {
	imageSquareSize int
	rangeAnglePos   []*option.RangeVal

	rangeThumbImageSquareSize []int
	thumbImageAlpha           float32
}

// GetImageSize .
func (o *Options) GetImageSize() int {
	return o.imageSquareSize
}

// GetRangeAngle .
func (o *Options) GetRangeAngle() []*option.RangeVal {
	var rv = make([]*option.RangeVal, len(o.rangeAnglePos))
	for i := 0; i < len(o.rangeAnglePos); i++ {
		rv[i] = &option.RangeVal{
			Min: o.rangeAnglePos[i].Min,
			Max: o.rangeAnglePos[i].Max,
		}
	}
	return rv
}

// GetThumbImageAlpha .
func (o *Options) GetThumbImageAlpha() float32 {
	return o.thumbImageAlpha
}

// GetRangeThumbImageSquareSize .
func (o *Options) GetRangeThumbImageSquareSize() []int {
	return o.rangeThumbImageSquareSize
}

type Option func(*Options)

// NewOptions .
func NewOptions() *Options {
	return &Options{}
}

//‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾
// Image
//_______________________________________________________________________

// WithImageSquareSize .
func WithImageSquareSize(val int) Option {
	return func(opts *Options) {
		opts.imageSquareSize = val
	}
}

// WithRangeAnglePos .
func WithRangeAnglePos(vals []option.RangeVal) Option {
	return func(opts *Options) {
		var newVals = make([]*option.RangeVal, 0)
		for i := 0; i < len(vals); i++ {
			val := vals[i]
			newVals = append(newVals, &option.RangeVal{Min: val.Min, Max: val.Max})
		}
		opts.rangeAnglePos = newVals
	}
}

//‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾
// Thumb Image
//_______________________________________________________________________

// WithRangeThumbImageSquareSize .
func WithRangeThumbImageSquareSize(val []int) Option {
	return func(opts *Options) {
		opts.rangeThumbImageSquareSize = val
	}
}

// WithThumbImageAlpha .
func WithThumbImageAlpha(val float32) Option {
	return func(opts *Options) {
		opts.thumbImageAlpha = val
	}
}

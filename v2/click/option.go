/**
 * @Author Awen
 * @Date 2024/05/01
 * @Email wengaolng@gmail.com
 **/

package click

import (
	"github.com/wenlng/go-captcha/v2/base/logger"
	"github.com/wenlng/go-captcha/v2/base/option"
	"golang.org/x/image/font"
)

// Options .
type Options struct {
	fontDPI     int
	fontHinting font.Hinting

	imageSize     *option.Size
	rangeLen      *option.RangeVal
	rangeAnglePos []*option.RangeVal
	rangeSize     *option.RangeVal
	rangeColors   []string
	displayShadow bool
	shadowColor   string
	shadowPoint   *option.Point
	imageAlpha    float32

	thumbImageSize          *option.Size
	rangeVerifyLen          *option.RangeVal
	rangeThumbSize          *option.RangeVal
	rangeThumbColors        []string
	rangeThumbBgColors      []string
	thumbBgDistort          int
	thumbBgCirclesNum       int
	thumbBgSlimLineNum      int
	isThumbNonDeformAbility bool
	thumbDisturbAlpha       float32

	useShapeOriginalColor bool
}

// GetImageSize .
func (o *Options) GetImageSize() *option.Size {
	return &option.Size{
		Width:  o.imageSize.Width,
		Height: o.imageSize.Height,
	}
}

// GetRangeLen .
func (o *Options) GetRangeLen() *option.RangeVal {
	return &option.RangeVal{
		Min: o.rangeLen.Min,
		Max: o.rangeLen.Max,
	}
}

// GetRangeAnglePos .
func (o *Options) GetRangeAnglePos() []*option.RangeVal {
	var rv = make([]*option.RangeVal, len(o.rangeAnglePos))
	for i := 0; i < len(o.rangeAnglePos); i++ {
		rv[i] = &option.RangeVal{
			Min: o.rangeAnglePos[i].Min,
			Max: o.rangeAnglePos[i].Max,
		}
	}
	return rv
}

// GetRangeSize .
func (o *Options) GetRangeSize() *option.RangeVal {
	return &option.RangeVal{
		Min: o.rangeSize.Min,
		Max: o.rangeSize.Max,
	}
}

// GetRangeColors .
func (o *Options) GetRangeColors() []string {
	var rv = make([]string, 0, len(o.rangeColors))
	rv = append(rv, o.rangeColors...)
	return rv
}

// GetDisplayShadow .
func (o *Options) GetDisplayShadow() bool {
	return o.displayShadow
}

// GetShadowColor .
func (o *Options) GetShadowColor() string {
	return o.shadowColor
}

// GetShadowPoint .
func (o *Options) GetShadowPoint() *option.Point {
	return &option.Point{
		X: o.shadowPoint.X,
		Y: o.shadowPoint.Y,
	}
}

// GetImageAlpha .
func (o *Options) GetImageAlpha() float32 {
	return o.imageAlpha
}

// GetThumbImageSize .
func (o *Options) GetThumbImageSize() *option.Size {
	return &option.Size{
		Width:  o.thumbImageSize.Width,
		Height: o.thumbImageSize.Height,
	}
}

// GetRangeVerifyLen .
func (o *Options) GetRangeVerifyLen() *option.RangeVal {
	return &option.RangeVal{
		Min: o.rangeVerifyLen.Min,
		Max: o.rangeVerifyLen.Max,
	}
}

// GetRangeThumbSize .
func (o *Options) GetRangeThumbSize() *option.RangeVal {
	return &option.RangeVal{
		Min: o.rangeThumbSize.Min,
		Max: o.rangeThumbSize.Max,
	}
}

// GetRangeThumbColors .
func (o *Options) GetRangeThumbColors() []string {
	var rv = make([]string, 0, len(o.rangeThumbColors))
	rv = append(rv, o.rangeThumbColors...)
	return rv
}

// GetRangeThumbBgColors .
func (o *Options) GetRangeThumbBgColors() []string {
	var rv = make([]string, 0, len(o.rangeThumbBgColors))
	rv = append(rv, o.rangeThumbBgColors...)
	return rv
}

// GetThumbBgDistort .
func (o *Options) GetThumbBgDistort() int {
	return o.thumbBgDistort
}

// GetThumbBgCirclesNum .
func (o *Options) GetThumbBgCirclesNum() int {
	return o.thumbBgCirclesNum
}

// GetThumbBgSlimLineNum .
func (o *Options) GetThumbBgSlimLineNum() int {
	return o.thumbBgSlimLineNum
}

// GetUseShapeOriginalColor .
func (o *Options) GetUseShapeOriginalColor() bool {
	return o.useShapeOriginalColor
}

// GetIsThumbNonDeformAbility .
func (o *Options) GetIsThumbNonDeformAbility() bool {
	return o.isThumbNonDeformAbility
}

// GetThumbDisturbAlpha .
func (o *Options) GetThumbDisturbAlpha() float32 {
	return o.thumbDisturbAlpha
}

type Option func(*Options)

// NewOptions .
func NewOptions() *Options {
	return &Options{}
}

// WithFontHinting .
func WithFontHinting(val font.Hinting) Option {
	return func(opts *Options) {
		opts.fontHinting = val
	}
}

//‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾
// Image
//_______________________________________________________________________

// WithImageSize .
func WithImageSize(val option.Size) Option {
	return func(opts *Options) {
		opts.imageSize = &option.Size{Width: val.Width, Height: val.Height}
	}
}

// WithRangeLen .
func WithRangeLen(val option.RangeVal) Option {
	return func(opts *Options) {
		opts.rangeLen = &option.RangeVal{Min: val.Min, Max: val.Max}
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

// WithRangeSize .
func WithRangeSize(val option.RangeVal) Option {
	return func(opts *Options) {
		opts.rangeSize = &option.RangeVal{Min: val.Min, Max: val.Max}
	}
}

// WithRangeColors .
func WithRangeColors(colors []string) Option {
	return func(opts *Options) {
		if len(colors) > 255 {
			logger.New().Errorf("withRangeColors error: the max value of rangColors must be less than or equal to 255")
			return
		}

		opts.rangeColors = colors
	}
}

// WithDisplayShadow .
func WithDisplayShadow(val bool) Option {
	return func(opts *Options) {
		opts.displayShadow = val
	}
}

// WithShadowColor .
func WithShadowColor(val string) Option {
	return func(opts *Options) {
		opts.shadowColor = val
	}
}

// WithShadowPoint .
func WithShadowPoint(val option.Point) Option {
	return func(opts *Options) {
		opts.shadowPoint = &option.Point{X: val.X, Y: val.Y}
	}
}

// WithImageAlpha .
func WithImageAlpha(val float32) Option {
	return func(opts *Options) {
		if val > 1 {
			val = 1
		}
		opts.imageAlpha = val
	}
}

//‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾
// Thumb Image
//_______________________________________________________________________

// WithRangeThumbImageSize .
func WithRangeThumbImageSize(val option.Size) Option {
	return func(opts *Options) {
		opts.thumbImageSize = &option.Size{Width: val.Width, Height: val.Height}
	}
}

// WithRangeVerifyLen .
func WithRangeVerifyLen(val option.RangeVal) Option {
	return func(opts *Options) {
		if val.Max > opts.rangeLen.Min {
			logger.New().Errorf("withRangeVerifyLen error: the max value of rangeVerifyLen must be less than or equal to the min value of rangeLen")
			return
		}

		opts.rangeVerifyLen = &option.RangeVal{Min: val.Min, Max: val.Max}
	}
}

// WithRangeThumbSize .
func WithRangeThumbSize(val option.RangeVal) Option {
	return func(opts *Options) {
		opts.rangeThumbSize = &option.RangeVal{Min: val.Min, Max: val.Max}
	}
}

// WithRangeThumbColors .
func WithRangeThumbColors(val []string) Option {
	return func(opts *Options) {
		if len(val) > 255 {
			logger.New().Errorf("withRangeThumbColors error: the max value of rangeThumbColors must be less than or equal to 255")
			return
		}
		opts.rangeThumbColors = val
	}
}

// WithRangeThumbBgColors .
func WithRangeThumbBgColors(val []string) Option {
	return func(opts *Options) {
		if len(val) > 255 {
			logger.New().Errorf("withRangeThumbBgColors error: the max value of rangeThumbBgColors must be less than or equal to 255")
			return
		}

		opts.rangeThumbBgColors = val
	}
}

// WithRangeThumbBgDistort .
func WithRangeThumbBgDistort(val int) Option {
	return func(opts *Options) {
		if val >= option.DistortNone || val <= option.DistortLevel5 {
			opts.thumbBgDistort = val
		} else {
			opts.thumbBgDistort = option.DistortNone
		}
	}
}

// WithRangeThumbBgCirclesNum .
func WithRangeThumbBgCirclesNum(val int) Option {
	return func(opts *Options) {
		opts.thumbBgCirclesNum = val
	}
}

// WithRangeThumbBgSlimLineNum .
func WithRangeThumbBgSlimLineNum(val int) Option {
	return func(opts *Options) {
		opts.thumbBgSlimLineNum = val
	}
}

// WithUseShapeOriginalColor .
func WithUseShapeOriginalColor(val bool) Option {
	return func(opts *Options) {
		opts.useShapeOriginalColor = val
	}
}

// WithIsThumbNonDeformAbility .
func WithIsThumbNonDeformAbility(val bool) Option {
	return func(opts *Options) {
		opts.isThumbNonDeformAbility = val
	}
}

// WithThumbDisturbAlpha .
func WithThumbDisturbAlpha(val float32) Option {
	return func(opts *Options) {
		opts.thumbDisturbAlpha = val
	}
}

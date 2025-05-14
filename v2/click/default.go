/**
 * @Author Awen
 * @Date 2024/06/01
 * @Email wengaolng@gmail.com
 **/

package click

import (
	"github.com/wenlng/go-captcha/v2/base/option"
	"golang.org/x/image/font"
)

// Default color list
var colors = []string{
	"#fde98e",
	"#60c1ff",
	"#fcb08e",
	"#fb88ff",
	"#b4fed4",
	"#cbfaa9",
	"#78d6f8",
}

// Default thumbnail color list
var thumbColors = []string{
	"#1f55c4",
	"#780592",
	"#2f6b00",
	"#910000",
	"#864401",
	"#675901",
	"#016e5c",
}

// Default shadow color
var shadowColor = "#101010"

// Default character set
var defaultChars = []string{"我", "是", "行", "为", "式", "验", "证", "码", "的", "随", "机", "文", "本", "种", "子"}

// getDefaultColors gets the default color list
// return: List of colors
func getDefaultColors() []string {
	return colors
}

// getDefaultShadowColor gets the default shadow color
// return: Shadow color
func getDefaultShadowColor() string {
	return shadowColor
}

// getDefaultThumbColors gets the default thumbnail color list
// return: List of thumbnail colors
func getDefaultThumbColors() []string {
	return thumbColors
}

// getDefaultChars gets the default character set
// return: Character set
func getDefaultChars() []string {
	return defaultChars
}

// defaultOptions sets the default captcha options
// return: Option function
func defaultOptions() Option {
	return func(opts *Options) {
		opts.fontDPI = 72
		opts.fontHinting = font.HintingNone

		opts.rangeLen = &option.RangeVal{Min: 6, Max: 7}
		opts.rangeAnglePos = []*option.RangeVal{
			{Min: 20, Max: 35},
			{Min: 35, Max: 45},
			{Min: 45, Max: 60},
			{Min: 290, Max: 305},
			{Min: 305, Max: 325},
			{Min: 325, Max: 330},
		}
		opts.rangeSize = &option.RangeVal{Min: 26, Max: 32}
		opts.rangeColors = getDefaultColors()
		opts.displayShadow = true
		opts.shadowColor = getDefaultShadowColor()
		opts.shadowPoint = &option.Point{X: -1, Y: -1}
		opts.imageSize = &option.Size{Width: 300, Height: 220}
		opts.imageAlpha = 1

		opts.rangeVerifyLen = &option.RangeVal{Min: 2, Max: 4}
		opts.disabledRangeVerifyLen = false
		opts.thumbImageSize = &option.Size{Width: 150, Height: 40}
		opts.rangeThumbSize = &option.RangeVal{Min: 22, Max: 28}
		opts.rangeThumbColors = getDefaultThumbColors()
		opts.rangeThumbBgColors = getDefaultThumbColors()
		opts.thumbBgDistort = option.DistortLevel4
		opts.thumbBgCirclesNum = 24
		opts.thumbBgSlimLineNum = 2
		opts.isThumbNonDeformAbility = true
		opts.thumbDisturbAlpha = 1
	}
}

// defaultResource sets the default captcha resources
// return: Resource function
func defaultResource() Resource {
	return func(resources *Resources) {
		resources.chars = getDefaultChars()
	}
}

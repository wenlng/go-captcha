/**
 * @Author Awen
 * @Date 2024/05/01
 * @Email wengaolng@gmail.com
 **/

package rotate

import (
	"github.com/wenlng/go-captcha/capts/base/option"
)

// defaultOptions is the default configuration
func defaultOptions() Option {
	return func(opts *Options) {
		opts.imageSquareSize = 240
		opts.rangeAnglePos = []*option.RangeVal{
			{Min: 290, Max: 305},
			{Min: 305, Max: 325},
			{Min: 325, Max: 330},
		}

		opts.thumbImageAlpha = 1
		opts.rangeThumbImageSquareSize = []int{150, 160, 170, 180}
	}
}

// defaultResource is the default resource
func defaultResource() Resource {
	return func(resources *Resources) {
		// ...
	}
}

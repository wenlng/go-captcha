/**
 * @Author Awen
 * @Date 2024/05/01
 * @Email wengaolng@gmail.com
 **/

package click

import (
	"image"

	"github.com/golang/freetype/truetype"
	"github.com/wenlng/go-captcha/v2/base/helper"
	"github.com/wenlng/go-captcha/v2/base/logger"
)

type Resources struct {
	chars                []string
	shapeMaps            map[string]image.Image
	shapes               []string
	rangFonts            []*truetype.Font
	rangBackgrounds      []image.Image
	rangThumbBackgrounds []image.Image
}

// NewResources .
func NewResources() *Resources {
	return &Resources{}
}

type Resource func(*Resources)

// WithChars is set characters
func WithChars(chars []string) Resource {
	return func(resources *Resources) {
		if len(chars) > 0 {
			for _, char := range chars {
				if helper.IsChineseChar(char) {
					if helper.LenChineseChar(char) > 1 {
						logger.New().Errorf("WithChars error: the chinese char [%s] must be equal to 1", char)
						return
					}
				} else if helper.LenChineseChar(char) > 2 {
					logger.New().Errorf("WithChars error: the char [%s] must be less than or equal to 2", char)
					return
				}
			}
		}

		resources.chars = chars
	}
}

// WithShapes is set shape
func WithShapes(shapeMaps map[string]image.Image) Resource {
	return func(resources *Resources) {
		resources.shapeMaps = shapeMaps
		var shapes = make([]string, 0, len(shapeMaps))
		for name, _ := range shapeMaps {
			shapes = append(shapes, name)
		}
		resources.shapes = shapes
	}
}

// WithFonts is set font
func WithFonts(fonts []*truetype.Font) Resource {
	return func(resources *Resources) {
		resources.rangFonts = fonts
	}
}

// WithBackgrounds is set background image
func WithBackgrounds(images []image.Image) Resource {
	return func(resources *Resources) {
		resources.rangBackgrounds = images
	}
}

// WithThumbBackgrounds is set thumbnail background image
func WithThumbBackgrounds(images []image.Image) Resource {
	return func(resources *Resources) {
		resources.rangThumbBackgrounds = images
	}
}

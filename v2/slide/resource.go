/**
 * @Author Awen
 * @Date 2024/06/01
 * @Email wengaolng@gmail.com
 **/

package slide

import (
	"image"
)

// GraphImage defines the graph resources for the slide CAPTCHA
type GraphImage struct {
	OverlayImage image.Image
	ShadowImage  image.Image
	MaskImage    image.Image
}

// Resources defines the resource collection for the slide CAPTCHA
type Resources struct {
	rangBackgrounds []image.Image
	rangGraphImage  []*GraphImage
}

// NewResources creates a new Resources instance
// return: Pointer to a Resources instance
func NewResources() *Resources {
	return &Resources{}
}

type Resource func(*Resources)

// WithBackgrounds sets the background images
// params:
//   - images: List of background images
//
// return: Resource function
func WithBackgrounds(images []image.Image) Resource {
	return func(resources *Resources) {
		resources.rangBackgrounds = images
	}
}

// WithGraphImages sets the graph images
// params:
//   - images: List of graph images
//
// return: Resource function
func WithGraphImages(images []*GraphImage) Resource {
	return func(resources *Resources) {
		resources.rangGraphImage = images
	}
}

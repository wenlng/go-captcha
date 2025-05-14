/**
 * @Author Awen
 * @Date 2024/06/08
 * @Email wengaolng@gmail.com
 **/

package click

// Builder defines an interface for building captcha
type Builder interface {
	SetOptions(opts ...Option)
	SetResources(resources ...Resource)
	Clear()
	Make() Captcha
	MakeShape() Captcha
	// Deprecated: As of 2.1.0, it will be removed, please use [MakeShape].
	MakeWithShape() Captcha
}

// Ensure that the builder struct implements the Builder interface
var _ Builder = (*builder)(nil)

// builder is the concrete implementation of the Builder interface
type builder struct {
	opts      []Option
	resources []Resource
}

// NewBuilder creates a new Builder instance
// opts: Optional initial options
// return: Builder interface instance
func NewBuilder(opts ...Option) Builder {
	build := &builder{
		opts:      make([]Option, 0),
		resources: make([]Resource, 0),
	}

	if len(opts) > 0 {
		build.opts = opts
	}

	return build
}

// Clear clears all options and resources in the builder
func (b *builder) Clear() {
	b.opts = make([]Option, 0)
	b.resources = make([]Resource, 0)
}

// SetOptions sets the options for the captcha
// opts: Options to add
func (b *builder) SetOptions(opts ...Option) {
	if len(opts) > 0 {
		b.opts = append(b.opts, opts...)
	}
}

// SetResources sets the resources for the captcha
// resources: Resources to add
func (b *builder) SetResources(resources ...Resource) {
	if len(resources) > 0 {
		b.resources = append(b.resources, resources...)
	}
}

// Make generates a text-mode captcha
// return: Captcha instance
func (b *builder) Make() Captcha {
	// Create text-mode captcha
	capt := newWithMode(ModeText)

	capt.setOptions(b.opts...)
	capt.setResources(b.resources...)
	return capt
}

// MakeShape generates a shape-mode captcha
// return: Captcha instance
func (b *builder) MakeShape() Captcha {
	capt := newWithMode(ModeShape)
	capt.setOptions(b.opts...)
	capt.setResources(b.resources...)
	return capt
}

// MakeWithShape generates a shape-mode captcha (deprecated)
// return: Captcha instance
func (b *builder) MakeWithShape() Captcha {
	capt := newWithMode(ModeShape)
	capt.setOptions(b.opts...)
	capt.setResources(b.resources...)
	return capt
}

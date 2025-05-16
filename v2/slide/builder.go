/**
 * @Author Awen
 * @Date 2024/06/08
 * @Email wengaolng@gmail.com
 **/

package slide

// Builder defines the interface for building slide CAPTCHAs
type Builder interface {
	SetOptions(opts ...Option)
	SetResources(resources ...Resource)
	Clear()
	Make() Captcha
	MakeDragDrop() Captcha
	// Deprecated: As of 2.1.0, it will be removed, please use [MakeDrag].
	MakeWithRegion() Captcha
}

var _ Builder = (*builder)(nil)

// builder is the concrete implementation of the Builder interface
type builder struct {
	opts      []Option
	resources []Resource
}

// NewBuilder creates a new Builder instance
// params:
//   - opts: Optional initial options
//
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

// SetOptions sets the CAPTCHA options
// params:
//   - opts: Options to add
func (b *builder) SetOptions(opts ...Option) {
	if len(opts) > 0 {
		b.opts = append(b.opts, opts...)
	}
}

// SetResources sets the CAPTCHA resources
// params:
//   - resources: Resources to add
func (b *builder) SetResources(resources ...Resource) {
	if len(resources) > 0 {
		b.resources = append(b.resources, resources...)
	}
}

// Make generates a slide CAPTCHA in basic mode
// params: Captcha interface instance
func (b *builder) Make() Captcha {
	capt := newWithMode(ModeBasic)
	capt.setOptions(b.opts...)
	capt.setResources(b.resources...)
	return capt
}

// MakeWithRegion generates a slide CAPTCHA in region mode (deprecated)
// return: Captcha interface instance
func (b *builder) MakeWithRegion() Captcha {
	capt := newWithMode(ModeDrag)
	capt.setOptions(b.opts...)
	capt.setResources(b.resources...)
	return capt
}

// MakeDragDrop generates a slide CAPTCHA in drag mode
// return: Captcha interface instance
func (b *builder) MakeDragDrop() Captcha {
	capt := newWithMode(ModeDrag)
	capt.setOptions(b.opts...)
	capt.setResources(b.resources...)
	return capt
}

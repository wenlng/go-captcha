/**
 * @Author Awen
 * @Date 2024/06/08
 * @Email wengaolng@gmail.com
 **/

package rotate

// Builder defines the interface for building rotate CAPTCHAs
type Builder interface {
	SetOptions(opts ...Option)
	SetResources(resources ...Resource)
	Clear()
	Make() Captcha
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

// Make generates a rotate CAPTCHA
// return: Captcha interface instance
func (b *builder) Make() Captcha {
	capt := newRotate()
	capt.setOptions(b.opts...)
	capt.setResources(b.resources...)
	return capt
}

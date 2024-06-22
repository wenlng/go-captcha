package click

// Builder .
type Builder interface {
	SetOptions(opts ...Option)
	SetResources(resources ...Resource)
	Clear()
	Make() Captcha
	MakeWithShape() Captcha
}

var _ Builder = (*builder)(nil)

// builder .
type builder struct {
	opts      []Option
	resources []Resource
}

// NewBuilder .
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

func (b *builder) Clear() {
	b.opts = make([]Option, 0)
	b.resources = make([]Resource, 0)
}

// SetOptions is the set option
func (b *builder) SetOptions(opts ...Option) {
	if len(opts) > 0 {
		b.opts = append(b.opts, opts...)
	}
}

// SetResources is the set resource
func (b *builder) SetResources(resources ...Resource) {
	if len(resources) > 0 {
		b.resources = append(b.resources, resources...)
	}
}

// Make .
func (b *builder) Make() Captcha {
	capt := newWithMode(ModeText)
	capt.setOptions(b.opts...)
	capt.setResources(b.resources...)
	return capt
}

// MakeWithShape .
func (b *builder) MakeWithShape() Captcha {
	capt := newWithMode(ModeShape)
	capt.setOptions(b.opts...)
	capt.setResources(b.resources...)
	return capt
}

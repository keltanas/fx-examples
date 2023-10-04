package baz

type FooDoer interface {
	DoFoo()
}

type Baz struct {
	foo FooDoer
}

func New(foo FooDoer) *Baz {
	return &Baz{
		foo: foo,
	}
}

func (b *Baz) DoBaz() {
	b.foo.DoFoo()
}

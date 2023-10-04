package bar

type BazDoer interface {
	DoBaz()
}

type Bar struct {
	baz BazDoer
}

func New(baz BazDoer) *Bar {
	return &Bar{
		baz: baz,
	}
}

func (b *Bar) Do() {
	b.baz.DoBaz()
}

package foo

import (
	"go.uber.org/zap"
)

type Foo struct {
	log *zap.Logger
}

func New(log *zap.Logger) *Foo {
	return &Foo{
		log: log,
	}
}

func (f *Foo) DoFoo() {
	f.log.Info("DoFoo")
}

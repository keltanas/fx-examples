package main

import (
	"context"
	"log"
	"time"

	"go.uber.org/fx"
	"go.uber.org/zap"

	"github.com/keltanas/fx-examples/internal/pkg/bar"
	"github.com/keltanas/fx-examples/internal/pkg/baz"
	"github.com/keltanas/fx-examples/internal/pkg/foo"
)

// decision 1
func main() {
	var b *bar.Bar
	app := fx.New(
		fx.Provide(zap.NewDevelopment),
		fx.Provide(foo.New),
		fx.Provide(
			func(f *foo.Foo) *baz.Baz {
				return baz.New(f)
			},
		),
		fx.Provide(
			func(b *baz.Baz) *bar.Bar {
				return bar.New(b)
			},
		),
		fx.Populate(&b),
	)

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	if err := app.Start(ctx); err != nil {
		log.Fatal(err)
	}

	defer func() {
		if err := app.Stop(ctx); err != nil {
			log.Fatal(err)
		}
	}()

	b.Do()
}

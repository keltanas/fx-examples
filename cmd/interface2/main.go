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

// decision 2
func main() {
	var b *bar.Bar
	app := fx.New(
		fx.Provide(zap.NewDevelopment),
		fx.Provide(foo.New),
		fx.Provide(fx.Annotate(baz.New, fx.From(new(*foo.Foo)))),
		fx.Provide(fx.Annotate(bar.New, fx.From(new(*baz.Baz)))),
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

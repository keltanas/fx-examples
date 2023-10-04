package main

import (
	"context"
	"time"

	"go.uber.org/fx"
	"go.uber.org/zap"

	"github.com/keltanas/fx-examples/internal/pkg/bar"
	"github.com/keltanas/fx-examples/internal/pkg/baz"
	"github.com/keltanas/fx-examples/internal/pkg/foo"
)

// It will panic
func main() {
	var b *bar.Bar
	app := fx.New(
		fx.Provide(zap.NewDevelopment),
		fx.Provide(foo.New),
		fx.Provide(baz.New),
		fx.Provide(bar.New),
		fx.Populate(&b),
	)

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	if err := app.Start(ctx); err != nil {
		panic(err)
	}

	defer func() {
		if err := app.Stop(ctx); err != nil {
			panic(err)
		}
	}()

	b.Do()
}

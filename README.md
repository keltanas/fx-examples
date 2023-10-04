# fx-examples
User Fx Examples

## Interfaces

### Issue with error

```go
_ = fx.New(
    fx.Provide(zap.NewDevelopment),
    fx.Provide(foo.New),
    fx.Provide(baz.New),
    fx.Provide(bar.New),
    fx.Populate(&b),
)
```

```shell
go run cmd/interface0/main.go
```

### Decision 1

```go
_ = fx.New(
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
```

```shell
go run cmd/interface1/main.go

[Fx] PROVIDE    *zap.Logger <= go.uber.org/zap.NewDevelopment()
[Fx] PROVIDE    *foo.Foo <= github.com/keltanas/fx-examples/internal/pkg/foo.New()
[Fx] PROVIDE    *baz.Baz <= main.main.func1()
[Fx] PROVIDE    *bar.Bar <= main.main.func2()
[Fx] PROVIDE    fx.Lifecycle <= go.uber.org/fx.New.func1()
[Fx] PROVIDE    fx.Shutdowner <= go.uber.org/fx.(*App).shutdowner-fm()
[Fx] PROVIDE    fx.DotGraph <= go.uber.org/fx.(*App).dotGraph-fm()
[Fx] INVOKE             reflect.makeFuncStub()
[Fx] RUN        provide: go.uber.org/zap.NewDevelopment()
[Fx] RUN        provide: github.com/keltanas/fx-examples/internal/pkg/foo.New()
[Fx] RUN        provide: main.main.func1()
[Fx] RUN        provide: main.main.func2()
[Fx] RUNNING
2023-10-04T17:50:12.977+0300    INFO    foo/foo.go:18   DoFoo
```

### Decision 2

> It's more preferred

```go
_ = fx.New(
    fx.Provide(zap.NewDevelopment),
    fx.Provide(foo.New),
    fx.Provide(fx.Annotate(baz.New, fx.From(new(*foo.Foo)))),
    fx.Provide(fx.Annotate(bar.New, fx.From(new(*baz.Baz)))),
    fx.Populate(&b),
)
```

```shell
go run cmd/interface2/main.go

[Fx] PROVIDE    *zap.Logger <= go.uber.org/zap.NewDevelopment()
[Fx] PROVIDE    *foo.Foo <= github.com/keltanas/fx-examples/internal/pkg/foo.New()
[Fx] PROVIDE    *baz.Baz <= fx.Annotate(github.com/keltanas/fx-examples/internal/pkg/baz.New(), fx.From([*foo.Foo])
[Fx] PROVIDE    *bar.Bar <= fx.Annotate(github.com/keltanas/fx-examples/internal/pkg/bar.New(), fx.From([*baz.Baz])
[Fx] PROVIDE    fx.Lifecycle <= go.uber.org/fx.New.func1()
[Fx] PROVIDE    fx.Shutdowner <= go.uber.org/fx.(*App).shutdowner-fm()
[Fx] PROVIDE    fx.DotGraph <= go.uber.org/fx.(*App).dotGraph-fm()
[Fx] INVOKE             reflect.makeFuncStub()
[Fx] RUN        provide: go.uber.org/zap.NewDevelopment()
[Fx] RUN        provide: github.com/keltanas/fx-examples/internal/pkg/foo.New()
[Fx] RUN        provide: fx.Annotate(github.com/keltanas/fx-examples/internal/pkg/baz.New(), fx.From([*foo.Foo])
[Fx] RUN        provide: fx.Annotate(github.com/keltanas/fx-examples/internal/pkg/bar.New(), fx.From([*baz.Baz])
[Fx] RUNNING
2023-10-04T17:50:27.753+0300    INFO    foo/foo.go:18   DoFoo
```

### Decision 3

```go
_ = fx.New(
    fx.Provide(zap.NewDevelopment),
    fx.Provide(fx.Annotate(foo.New, fx.As(new(baz.FooDoer)))),
    fx.Provide(fx.Annotate(baz.New, fx.As(new(bar.BazDoer)))),
    fx.Provide(bar.New),
    fx.Populate(&b),
)
```

```shell
go run cmd/interface3/main.go

[Fx] PROVIDE    *zap.Logger <= go.uber.org/zap.NewDevelopment()
[Fx] PROVIDE    baz.FooDoer <= fx.Annotate(github.com/keltanas/fx-examples/internal/pkg/foo.New(), fx.As([[baz.FooDoer]])
[Fx] PROVIDE    bar.BazDoer <= fx.Annotate(github.com/keltanas/fx-examples/internal/pkg/baz.New(), fx.As([[bar.BazDoer]])
[Fx] PROVIDE    *bar.Bar <= github.com/keltanas/fx-examples/internal/pkg/bar.New()
[Fx] PROVIDE    fx.Lifecycle <= go.uber.org/fx.New.func1()
[Fx] PROVIDE    fx.Shutdowner <= go.uber.org/fx.(*App).shutdowner-fm()
[Fx] PROVIDE    fx.DotGraph <= go.uber.org/fx.(*App).dotGraph-fm()
[Fx] INVOKE             reflect.makeFuncStub()
[Fx] RUN        provide: go.uber.org/zap.NewDevelopment()
[Fx] RUN        provide: fx.Annotate(github.com/keltanas/fx-examples/internal/pkg/foo.New(), fx.As([[baz.FooDoer]])
[Fx] RUN        provide: fx.Annotate(github.com/keltanas/fx-examples/internal/pkg/baz.New(), fx.As([[bar.BazDoer]])
[Fx] RUN        provide: github.com/keltanas/fx-examples/internal/pkg/bar.New()
[Fx] RUNNING
2023-10-04T17:50:37.224+0300    INFO    foo/foo.go:18   DoFoo
```

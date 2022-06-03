# fp

[![Go Reference](https://pkg.go.dev/badge/github.com/jaz303/fp.svg)](https://pkg.go.dev/github.com/jaz303/fp)

`fp` is a small utility library of functional programming primitives for Go.

If was originally written to get my head round 1.18's new generics features but it's grown to become pretty useful and has 100% test coverage so why not release it...

  - `Foreach()`, `Map()`, `Reduce()`
  - `All()`, `Any()`, `None()`
  - `First()`, `FirstIndexOf()`, `Count()`, `Filter()`, `Partition()`
  - `Keys()`, `Values()`
  - `Assoc()`
  - `Compose()`
  - `Repeat()`

Many of the functions an `-Index()` suffixed variant; the only difference with these is that their callbacks receive the current iteration index as an additional final parameter.
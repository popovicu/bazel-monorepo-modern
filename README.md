# bazel-monorepo-modern

The example project for
[Build like a Google engineer: modern Bazel with modules, Go and Typst](https://popovicu.com/posts/build-like-a-google-engineer-modern-bazel/).

One command fetches a Go SDK, a Typst compiler, a font and a Go library,
renders three PNG cards with Typst, builds a gallery page with a Go tool
written in this repository, and embeds all of it into a Go web server:

```bash
bazel run //server
```

Then open http://localhost:3333.

Other things to try:

```bash
bazel test //...
bazel build //server --platforms=@rules_go//go/toolchain:linux_amd64
bazel query 'rdeps(//..., //cards:card.typ)'
bazel aquery //server:index
```

Only Bazel (or [Bazelisk](https://github.com/bazelbuild/bazelisk)) needs to
be installed. `.bazelversion` pins Bazel itself.

# go-binary-size
This repo is companion to an article published here (TODO add link).

Contents :
* `app`: a Go module intended as an "application" (vs library) with 4 packages compilable as single binaries
  * `a`, `a2` and `b` depend exclusively on matching packages from `mod1`
  * `control` depends only on stdlib's `fmt` and is used as a control for what a small Go binary can weigh
* `mod1`: a Go module intended to represent a library that `app` would include. It contains 3 packages:
  * `a`: depends on a large external dependency (`github.com/gin-gonic/gin`) in the "normal" files and also depends on some other external dependencies exclusively used in the test files, whether in the same package or in the `a_test` package.
  * `a2`: same as `a` but doesn't contain test files, to show that test dependencies don't bloat the built binary
  * `b`: depends only on `fmt`

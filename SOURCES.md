# SOURCES.md

Canonical references for this repo. Answers should come from here (or from the linked
docs for a specific third-party package), not from memory or a random blog post.

Agents: read this file only when you need a link. The sourcing *rules* live in
[AGENTS.md](./AGENTS.md#sources) — this is just the index.

## Language

| What | Where |
| --- | --- |
| Spec (authoritative on syntax + semantics) | https://go.dev/ref/spec |
| Effective Go | https://go.dev/doc/effective_go |
| Go Code Review Comments (idiom/style) | https://go.dev/wiki/CodeReviewComments |
| Go Doc Comments (comment conventions) | https://go.dev/doc/comment |
| Memory model (happens-before, races) | https://go.dev/ref/mem |
| FAQ ("why does Go do X") | https://go.dev/doc/faq |
| Go Proverbs | https://go-proverbs.io/ |
| Official blog | https://go.dev/blog/ |
| Release notes (what changed in which version) | https://go.dev/doc/devel/release |

## Standard library

| What | Where |
| --- | --- |
| Package reference | https://pkg.go.dev/ |
| Full stdlib index | https://pkg.go.dev/std |
| Source (browsable) | https://cs.opensource.google/go/go/+/refs/heads/master:src/ |
| Source (GitHub mirror) | https://github.com/golang/go/tree/master/src |

Reading the source is the tiebreaker when the docs are ambiguous.

## Tooling and modules

| What | Where |
| --- | --- |
| `go` command reference | https://pkg.go.dev/cmd/go |
| Modules reference | https://go.dev/ref/mod |
| Workspaces (`go.work`) | https://go.dev/ref/mod#workspaces |
| Workspace tutorial | https://go.dev/doc/tutorial/workspaces |
| Organizing a module (layout) | https://go.dev/doc/modules/layout |
| Managing dependencies | https://go.dev/doc/modules/managing-dependencies |
| `go vet` | https://pkg.go.dev/cmd/vet |
| `gofmt` | https://pkg.go.dev/cmd/gofmt |
| `testing` package | https://pkg.go.dev/testing |
| Race detector | https://go.dev/doc/articles/race_detector |

## Topic deep-dives (official)

| Topic | Where |
| --- | --- |
| Errors + wrapping (`%w`, `errors.Is/As`) | https://go.dev/blog/go1.13-errors |
| Error handling basics | https://go.dev/blog/error-handling-and-go |
| `context` | https://go.dev/blog/context |
| Concurrency pipelines / cancellation | https://go.dev/blog/pipelines |
| Slices internals (append, aliasing) | https://go.dev/blog/slices-intro |
| Strings, bytes, runes, UTF-8 | https://go.dev/blog/strings |
| `defer`, `panic`, `recover` | https://go.dev/blog/defer-panic-and-recover |
| Table-driven tests | https://go.dev/wiki/TableDrivenTests |

## Packages these exercises lean on

Prefix each with `https://pkg.go.dev/`:

`encoding/csv` · `flag` · `time` · `bufio` · `os` · `io` · `strings` · `strconv` ·
`errors` · `fmt` · `sort` · `math/rand` · `sync` · `context` · `net/http` ·
`html/template` · `encoding/json` · `testing`

Outside stdlib but official-adjacent (`golang.org/x`): `golang.org/x/net/html` for the
link-parser and sitemap exercises.

## The exercises themselves

| What | Where |
| --- | --- |
| Gophercises | https://gophercises.com/ |

## Third-party packages

No index here on purpose. Use the maintainer's own README / docs site / godoc on
pkg.go.dev — not a tutorial written about it by someone else.

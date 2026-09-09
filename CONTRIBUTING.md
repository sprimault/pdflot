# Contributing

French version: `CONTRIBUTING.fr.md`, which is authoritative.

## How this repository is written

The code is written with the assistance of a language model, under review. The
resulting project rules live in `docs/` and apply to every contributor, human or
otherwise.

## Branch and publication

One branch per batch, named `<type>/<subject>`. One commit. Merged through a
squashed pull request.

A batch touching `internal/sandbox` or the manifest schema ships alone: those
are the two places where a regression does not show up in the diff.

## Before pushing

1. `make lint` passes — `gofmt -l .` returns nothing, `go vet ./...` is clean.
2. `make test` passes.
3. `make sec` and `make vulncheck` pass.
4. `make rendu-verif` passes, inside the reference image.
5. Documentation touched by the batch ships in the same commit.

The first four go through `make` rather than by hand: it is the same definition
continuous integration runs, and two lists kept in parallel drift apart without
saying so.

## Messages

Conventional types untranslated. Scope follows the package. Bilingual messages,
halves separated by `***`, four lines of body per language at most.

## Test documents

No excerpt of a rendered document in a message, pull request or issue. Refer to
a document by its identifier.

## Dependencies

Standard library by default. Every new dependency is justified in the pull
request, with its license. No AGPL dependency: it would contaminate users of the
service.

## Security

See `docs/securite.md` for the threat model and `SECURITY.md` for reporting.

A contribution adding an option that relaxes a sandbox guarantee is refused on
principle.

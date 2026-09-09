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

1. `make controles` passes — formatting, `go vet`, `go.sum` checksums, `gosec`,
   `govulncheck` and tests.
2. `make rendu-verif` passes, inside the reference image.
3. Documentation touched by the batch ships in the same commit.

The first runs on its own: `.githooks/pre-push` calls it before every push, once
the hook is enabled with `git config core.hooksPath .githooks` — once per clone.
`git push --no-verify` bypasses it, for pushing a work branch known to be broken.

One target, called by the hook and by continuous integration alike. Not a
convenience: as long as the checks were listed on both sides, the two lists
drifted apart silently, and the job refused formatting the local check accepted.

CodeQL is the only check not replayed locally, its analysis requiring a bundle of
close to a gigabyte. It runs on every proposal on GitHub, and its absence from
`make controles` is deliberate.

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

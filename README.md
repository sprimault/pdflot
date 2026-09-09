# pdflot

Batch HTML to PDF conversion, with grouping rules, duplex padding and packaged
delivery.

You submit a batch — an archive of HTML files, their assets and an optional
manifest. pdflot renders the documents with Chromium, assembles them according
to the requested rules, and returns an archive containing the PDFs, a control
sheet and an execution report. The target scale is 6000 documents per batch.

French version: `README.fr.md`.

## Warning

pdflot runs third-party HTML in a browser engine. All incoming HTML is treated
as hostile: no network access, no filesystem access outside the working
directory, assets resolved within the batch only. These guarantees cannot be
disabled through configuration.

Do not expose the service on a public network without an authenticating gateway
in front of it.

## What pdflot is not

Not a single-document HTML to PDF converter — Gotenberg already fills that role
well. The value here is the orchestration layer above it: grouping, padding,
packaging, multi-tenant scheduling.

Not a wkhtmltopdf fork either. Only the option vocabulary is carried over, so
that migration is immediate.

## Status

Early stage. See `ROADMAP.md`.

## License

Apache 2.0.

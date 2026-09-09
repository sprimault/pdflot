# Security

French version: `SECURITY.fr.md`, which is authoritative.

## Reporting a vulnerability

Through the repository's **Security** tab, **Report a vulnerability**. The
report stays private until a fix is published.

No public issue: a sandbox escape described in the open before a fix amounts to
instructions.

The project has no dedicated team. An answer comes as soon as possible, with no
guaranteed delay.

## Affected versions

No version has been released. Only the `master` branch is concerned.

## What counts as a vulnerability

The full threat model is in `docs/securite.md`. The starting assumption is that
submitted HTML is hostile, including from an authenticated tenant: any escape
from the sandbox guarantees counts.

- Outbound network access from a rendered document.
- Reading or writing outside the working directory of the job at hand.
- Resolving a resource outside the batch.
- Crossing tenant isolation, whether disk or concurrency.
- Document content or personal data surfacing in the logs.

## What does not

`docs/securite.md` lists what the project does not guarantee. A report on those
is handled as a feature request, not as a vulnerability:

- Isolation between documents of a single batch. The batch is the unit of trust.
- The archive password seen as protecting one recipient from another: it
  protects transport, and per-PDF encryption is what protects the recipient.
- Denial of service through a pathological document, beyond the configured
  quotas and deadlines.

A contribution adding an option that relaxes a sandbox guarantee is refused on
principle, whatever its quality otherwise.

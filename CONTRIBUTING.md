# Contributing to betriebsrat

German labour law CLI and Claude skill — grounded in gesetze-im-internet.de and the full BetrVG. Contributions welcome.

## Before you start

- Check [open issues](https://github.com/googlarz/betriebsrat/issues) and [discussions](https://github.com/googlarz/betriebsrat/discussions)
- For new legal content (cases, paragraphs, topic expansions), open an issue first — accuracy matters more than coverage

## Setup

```bash
git clone https://github.com/googlarz/betriebsrat.git
cd betriebsrat
make build
```

Requires Go 1.21+.

## Development

```bash
make build      # compile
make test       # run tests
make lint       # golangci-lint
```

## What to contribute

- **Bug fixes** — CLI output, parsing, command behaviour
- **Legal accuracy** — corrections to BetrVG citations, case references, or advice text (cite your source)
- **New commands** — covering BetrVG paragraphs not yet implemented
- **Language** — German/English phrasing improvements
- **Case citations** — verified BAG cases with working dejure.org or bundesarbeitsgericht.de links only

## Legal content standard

All legal content must be grounded in primary sources:
- Statute text: [gesetze-im-internet.de](https://www.gesetze-im-internet.de)
- Case law: [bundesarbeitsgericht.de](https://www.bundesarbeitsgericht.de) or [dejure.org](https://dejure.org) (verify link resolves)

Do not add case citations you cannot verify. Unverifiable citations are worse than no citations.

## Submitting a PR

1. Fork → branch from `main`
2. `make build && make test && make lint` must pass
3. Include source links for any new legal content
4. English PR description is fine; German welcome too

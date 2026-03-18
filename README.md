# docs-code

[![GoDoc](https://godoc.org/github.com/sshaplygin/docs-code?status.svg)](https://godoc.org/github.com/sshaplygin/docs-code)
[![Go Coverage](https://github.com/sshaplygin/docs-code/wiki/coverage.svg)](https://raw.githack.com/wiki/sshaplygin/docs-code/coverage.html)
[![License](https://img.shields.io/badge/license-MIT-blue.svg)](https://github.com/sshaplygin/docs-code/blob/master/LICENSE)
[![Go Report Card](https://goreportcard.com/badge/github.com/sshaplygin/docs-code)](https://goreportcard.com/report/github.com/sshaplygin/docs-code)

Go library for validation and generation of Russian official document codes.

> **Note:** The API is not yet stable and may change in future versions.

## Supported Document Codes

| Code   | Description                                          | Validate | Generate |
|--------|------------------------------------------------------|:--------:|:--------:|
| BIK    | Bank Identification Code (БИК)                       | +        | +        |
| INN    | Taxpayer Identification Number (ИНН)                 | +        | +        |
| KPP    | Tax Registration Reason Code (КПП)                   | +        | +        |
| OGRN   | Primary State Registration Number (ОГРН)             | +        | +        |
| OGRNIP | Primary State Registration Number for IE (ОГРНИП)    | +        | +        |
| SNILS  | Insurance Individual Account Number (СНИЛС)          | +        | +        |
| OKATO  | Russian Classification of Administrative Territories | -        | -        |

## Requirements

- Go 1.25+

## Installation

```bash
go get github.com/sshaplygin/docs-code
```

## Usage

### Validate a document code

```go
import (
	"log"

	docs_code "github.com/sshaplygin/docs-code"
)

isValid, err := docs_code.Validate(docs_code.INN, "526317984689")
if err != nil {
	log.Fatal(err)
}
if !isValid {
	log.Println("INN is invalid")
} else {
	log.Println("INN is valid")
}
```

### Generate a document code

```go
import (
	"fmt"

	docs_code "github.com/sshaplygin/docs-code"
)

code := docs_code.Generate(docs_code.INN)
fmt.Println("Generated INN:", code)
```

### Use a specific package directly

Each document type also exposes its own package with additional functions:

```go
import "github.com/sshaplygin/docs-code/inn"

// Generate specific INN types
legalINN := inn.GenerateLegal()
physicalINN := inn.GeneratePhysical()

// Validate
ok, err := inn.Validate("526317984689")
```

## Benchmarks

See [BENCHMARKS.md](BENCHMARKS.md) for performance data.

## Project Structure

```text
docs-code/
├── bik/          # BIK validation and generation
├── inn/          # INN validation and generation
├── kpp/          # KPP validation and generation
├── ogrn/         # OGRN validation and generation
├── ogrnip/       # OGRNIP validation and generation
├── okato/        # OKATO (in progress)
├── snils/        # SNILS validation and generation
├── fts/          # Federal Tax Service data (regions, departments)
├── models/       # Shared error types
├── utils/        # Internal helpers
└── parser/       # Data parsers (BIK, OKATO, subjects)
```

## References

- [INN (ИНН)](https://ru.wikipedia.org/wiki/%D0%98%D0%B4%D0%B5%D0%BD%D1%82%D0%B8%D1%84%D0%B8%D0%BA%D0%B0%D1%86%D0%B8%D0%BE%D0%BD%D0%BD%D1%8B%D0%B9_%D0%BD%D0%BE%D0%BC%D0%B5%D1%80_%D0%BD%D0%B0%D0%BB%D0%BE%D0%B3%D0%BE%D0%BF%D0%BB%D0%B0%D1%82%D0%B5%D0%BB%D1%8C%D1%89%D0%B8%D0%BA%D0%B0)
- [SNILS (СНИЛС)](http://www.consultant.ru/document/cons_doc_LAW_124607/68ac3b2d1745f9cc7d4332b63c2818ca5d5d20d0/)
- [OGRN (ОГРН)](https://ru.wikipedia.org/wiki/%D0%9E%D1%81%D0%BD%D0%BE%D0%B2%D0%BD%D0%BE%D0%B9_%D0%B3%D0%BE%D1%81%D1%83%D0%B4%D0%B0%D1%80%D1%81%D1%82%D0%B2%D0%B5%D0%BD%D0%BD%D1%8B%D0%B9_%D1%80%D0%B5%D0%B3%D0%B8%D1%81%D1%82%D1%80%D0%B0%D1%86%D0%B8%D0%BE%D0%BD%D0%BD%D1%8B%D0%B9_%D0%BD%D0%BE%D0%BC%D0%B5%D1%80)
- [OGRNIP (ОГРНИП)](http://www.temabiz.com/terminy/chto-takoe-ogrnip.html)
- [BIK (БИК)](https://ru.wikipedia.org/wiki/%D0%91%D0%B0%D0%BD%D0%BA%D0%BE%D0%B2%D1%81%D0%BA%D0%B8%D0%B9_%D0%B8%D0%B4%D0%B5%D0%BD%D1%82%D0%B8%D1%84%D0%B8%D0%BA%D0%B0%D1%86%D0%B8%D0%BE%D0%BD%D0%BD%D1%8B%D0%B9_%D0%BA%D0%BE%D0%B4)
- [KPP (КПП)](https://dic.academic.ru/dic.nsf/ruwiki/239834)

## License

[MIT](LICENSE)

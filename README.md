# docs-code

[![GoDoc](https://godoc.org/github.com/sshaplygin/docs-code?status.svg)](https://godoc.org/github.com/sshaplygin/docs-code) [![Go Coverage](https://github.com/sshaplygin/docs-code/wiki/coverage.svg)](https://raw.githack.com/wiki/sshaplygin/docs-code/coverage.html) [![License](https://img.shields.io/badge/license-MIT-blue.svg)](https://github.com/sshaplygin/docs-code/blob/master/LICENSE) [![Go Report Card](https://goreportcard.com/badge/github.com/sshaplygin/docs-code)](https://goreportcard.com/report/github.com/sshaplygin/docs-code)

It is validator about official of documents codes in Go

## Alert

It is not production ready public API! It is API could be change it the future. To be attentive for this.

## Status

Status of implementation by code package

- [ ] BIK
  - [x] Generation method
  - [ ] Validation method
- [x] INN
  - [x] Generation method
  - [x] Validation method
- [ ] KPP
  - [x] Generation method
  - [ ] Validation method
- [ ] OGRN
  - [x] Generation method
  - [ ] Validation method
- [ ] OGRNIP
  - [x] Generation method
  - [ ] Validation method
- [ ] OKATO
  - [ ] Generation method
  - [ ] Validation method
- [ ] OKTMO
  - [ ] Generation method
  - [ ] Validation method
- [ ] SNILS
  - [x] Generation method
  - [ ] Validation method
- [ ] Swift
  - [x] Generation method
  - [ ] Validation method

## Usage

``` bash
go get github.com/sshaplygin/docs-code
```

### Example

```go

import (
	"log"
	
	"github.com/sshaplygin/docs-code/inn"
)

...

isValid, err := inn.Validate("526317984689")
if err != nil {
    log.Error(err)
}
if isValid {
    log.Println("INN is valid")
}
```

## Documentation

* Info about INN  - [ИНН](https://ru.wikipedia.org/wiki/%D0%98%D0%B4%D0%B5%D0%BD%D1%82%D0%B8%D1%84%D0%B8%D0%BA%D0%B0%D1%86%D0%B8%D0%BE%D0%BD%D0%BD%D1%8B%D0%B9_%D0%BD%D0%BE%D0%BC%D0%B5%D1%80_%D0%BD%D0%B0%D0%BB%D0%BE%D0%B3%D0%BE%D0%BF%D0%BB%D0%B0%D1%82%D0%B5%D0%BB%D1%8C%D1%89%D0%B8%D0%BA%D0%B0)
* Info about SNILS - [СНИЛС](http://www.consultant.ru/document/cons_doc_LAW_124607/68ac3b2d1745f9cc7d4332b63c2818ca5d5d20d0/)
* Info about OGRN - [ОГРН](https://ru.wikipedia.org/wiki/%D0%9E%D1%81%D0%BD%D0%BE%D0%B2%D0%BD%D0%BE%D0%B9_%D0%B3%D0%BE%D1%81%D1%83%D0%B4%D0%B0%D1%80%D1%81%D1%82%D0%B2%D0%B5%D0%BD%D0%BD%D1%8B%D0%B9_%D1%80%D0%B5%D0%B3%D0%B8%D1%81%D1%82%D1%80%D0%B0%D1%86%D0%B8%D0%BE%D0%BD%D0%BD%D1%8B%D0%B9_%D0%BD%D0%BE%D0%BC%D0%B5%D1%80)
* Info about OGRNIP - [ОГРНИП](http://www.temabiz.com/terminy/chto-takoe-ogrnip.html)
* Info about BIK - [БИК](https://ru.wikipedia.org/wiki/%D0%91%D0%B0%D0%BD%D0%BA%D0%BE%D0%B2%D1%81%D0%BA%D0%B8%D0%B9_%D0%B8%D0%B4%D0%B5%D0%BD%D1%82%D0%B8%D1%84%D0%B8%D0%BA%D0%B0%D1%86%D0%B8%D0%BE%D0%BD%D0%BD%D1%8B%D0%B9_%D0%BA%D0%BE%D0%B4)
* Info about KPP - [КПП](https://dic.academic.ru/dic.nsf/ruwiki/239834)

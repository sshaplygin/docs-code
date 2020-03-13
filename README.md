# ru-doc-codes
[![GoDoc](https://godoc.org/github.com/mrfoe7/ru-doc-code?status.svg)](https://godoc.org/github.com/mrfoe7/ru-doc-code) [![Build Status](https://travis-ci.org/mrfoe7/ru-doc-code.svg)](https://travis-ci.org/mrfoe7/ru-doc-code.svg?branch=master) [![License](https://img.shields.io/badge/license-MIT-blue.svg)](https://github.com/mrfoe7/ru-doc-code/blob/master/LICENSE) [![Go Report Card](https://goreportcard.com/badge/github.com/mrfoe7/ru-doc-code)](https://goreportcard.com/report/github.com/mrfoe7/ru-doc-code)

It is validator about official codes of documents from Russia in Go 

## Usage 

* go get github.com/mrfoe7/ru-doc-codes

### Example
 
```go

isValid, err := ru-doc-code.IsSNILSValid("112-233-445 95")
if err != nil {
    log.Fatal(err)
}
if isValid {
    log.Println("SNILS is valid")
}

```

## Documentation

- Check to valid INN  - [ИНН](https://ru.wikipedia.org/wiki/%D0%98%D0%B4%D0%B5%D0%BD%D1%82%D0%B8%D1%84%D0%B8%D0%BA%D0%B0%D1%86%D0%B8%D0%BE%D0%BD%D0%BD%D1%8B%D0%B9_%D0%BD%D0%BE%D0%BC%D0%B5%D1%80_%D0%BD%D0%B0%D0%BB%D0%BE%D0%B3%D0%BE%D0%BF%D0%BB%D0%B0%D1%82%D0%B5%D0%BB%D1%8C%D1%89%D0%B8%D0%BA%D0%B0)

- Check to valid SNILS - [СНИЛС](http://www.consultant.ru/document/cons_doc_LAW_124607/68ac3b2d1745f9cc7d4332b63c2818ca5d5d20d0/)

- Check to valid OGRN - [ОГРН](https://ru.wikipedia.org/wiki/%D0%9E%D1%81%D0%BD%D0%BE%D0%B2%D0%BD%D0%BE%D0%B9_%D0%B3%D0%BE%D1%81%D1%83%D0%B4%D0%B0%D1%80%D1%81%D1%82%D0%B2%D0%B5%D0%BD%D0%BD%D1%8B%D0%B9_%D1%80%D0%B5%D0%B3%D0%B8%D1%81%D1%82%D1%80%D0%B0%D1%86%D0%B8%D0%BE%D0%BD%D0%BD%D1%8B%D0%B9_%D0%BD%D0%BE%D0%BC%D0%B5%D1%80)

- Check to valid OGRNIP - [ОГРНИП](http://www.temabiz.com/terminy/chto-takoe-ogrnip.html)

- Check to valid BIK - [БИК](https://ru.wikipedia.org/wiki/%D0%91%D0%B0%D0%BD%D0%BA%D0%BE%D0%B2%D1%81%D0%BA%D0%B8%D0%B9_%D0%B8%D0%B4%D0%B5%D0%BD%D1%82%D0%B8%D1%84%D0%B8%D0%BA%D0%B0%D1%86%D0%B8%D0%BE%D0%BD%D0%BD%D1%8B%D0%B9_%D0%BA%D0%BE%D0%B4)

- Check to valid KPP - [КПП](https://ru.wikipedia.org/wiki/%D0%98%D0%B4%D0%B5%D0%BD%D1%82%D0%B8%D1%84%D0%B8%D0%BA%D0%B0%D1%86%D0%B8%D0%BE%D0%BD%D0%BD%D1%8B%D0%B9_%D0%BD%D0%BE%D0%BC%D0%B5%D1%80_%D0%BD%D0%B0%D0%BB%D0%BE%D0%B3%D0%BE%D0%BF%D0%BB%D0%B0%D1%82%D0%B5%D0%BB%D1%8C%D1%89%D0%B8%D0%BA%D0%B0#%D0%9A%D0%BE%D0%B4_%D0%BF%D1%80%D0%B8%D1%87%D0%B8%D0%BD%D1%8B_%D0%BF%D0%BE%D1%81%D1%82%D0%B0%D0%BD%D0%BE%D0%B2%D0%BA%D0%B8_%D0%BD%D0%B0_%D1%83%D1%87%D1%91%D1%82_(%D0%9A%D0%9F%D0%9F)

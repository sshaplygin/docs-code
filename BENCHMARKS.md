# Benchmarks

## Notice

Command to update result:

```bash
go test -benchmem -run=^$ -bench . github.com/sshaplygin/docs-code/snils >> BENCHMARKS.md
```

## BIK

```
goos: darwin
goarch: arm64
pkg: github.com/sshaplygin/docs-code/bik
BenchmarkValidateCorrect-10    	 8064954	       140.7 ns/op	     256 B/op	       3 allocs/op
BenchmarkGenerate-10           	  615589	      1972 ns/op	     240 B/op	      18 allocs/op
PASS
ok  	github.com/sshaplygin/docs-code/bik	2.635s
```

## INN

```
goos: darwin
goarch: arm64
pkg: github.com/sshaplygin/docs-code/inn
BenchmarkValidateCorrectLegal-10       	 2719714	       443.8 ns/op	     616 B/op	      23 allocs/op
BenchmarkValidateCorrectPhysical-10    	 2076355	       576.6 ns/op	     936 B/op	      30 allocs/op
BenchmarkGenerate-10                   	  394204	      3133 ns/op	     875 B/op	      42 allocs/op
BenchmarkGenerateLegal-10              	  354616	      3213 ns/op	     801 B/op	      41 allocs/op
BenchmarkGeneratePhysical-10           	  492985	      2419 ns/op	     974 B/op	      41 allocs/op
PASS
ok  	github.com/sshaplygin/docs-code/inn	7.215s
```

## KPP

```
goos: darwin
goarch: arm64
pkg: github.com/sshaplygin/docs-code/kpp
BenchmarkValidateCorrect-10    	 5280958	       218.9 ns/op	     216 B/op	       8 allocs/op
BenchmarkGenerate-10           	  484114	      2434 ns/op	     385 B/op	      22 allocs/op
PASS
ok  	github.com/sshaplygin/docs-code/kpp	2.810s
```

## OGRN

```
goos: darwin
goarch: arm64
pkg: github.com/sshaplygin/docs-code/ogrn
BenchmarkValidateCorrect-10    	 2583738	       457.3 ns/op	     728 B/op	      18 allocs/op
BenchmarkGenerate-10           	  294908	      3938 ns/op	     841 B/op	      45 allocs/op
PASS
ok  	github.com/sshaplygin/docs-code/ogrn	3.074s
```

## OGRNIP

```
goos: darwin
goarch: arm64
pkg: github.com/sshaplygin/docs-code/ogrnip
BenchmarkValidateCorrect-10    	 1991065	       580.4 ns/op	    1008 B/op	      24 allocs/op
BenchmarkGenerate-10           	  403179	      3100 ns/op	    1010 B/op	      46 allocs/op
PASS
ok  	github.com/sshaplygin/docs-code/ogrnip	3.411s
```

## SNILS

```
goos: darwin
goarch: arm64
pkg: github.com/sshaplygin/docs-code/snils
BenchmarkValidateCorrect-10    	 4451258	       263.2 ns/op	     336 B/op	       5 allocs/op
BenchmarkGenerate-10           	 1302042	       895.3 ns/op	     568 B/op	      25 allocs/op
PASS
ok  	github.com/sshaplygin/docs-code/snils	3.768s
```

# jsonlang-test
[![Build Status](https://travis-ci.org/DLag/jsonlang-test.svg?branch=master)](https://travis-ci.org/DLag/jsonlang-test)
[![codecov](https://codecov.io/gh/DLag/jsonlang-test/branch/master/graph/badge.svg)](https://codecov.io/gh/DLag/jsonlang-test)
[![Go Report Card](https://goreportcard.com/badge/github.com/DLag/jsonlang-test)](https://goreportcard.com/report/github.com/DLag/jsonlang-test)

Test task

### Setup
```bash
go get github.com/DLag/jsonlang-test
```
### Give it a try
```bash
$GOPATH/bin/jsonlang-test examples/script1.json examples/script2.json
```
### Test maximum execution depth
```bash
$GOPATH/bin/jsonlang-test --maxdepth=5 examples/recursion.json
```
### Test consistence checks
```bash
$GOPATH/bin/jsonlang-test examples/erroneous.json
```
### Check test coverage
```bash
go test -v -cover ./...
```
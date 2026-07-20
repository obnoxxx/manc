[![Build Status](https://github.com/obnoxxx/manc/actions/workflows/ci.yml/badge.svg?branch=main)](https://github.com/obnoxxx/manc/actions)
[![Go Report Card](https://goreportcard.com/badge/github.com/gontract/gontract)](https://goreportcard.com/report/github.com/obnoxxx/manc)
[![GPL license](https://img.shields.io/badge/license-LGPL-blue.svg)](http://opensource.org/licenses/LGPL)
[![Go Reference](https://pkg.go.dev/badge/github.com/gontract/gontract.svg)](https://pkg.go.dev/github.com/obnoxxx/manc)

# manc

`manc` provides ancillary floating point math functions for Go. It has essentially been extracted from the example code of [gontract](https://github.com/gontract/gontract)


## FloatEquals

The function `FloatEquals` reports if two floats are approximately equal,
taking into account imprecisions of binary representation and rounding.

```go
if manc.FloatEquals(result, expected) {
	// Values are approximately equal.
}
```

## Development

For building and testing, manc needs [mage](https://magefile.org/) installed in the path
as it uses a magefile instead of a Makefile.

Examples:

```sh
mage test
mage vet
mage check
```

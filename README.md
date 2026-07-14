# manc

`manc` provides ancillary math functions for Go.

## FoatEquals

The `FloatEquals` function reports if two floats are approximately eual,
takin into account imprecisions of binary representation and rouynding.

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

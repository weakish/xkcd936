# Xkcd936

[xkcd936](https://www.xkcd.com/936/) style passphrase generator with BIP32 word list.

## Installation

To compile from source, you need to have `go` installed.

    git clone https://github.com/weakish/xkcd936
    cd xkcd936
    make install

To change installation path (default to `/usr/local/bin`),
edit `config.mk` before running `make`.

To uninstall, run `make uninstall`.

The Makefile is compatible with both GNU make and BSD make.


## Usage

### As a Commandline Utility

Example output:

```sh
; xkcd936
subway modify spoil basic
```

To fulfill password requirements from some websites
(white space is not allowed and/or mix case is required),
xkcd936 can also output titlized words:

```sh
; xkcd936 -t
AdvanceJealousDevelopSenior
```

By default, xkcd936 will generate four words,
consistent with the [xkcd936] comic.
But xkcd936 can generate up to 12 words with the number specified in `-n`.

### As a Library

```go
import (
	"github.com/weakish/xkcd936/xkcd936"
)

var words []string = xkcd936.Words(4) // e.g. []string{"cannon", "isolate", "soccer", "word"}
var phrase string = xkcd936.Phrase(words, true) // e.g. "CannonIsolateSoccerWord"
```

## License

0BSD

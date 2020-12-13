# Xkcd936

[xkcd936](https://www.xkcd.com/936/) style passphrase generator with BIP32 word list.

## Installation

To compile from source, you need to have `go` installed.

    git clone https://github.com/weakish/xkcd936
    cd xkcd936
    make
    sudo make install

To change installation path (default to `/usr/local/bin`),
edit `config.mk` before running `make`.

To uninstall, run `make uninstall`.

The Makefile is compatible with both GNU make and BSD make.


## Usage

### As a Commandline Utility

```sh
xkcd936 n # n: from 1 to 12
xkcd936   # defaults to 4
```

Example output:

```go
AdvanceJealousDevelopSenior
```

### As a Library

```go
import (
	"github.com/weakish/xkcd936/xkcd936"
)

var words []string = xkcd936.Words(4) // e.g. []string{"cannon", "isolate", "soccer", "word"}
var phrase string = xkcd936.Phrase(words) // e.g. "CannonIsolateSoccerWord"
```

## License

0BSD

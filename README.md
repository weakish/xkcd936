# Xkcd936

[xkcd936](https://www.xkcd.com/936/) style passphrase generator with BIP32 word list.

## Installation

Installation from the source requires Go installed on the system.

    git clone https://github.com/weakish/xkcd935
    cd xkcd935

### Make

    make install

To change installation path (default to `/usr/local/bin`),
edit `config.mk` before running `make`.

To uninstall, run `make uninstall`.

The Makefile is compatible with both GNU make and BSD make.

### DoIt

It is also possible to use [DoIt] with Python >= 3.10 instead of make.

```sh
doit install
```

[DoIt]: https://pydoit.org

To change installation path (default to `/usr/local/bin`),
run `doit install --prefix /absolute/path` instead.

To uninstall, run `doit uninstall`.

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

By default, the BIP32 English word list is used.
To use a BIP32 word list in another language, specify it in `-l`.

Some diceware-derived lists can also be used:

    - `diceware{1,2,3,4}`: four lists split from [diceware8k], containing 2048 words each.
    - `diceware2k`: [a wordlist derived from diceware8k][diceware2k], containing 2048 short words.

[diceware8k]: https://theworld.com/~reinhold/diceware8k.txt
[diceware2k]: https://diceware2k.surge.sh

### As a Library

```go
import (
	"github.com/weakish/xkcd936/xkcd936"
)

// nil means using the default BIP32 English word list
var words []string = xkcd936.Words(4, nil) // e.g. []string{"cannon", "isolate", "soccer", "word"}
var phrase string = xkcd936.Phrase(words, true) // e.g. "CannonIsolateSoccerWord"
```

## License

0BSD

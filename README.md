# Xkcd936

[xkcd936](https://www.xkcd.com/936/) style passphrase generator.

## Dependencies

- libsodium

## Installation

    git clone https://github.com/weakish/xkcd936
    cd xkcd936
    make

To change C compiler (default to clang) and installation path (default to `/usr/local/bin`),
edit `config.mk` before running `make`.

To uninstall, run `make uninstall`.

The Makefile is compatible with both GNU make and BSD make.

## Usage

    xkcd936

## License

0BSD

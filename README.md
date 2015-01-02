# Xkcd936

[![Gem Version](https://badge.fury.io/rb/xkcd936.svg)](http://badge.fury.io/rb/xkcd936)

[xkcd936](https://www.xkcd.com/936/) style passphrase generator.

But we use 5 words instead of 4, since 5 words will achieve an entropy of:

    log(99171)/log(2)*5 = 82.988

This is enough.
NIST recommends 80-bits for the most secure passwords.
And it roughly needs 6 billions USD to break.

## Installation

Add this line to your application's Gemfile:

```ruby
gem 'xkcd936'
```

And then execute:

    $ bundle

Or install it yourself as:

    $ gem install xkcd936

## Usage

As a library:

```ruby
require 'xkcd936'
Xkcd936.generate_passphrase('path/to/dictionary/file')
```

If you does not give a path, it will use `/usr/share/dict/words` as default.

As a command line tool:

```ruby
; xkcd936
```

## Contributing

1. Fork it ( https://github.com/weakish/xkcd936/fork )
2. Create your feature branch (`git checkout -b my-new-feature`)
3. Commit your changes (`git commit -am 'Add some feature'`)
4. Push to the branch (`git push origin my-new-feature`)
5. Create a new Pull Request

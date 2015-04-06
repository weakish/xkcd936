# coding: utf-8
lib = File.expand_path('../lib', __FILE__)
$LOAD_PATH.unshift(lib) unless $LOAD_PATH.include?(lib)
require 'xkcd936/version'

Gem::Specification.new do |spec|
  spec.name          = 'xkcd936'
  spec.version       = Xkcd936::VERSION
  spec.authors       = ['Jakukyo Friel']
  spec.email         = ['weakish@gmail.com']
  spec.summary       = %q{xkcd936 style passphrase generator.}
  spec.homepage      = 'https://github.com/weakish/xkcd936'
  spec.license       = 'MIT'
=begin gem < 2.0 does not support metadata
  spec.metadata      = {
      'repository' => 'https://github.com/weakish/xkcd936.git',
      'documentation' => 'http://www.rubydoc.info/gems/xkcd936',
      'issues' => 'https://github.com/weakish/xkcd936/issues',
      'wiki' => 'https://github.com/weakish/xkcd936/wiki'
  }
=end
  spec.files         = `git ls-files -z`.split("\x0")
  spec.executables   = spec.files.grep(%r{^bin/}) { |f| File.basename(f) }
  spec.test_files    = spec.files.grep(%r{^(test|spec|features)/})
  spec.require_paths = ['lib']

  spec.required_ruby_version = '>= 1.9.3'
  spec.add_development_dependency 'bundler', '~> 1.3'
  spec.add_development_dependency 'rake', '>= 10.0.0'
end

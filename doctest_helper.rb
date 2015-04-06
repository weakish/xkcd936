require 'lib/xkcd936'
require 'pry-rescue/minitest'

open('lib/xkcd936.rb') do |f|
  module_lines = f.find_all { |line| line[0..5] == 'module' }
  modules = module_lines.map { |s| s.strip.tr(' ', '')[6..-1] }
  modules.each do |mod|
    mod = Object.const_get(mod)
    include mod
    YARD::Doctest.configure do |doctest|
      mod.private_instance_methods.each do |m|
        doctest.skip "#{mod}.#{m.to_s}"
      end
    end
  end
end
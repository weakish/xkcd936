require "xkcd936/version"

module Xkcd936
  module_function

  # Inspired by http://xkcd.com/936/
  # But I use 5 words instead of 4, since 5 words will achieve an entropy of:
  #     log(99171)/log(2)*5 = 82.988
  # This is enough.
  # NIST recommends 80-bits for the most secure passwords.
  # And it roughly needs 6 billions USD to break.
  #
  # @param [String] dict file path
  # @return [String] passphrase
  def generate_passphrase(dict='/usr/share/dict/words')
    (IO.readlines(dict).sample(5).map { |s| s.chomp.capitalize }).join
  end
end

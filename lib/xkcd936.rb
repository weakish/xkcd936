require 'xkcd936/version'
require 'i18n'

module Xkcd936
  module_function

  # Inspired by http://xkcd.com/936/
  # But I use 5 words as default instead of 4, since 5 words will achieve
  # an entropy of:
  #     log(99171)/log(2)*5 = 82.988
  # This is enough.
  # NIST recommends 80-bits for the most secure passwords.
  # And it roughly needs 6 billions USD to break.
  #
  # @param [Fixnum] number of words
  # @param [String] dict file path
  # @return [String] passphrase
  #
  # @example default values
  #   generate_passphrase
  # @example pick up four words
  #   generate_passphrase(4)
  def generate_passphrase(number=5, dict='/usr/share/dict/words')
    pick_five = IO.readlines(dict).sample(number)
    passphrase = pick_five.map { |s| s.chomp.capitalize }.join
    # Let I18n skip validation and fall back to built in :en locale.
    # Otherwise you need to have a locale file.
    I18n.enforce_available_locales = false
    I18n.transliterate(passphrase)
  end
end

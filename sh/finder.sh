####################################################################################################

# turn finder hidden files on
finder_on() {
  defaults write com.apple.Finder AppleShowAllFiles true && killall Finder
}

# turn finder hidden files off
finder_off() {
  defaults write com.apple.Finder AppleShowAllFiles false && killall Finder
}

####################################################################################################

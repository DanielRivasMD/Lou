####################################################################################################

# display fpath
fpath() {
  source "${IN_SILICO}/.config/config.sh"
  source "${IN_SILICO}/.break/DL.sh"  $COLUMNS
  echo $FPATH | xsv flatten -d ":" -n | awk '{print $2}'
  source "${IN_SILICO}/.break/DL.sh"  $COLUMNS
}

# display GOPATH
gopath() {
  source "${IN_SILICO}/.config/config.sh"
  source "${IN_SILICO}/.break/DL.sh" $COLUMNS
  echo $GOPATH | xsv flatten -d ":" -n | awk '{print $2}'
  source "${IN_SILICO}/.break/DL.sh" $COLUMNS
}

# display PATH
path() {
  source "${IN_SILICO}/.config/config.sh"
  source "${IN_SILICO}/.break/DL.sh" $COLUMNS
  echo $PATH | xsv flatten -d ":" -n | awk '{print $2}'
  source "${IN_SILICO}/.break/DL.sh" $COLUMNS
}

####################################################################################################

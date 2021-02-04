################################################################################

# install Lou locally
install:
  echo "Install..."
  # Lou
  cd $HOME/.go/src/github.com/DanielRivasMD/Lou
  go install
  cd $HOME/.go/bin
  mv -v lou $HOME/bin/goTools/

################################################################################

# link Lou repository to GOPATH
link:
  echo "Linking..."
  ln -svf $HOME/Factorem/Lou $HOME/.go/src/github.com/DanielRivasMD

################################################################################


################################################################################

_default:
  @just --list

################################################################################

# install Lou locally
@ install:
  echo "Install..."
  # Lou
  cd ${HOME}/.go/src/github.com/DanielRivasMD/Lou && go install
  mv -v ${HOME}/.go/bin/lou ${HOME}/bin/goTools/

################################################################################

# link Lou repository to GOPATH
@ link:
  echo "Linking..."
  ln -svf ${HOME}/Factorem/Lou ${HOME}/.go/src/github.com/DanielRivasMD

################################################################################

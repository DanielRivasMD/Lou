####################################################################################################

_default:
  @just --list

####################################################################################################

# print justfile
print:
  bat .justfile --language make

####################################################################################################

# edit justfile
edit:
  micro .justfile

####################################################################################################

# aliases
alias b := buildOSX
alias i := install

####################################################################################################

# build bender for osx & store `excalibur`
buildOSX:
  #!/bin/bash
  set -euo pipefail

  # declarations
  source .just.sh

  echo "Building..."
  go build -v -o ${lou}/excalibur/lou

####################################################################################################

# build bender for linux & store `excalibur`
buildLinux:
  #!/bin/bash
  set -euo pipefail

  # declarations
  source .just.sh

  echo "Building..."
  env GOOS=linux GOARCH=amd64 go build -v -o ${lou}/excalibur/lou

####################################################################################################

# install Lou locally
install:
  #!/bin/bash
  set -euo pipefail

  # declarations
  source .just.sh

  echo "Install..."
  # Lou
  go install
  mv -v ${HOME}/.go/bin/Lou ${HOME}/.go/bin/lou

####################################################################################################

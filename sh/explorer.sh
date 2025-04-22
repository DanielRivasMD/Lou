#!/bin/bash

################################################################################

cmd=$1
file=$2

if [[ "$cmd" == "lsd" ]]; then
  cmd="lsd --header --long --classify --git"
fi

# run
zellij run \
  --name canvas \
  --floating \
  --height 100 \
  --width 100 \
  --x 100 \
  --y 0 \
  -- "$cmd" "$file"

################################################################################

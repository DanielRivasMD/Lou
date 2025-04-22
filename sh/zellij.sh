#!/bin/bash

################################################################################

cmd=$1
file=$2

if [[ $1 == "lsd" ]]; then
  zellij run \
    --name canvas \
    --floating \
    --height 100 \
    --width 100 \
    --x 100 \
    --y 0 \
    -- "$cmd" --header --long --classify --git "$file"
else
  zellij run \
    --name canvas \
    --floating \
    --height 100 \
    --width 100 \
    --x 100 \
    --y 0 \
    -- "$cmd" "$file"
fi


################################################################################

#!/bin/bash

################################################################################

# TODO: take arguments to control movement. defaults '~/Downloads' & '~/Articles/Thesis'
function theses() {

  # declare variables
  cd /Users/drivas/Downloads/
  shDirectory="/Users/drivas/Factorem/Lou/sh/"
  typeArr=(pdf)
  folderArr=(Thesis)

  for (( ix = 0; ix < ${#typeArr[@]}; ++ix ))
  do

    # TODO: modify regex to cover two letter author last name
    # collect files
    fs2move=$( find -E . -type f -regex "./[A-Z][a-z-]+_[A-Z][a-z-]+_Thesis_[0-9]{4}.${typeArr[ix]}" )

    if [[ -z ${fs2move} ]]
    then
      emptyMessage=`cat <<- TESTA
\tNo thesis files to relocate:	${folderArr[ix]}`

      echo -e "${emptyMessage}"
    else

      # relocate
      while read ori flies
      do
        mv -v "${ori}" ~/Articulos/${folderArr[ix]}/"${flies}" | \
        awk '{sub("/Users/drivas", "~", $3); printf "%-70s ->     %25s\n", $1, $3}'
      done < <( echo -e "${fs2move}" )
    fi
  done

  cd - > /dev/null
}

################################################################################

theses

################################################################################


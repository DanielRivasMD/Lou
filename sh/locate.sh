#!/bin/bash

################################################################################

function move() {

  # declare variables
  cd /Users/drivas/Downloads/
  shDirectory="/Users/drivas/Factorem/Lou/sh/"
  typeArr=(pdf ris)
  folderArr=(PDFs Refs)

  for (( ix = 0; ix < ${#typeArr[@]}; ++ix ))
  do

    # collect files
    fs2move=$( find -E . -type f -regex "./[A-Z][a-z]+[-]{1}[A-Za-z_-]+.${typeArr[ix]}")

    if [[ -z ${fs2move} ]]
    then
      emptyMessage=`cat <<- TESTA
\tNo files to relocate:	${folderArr[ix]}`

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

move

################################################################################

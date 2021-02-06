#!/bin/bash

################################################################################

# TODO: take arguments to control movement. defaults '~/Downloads' & '~/Articles/PDFs'
function form() {

  # declare variables
  cd /Users/drivas/Downloads/
  shDirectory="/Users/drivas/Factorem/Lou/sh/"
  typeArr=(pdf ris)
  folderArr=(PDFs Refs)

  for (( ix = 0; ix < ${#typeArr[@]}; ++ix ))
  do

    # TODO: modify regex to cover two letter author last name
    # collect files
    fs2move=$( find -E . -type f -regex "./[A-Z][a-z-]+[0-9]{4}[A-Za-z-]+.${typeArr[ix]}" | \
    awk -f ${shDirectory}article_formatter.awk -v suffix=${typeArr[ix]} 2>&1 )

    if [[ -z ${fs2move} ]]
    then
      emptyMessage=`cat <<- TESTA
\No files to reformat:	${folderArr[ix]}`

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

form

################################################################################

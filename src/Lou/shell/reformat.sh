#!/bin/bash

################################################################################

function form() {

  # declare variables
  cd /Users/drivas/Downloads/
  shellDirectory="/Users/drivas/Factorem/Lou/src/papel/shell/"
  typeArr=(pdf ris)
  folderArr=(PDFs Refs)

  for (( ix = 0; ix < ${#typeArr[@]}; ++ix ))
  do

    # collect files
    fs2move=$( find -E . -type f -regex "./[A-Z][a-z-]+[0-9]{4}[A-Za-z-]+.${typeArr[ix]}" | \
    awk -f ${shellDirectory}article_formatter.awk -v suffix=${typeArr[ix]} 2>&1 )

    if [[ -z ${fs2move} ]]
    then
      emptyMessage=`cat <<- TESTA
\tThere were no files:	${folderArr[ix]}`

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

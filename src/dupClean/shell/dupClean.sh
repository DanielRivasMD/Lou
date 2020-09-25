#!/bin/bash

################################################################################

function cleaning() {
  # duplicates to clean at ~/Downloads/
  dups_to_clean=$( ls ~/Downloads/*\(*\)* 2>&1 )

  expectedError="No such file or directory"
  if [[ ${dups_to_clean} =~ ${expectedError} ]]
  then
    emptyMessage=`cat <<- TESTA
    \t\t\tThere were no files to remove`

    echo -e "${emptyMessage}"
  else
    removingFiles=`cat <<- TESTA
    \tFiles to remove:
    ${dups_to_clean}
    TESTA`
    echo -e "${removingFiles}"
    rm /Users/drivas/Downloads/*\(*\)*
  fi
}

################################################################################

cleaning

################################################################################

#!/bin/bash

################################################################################

location=$1

################################################################################

# function declaration
function cleaning() {
  # duplicates to clean
  dups_to_clean=$( ls ${location}/*\(*\)* 2>&1 )

  expectedError="No such file or directory"
  if [[ ${dups_to_clean} =~ ${expectedError} ]]
  then
    emptyMessage=`cat <<- TESTA
\tThere were no files to remove
TESTA`

    echo -e "${emptyMessage}"
  else
    removingFiles=`cat <<- TESTA
\tFiles to remove:
${dups_to_clean}
TESTA`

    echo -e "${removingFiles}"
    rm ${location}/*\(*\)*
  fi
}

################################################################################

# function call
cleaning

################################################################################

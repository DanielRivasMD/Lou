#!/bin/bash

################################################################################

# declare variables
articleDirectory="/Users/drivas/Articulos"
articlePDF=$1

if [[ ! -e ${articleDirectory}/PDFs/${articlePDF} ]]
then
  echo "File does not exist" && exit 1    # test whether file exists
elif [[ -L ${articleDirectory}/PDFs/${articlePDF} ]]
then
  echo "File is already linked" && exit 1 # test whether file is linked
else
  mv ${articleDirectory}/PDFs/${articlePDF} ${articleDirectory}/Notes/
  ln -s ${articleDirectory}/Notes/${articlePDF} ${articleDirectory}/PDFs/
  echo "Annotate -> ${articlePDF}"
fi

################################################################################

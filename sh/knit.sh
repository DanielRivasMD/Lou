####################################################################################################

# rmarkdown knit
knit() {
  R --slave -e "rmarkdown::render('$1')" > /dev/null
}

####################################################################################################

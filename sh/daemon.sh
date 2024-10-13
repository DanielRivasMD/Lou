####################################################################################################

# check daemon status
check_daemon() {
  if test ! -f "${pueuedTracker}" && test ! -f "${pueueSocket}"
  then
    source "${IN_SILICO}/.break/Cross.sh" $COLUMNS
    echo ''
    echo "${CYAN}pueue${YELLOW} daemon${NC} is not currently running${NC}"
    echo ''
    echo "start by calling ${CYAN}pueued${GREEN} --daemonize${NC}"
    echo ''
    source "${IN_SILICO}/.break/Cross.sh" $COLUMNS
    return 1
  else
    return 0
  fi
}

####################################################################################################

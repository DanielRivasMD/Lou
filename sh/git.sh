####################################################################################################

status() {
  # git status & git stash & git log
  source "${IN_SILICO}/.config/config.sh"

  source "${IN_SILICO}/.break/Cross.sh" $COLUMNS

  git status --short
  source "${IN_SILICO}/.break/SL.sh" $COLUMNS

  git stash list
  source "${IN_SILICO}/.break/SL.sh" $COLUMNS

  git log --graph --topo-order --abbrev-commit --date=relative --decorate --all --boundary --pretty=format:'%Cgreen%ad %Cred%h%Creset -%C(yellow)%d%Creset %s %C(dim white)%cn%Creset' -10
  source "${IN_SILICO}/.break/Cross.sh" $COLUMNS
}

####################################################################################################
####################################################################################################

# bat delta
diff() {
  git diff --name-only --relative --diff-filter=d | xargs bat --diff
}

####################################################################################################

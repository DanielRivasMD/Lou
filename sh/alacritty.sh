####################################################################################################
# alacritty
####################################################################################################

# config
source "${HOME}/.lou/sh/config.sh"

# check & remove target
if test -f "${config}/alacritty/alacritty.toml"
then
  rm "${config}/alacritty/alacritty.toml"
fi

# migrate config
alacritty migrate

####################################################################################################


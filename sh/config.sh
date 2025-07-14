####################################################################################################
# config
####################################################################################################

# home
export home="${HOME}"
export config="${home}/.config"
export archive="${home}/.archive"
export in_situ="${archive}/in-situ"
export ex_situ="${archive}/ex-situ"
export in_silico="${archive}/in-silico"
export forked="${home}/Forked"
export observatory="${home}/Observatory"
export completion="${home}/.completion"
export linked="${home}/Linked"
export appSupport="${home}/Library/Application Support"
export localShare="${home}/.local/share"
export completionArch="${archive}/completions"

####################################################################################################
# in-situ
####################################################################################################

# git files
export git="${in_situ}/git"
export gitconfig="${git}/gitconfig"
export gitignore="${git}/gitignore_global"

# rc config
export mplayer="${in_situ}/mplayer/mplayer"

# toml files
export procs="${in_situ}/procs/procs.toml"

# config files
export atuin="${in_situ}/atuin"
export alacritty="${in_situ}/alacritty"
export bottom="${in_situ}/bottom"
export joplin="${in_situ}/joplin"

export gh="${in_situ}/gh"
export spotify="${in_situ}/spotify"
export saiyajin="${home}/.saiyajin"
export karabiner="${saiyajin}/karabiner"
export frag="${saiyajin}/frag"
export fapps="${frag}/apps"
export fprofile="${frag}/profile"
export fmode="${frag}/mode"
export fsimple="${frag}/simple"

# config dirs
export gitui="${in_situ}/gitui"
export zellij="${in_situ}/zellij"

# distant
export julia="${in_situ}/julia"
export sshConfig="${in_situ}/ssh/config"
export lazycli="${in_situ}/lazycli"
export lazycliConf="${appSupport}/lazycli"
export lazygit="${in_situ}/lazygit"
export lazygitConf="${appSupport}/jesseduffield/lazygit"
export halp="${in_situ}/halp"
export halpConf="${appSupport}/halp"

# broot
export broot="${in_situ}/broot"
export brootConf="${home}/.config/broot"

# espanso
export espanso="${in_situ}/espanso"
export espansoConfig="${espanso}/config"
export espansoMatch="${espanso}/match"
export espansoConf="${appSupport}/espanso"

# navi
export navi="${in_situ}/navi"
export naviConf="${appSupport}/navi/"

####################################################################################################
# ex-situ
####################################################################################################

# @HOME
export screen="${ex_situ}/screen"
export mycli="${ex_situ}/mycli"

# @config
export bpython="${ex_situ}/bpython"
export helix="${ex_situ}/helix"
export moded="${helix}/modes"
export litecli="${ex_situ}/litecli"
export lsd="${ex_situ}/lsd"
export micro="${ex_situ}/micro"
export sheldon="${ex_situ}/sheldon"
export starship="${ex_situ}/starship"

# shell
export shell="${ex_situ}/shell"
export term="${shell}/term"

# bash
export bash="${shell}/bash"

# zsh
export zsh="${shell}/zsh"

# nushell
export nushell="${shell}/nu"
export nushellConf="${appSupport}/nu"

####################################################################################################
# ergo
####################################################################################################

# completions
export zshcomp="${config}/zsh_completion"

####################################################################################################

# remote
export remoteBin="${home}/RemoteBin"
export homeRemote="/home/drivas"
export ex_situRemote="${homeRemote}/.archive/ex-situ"
export shellRemote="${ex_situRemote}/shell"
export termRemote="${shellRemote}/term"
export bashRemote="${shellRemote}/bash"
export zshRemote="${shellRemote}/zsh"

# pawsey
export pawseyID="drivas@topaz.pawsey.org.au"
export softwarePawsey="/scratch/pawsey0263/drivas/software"

####################################################################################################

# pueue
export pueuedTracker="${appSupport}/pueue/pueue.pid"
export pueuedSocket="${appSupport}/pueue/pueue_drivas.socket="
export pueueFG="forge"
export pueueUG="update"

####################################################################################################

# # load colors
# source "${archive}/.just/colors.sh"

####################################################################################################

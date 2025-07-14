####################################################################################################
# goku
####################################################################################################

# config
source "${HOME}/.lou/sh/config.sh"

# format
cljfmt fix "${frag}/"*

# create temporary files
cat << HEAD >> "${saiyajin}/.profile.tmp"

{:profiles

HEAD

cat << HEAD >> "${saiyajin}/.main.tmp"

:main [

HEAD

cat << HEAD >> "${saiyajin}/.eof.tmp"

  ]}]}

HEAD

# concatenate
mbombo forge --path "${frag}" --out "${karabiner}/karabiner.edn" \
  --files "../.profile.tmp" \
  --files "profile/profile.edn" \
  --files "../.main.tmp" \
  --files "apps/alacritty.edn" \
  --files "apps/browser.edn" \
  --files "apps/mail.edn" \
  --files "apps/zellij.edn" \
  --files "apps/zoom.edn" \
  --files "double/keypad.edn" \
  --files "double/lcmd.edn" \
  --files "double/lctl.edn" \
  --files "mode/mouse.edn" \
  --files "mode/obracket.edn" \
  --files "mode/cbracket.edn" \
  --files "mode/semicolon.edn" \
  --files "mode/quote.edn" \
  --files "mode/backslash.edn" \
  --files "mode/slash.edn" \
  --files "simple/bs.edn" \
  --files "simple/claw.edn" \
  --files "simple/esc.edn" \
  --files "simple/hyper.edn" \
  --files "simple/joker.edn" \
  --files "simple/loptcmd.edn" \
  --files "simple/loptctl.edn" \
  --files "simple/lctlcmd.edn" \
  --files "simple/lcmd.edn" \
  --files "simple/lctl.edn" \
  --files "simple/lopt.edn" \
  --files "simple/lshift.edn" \
  --files "simple/patch.edn" \
  --files "simple/rcmd.edn" \
  --files "simple/rctl.edn" \
  --files "simple/rshift.edn" \
  --files "simple/ropt.edn" \
  --files "simple/tab.edn" \
  --files "simple/zero.edn" \
  --files "profile/keys.edn" \
  --files "../.eof.tmp"

# purge temporary files
rm "${saiyajin}/.profile.tmp"
rm "${saiyajin}/.main.tmp"
rm "${saiyajin}/.eof.tmp"

# render config
goku

####################################################################################################

####################################################################################################
# helix
####################################################################################################

# config
source "${HOME}/.lou/sh/config.sh"

# create temporary files
for type in normal insert select
do
  mbombo forge --path "${moded}" --out "${helix}/.${type}.tmp" --files "common.toml" --old MODE --new "${type}_mode"
done

# concatenate
mbombo forge --path "${moded}" --out "${helix}/config.toml" \
  --files "theme.toml" \
  --files "editor.toml" \
  --files "mini-mode.toml" \
  --files "normal.toml" \
  --files "../.normal.tmp" \
  --files "insert.toml" \
  --files "../.insert.tmp" \
  --files "select.toml" \
  --files "../.select.tmp"

# purge temporary files
rm "${helix}/.normal.tmp"
rm "${helix}/.insert.tmp"
rm "${helix}/.select.tmp"

####################################################################################################

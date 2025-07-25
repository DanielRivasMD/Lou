####################################################################################################

_default:
  @just --list

####################################################################################################

# print justfile
[group('just')]
@show:
  bat .justfile --language make

####################################################################################################

# edit justfile
[group('just')]
@edit:
  micro .justfile

####################################################################################################

# compile roadmap action items
[group('dev')]
@roadmap:
  { echo '\n=================================================='; todor -s; echo '=================================================='; } >> ROADMAP.txt

####################################################################################################
# aliases
####################################################################################################

####################################################################################################
# import
####################################################################################################

# config
import '.just/go.conf'

####################################################################################################
# jobs
####################################################################################################

# build for OSX
[group('go')]
osx app=goapp:
  @echo "\n\033[1;33mBuilding\033[0;37m...\n=================================================="
  go build -v -o excalibur/{{app}}

####################################################################################################

# build for linux
[group('go')]
linux app=goapp:
  @echo "\n\033[1;33mBuilding\033[0;37m...\n=================================================="
  env GOOS=linux GOARCH=amd64 go build -v -o excalibur/{{app}}

####################################################################################################

# install locally
[group('go')]
install app=goapp exe=goexe dir=dir:
  @echo "\n\033[1;33mInstalling\033[0;37m...\n=================================================="
  go install
  @echo "\n\033[1;33mLinking\033[0;37m...\n=================================================="
  @mv -v "${HOME}/go/bin/{{app}}" "${HOME}/go/bin/{{exe}}"
  @echo "\n\033[1;33mCopying\033[0;37m...\n=================================================="
  @if [ ! -d "${HOME}/{{dir}}" ]; then mkdir "${HOME}/{{dir}}"; fi
  @if test -e "${HOME}/{{sh}}"; then rm -r "${HOME}/{{sh}}"; fi && echo "\033[1;33msh\033[0;37m" && cp -v -R "sh" "${HOME}/{{sh}}"
  @if test -e "${HOME}/{{layouts}}"; then rm -r "${HOME}/{{layouts}}"; fi && echo "\033[1;33mlayouts\033[0;37m" && cp -v -R "layouts" "${HOME}/{{layouts}}"

####################################################################################################

# watch changes
[group('go')]
watch:
  watchexec --clear --watch cmd -- 'just install'

####################################################################################################

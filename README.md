
## Lou is a personal assitant to automatize routine tasks

![](assets/Lou.png)


## Table of contents

  - [Table of contents](#table-of-contents)
  - [Overview](#overview)
  - [Installation](#installation)
    - [Via Homebrew (for macOS)](#via-homebrew-for-macos)
    - [Via APT (for Debian-based Linux distros)](#via-apt-for-debian-based-linux-distros)
    - [From Github release](#from-github-release)
  - [Documentation](#documentation)
    - [Usage](#usage)
  - [Examples](#examples)
  - [Acknowledgements](#acknowledgements)
  - [License](#license)


## Overview


## Installation


<!-- TODO: -->
**Not currently working**

### Via Homebrew (for macOS)

Prerequisites:

- [Homebrew](https://brew.sh/)

```
brew install danielrivasmd/Lou
```



<!-- TODO: -->
**Not currently working**

### Via APT (for Debian-based Linux distros)

```
curl -SsL https://fbecart.github.io/ppa/debian/KEY.gpg | sudo apt-key add -
sudo curl -SsL -o /etc/apt/sources.list.d/fbecart.list https://fbecart.github.io/ppa/debian/fbecart.list
sudo apt update
sudo apt install lou
```



<!-- TODO: -->
**Not currently working**

### From Github release

<!-- TODO: -->
**Not currently working**



## Documentation

### Usage

Use `lou -h` or `lou --help` to display help on commandline.

```
Lou, personal assitant at your service

Usage:
  lou [command]

Available Commands:
  help                  Help about any command

Flags:
  -h, --help            help for Lou
  -o, --outDir string   Output directory. creates if not exitst

Use "lou [command] --help" for more information about a command.
```



## Examples

```
lou help
lou clean -l $(pwd)
lou biblo format
```


## Acknowledgements



## License

Lou is distributed under the terms of the GNU GENERAL PUBLIC LICENSE.

See [LICENSE](LICENSE) for details.


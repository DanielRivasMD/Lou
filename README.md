
## Lou is a personal assitant to automatize routine tasks

![](assets/Lou.png)


## Table of contents

- [Overview](#overview)
- [Installation](#installation)
  - [Homebrew](#via-homebrew-for-macos)
  - [APT](#via-apt-for-debian-based-linux-distros)
  - [GitHub](#from-github-release)
- [Documentation](#documentation)
  - [Usage](#usage)
- [Examples](#examples)
  - [Config](#example-Lou-config)
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
sudo apt install Lou
```



<!-- TODO: -->
**Not currently working**

### From Github release

<!-- TODO: -->
**Not currently working**



## Documentation

### Usage

Use `Lou -h` or `Lou --help` to display help on commandline.

```
Lou, personal assitant at your service

Usage:
  Lou [command]

Available Commands:
  help                  Help about any command

Flags:
  -h, --help            help for Lou
  -o, --outDir string   Output directory. creates if not exitst

Use "Lou [command] --help" for more information about a command.
```



## Examples

Lou help
Lou clean -l $(pwd)
Lou biblo reformat



## Acknowledgements



## License

Lou is distributed under the terms of the GNU GENERAL PUBLIC LICENSE.

See [LICENSE](LICENSE) for details.


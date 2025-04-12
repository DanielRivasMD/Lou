

# Lou is a personal assitant to automatize routine tasks

**Lou** is a lightweight CLI tool for streamlining local file operations. Perform batch actions, deduplicate files, and automate mundane tasks with simple commands.

![Lou Logo](./assets/Lou01.png)

---

## Features

- **Smart Copy/Move**: Preserve structure or flatten directories
- **Duplicate Detection**: Find and remove redundant files (MD5 checksum)
- **Batch Renaming**: Regex support for pattern-based renaming
- **Dry Run Mode**: Preview changes before executing
- **Cross-Platform**: Windows/macOS/Linux support

## Installation

### Pre-built Binaries
Download for your OS from [Releases](https://github.com/yourusername/lou/releases).

### From Source (Go 1.20+)
```bash
git clone https://github.com/yourusername/lou.git
cd lou
go build -o lou cmd/main.go## Documentation
```

```
go install github.com/yourusername/lou@latest
```

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

Contributing
Fork the repository

Create a feature branch

Submit a pull request

License
MIT License - See LICENSE for details.

Contact
Daniel Rivas, MD
danielrivasmd@gmail.com

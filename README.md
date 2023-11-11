# Ignite version manager
Makes dealing with different [ignite](https://github.com/ignite/cli) versions easy

## Installation
First remove all instances of ignite, you have already installed, then
```bash
go install github.com/lxgr-linux/ignite-vm@v0.1.1
```

## Usage
```bash
ignite-vm list  # Lists all versions
ignite-vm install v0.27.1  # Downloads and installs ignite v0.27.1
ignite-vm set v0.27.1  # Sets ignite to version v0.27.1
```

## Contributing
Just create an issue or an PR
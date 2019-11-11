#!/bin/bash

# switch current working directory to script path
cd "$(dirname "$0")"

# Install homebrew if it's not installed yet
hash brew || /usr/bin/ruby -e "$(curl -fsSL https://raw.githubusercontent.com/Homebrew/install/master/install)"

# installation based on Brewfile
brew bundle -v

# installs Go Language Server
GO111MODULE=on go get golang.org/x/tools/gopls@latest

# installs Go Debugger
GO111MODULE=on go get -u github.com/go-delve/delve/cmd/dlv

# installs Go extension for Go
code --install-extension ms-vscode.go

tput setaf 2 && echo 'Go Lang dev setup is ready!!!' && tput sgr0
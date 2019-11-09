#!/bin/bash

cd "$(dirname "$0")"

hash brew || /usr/bin/ruby -e "$(curl -fsSL https://raw.githubusercontent.com/Homebrew/install/master/install)"
brew bundle -v

GO111MODULE=on go get golang.org/x/tools/gopls@latest
GO111MODULE=on go get -u github.com/go-delve/delve/cmd/dlv

code --install-extension ms-vscode.go

tput setaf 2 && echo 'Go Lang dev setup is ready!!!' && tput sgr0
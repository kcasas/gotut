#!/bin/bash

# switch current working directory to script path
cd "$(dirname "$0")"

# Install homebrew if it's not installed yet
hash brew || /usr/bin/ruby -e "$(curl -fsSL https://raw.githubusercontent.com/Homebrew/install/master/install)"

# installation based on Brewfile
brew bundle -v

# globally enable go modules
cat ~/.bash_profile | grep GO111MODULE || echo 'export GO111MODULE=on' >> ~/.bash_profile
source ~/.bash_profile

# installs VS Code extension for Go
code --install-extension ms-vscode.go

# installs Go Development modules used by extension
# reference: https://github.com/Microsoft/vscode-go/wiki/Go-tools-that-the-Go-extension-depends-on
go get -u -v github.com/ramya-rao-a/go-outline
go get -u -v github.com/acroca/go-symbols
go get -u -v golang.org/x/lint/golint
go get -u -v github.com/fatih/gomodifytags
go get -u -v golang.org/x/tools/cmd/gorename
go get -u -v github.com/cweill/gotests/...
go get -u -v golang.org/x/tools/cmd/guru
go get -u -v github.com/josharian/impl
go get -u -v github.com/haya14busa/goplay/cmd/goplay
go get -u -v github.com/uudashr/gopkgs/cmd/gopkgs
go get -u -v github.com/davidrjenni/reftools/cmd/fillstruct
go get -u -v github.com/go-delve/delve/cmd/dlv
# install Go Language Server
go get -v golang.org/x/tools/gopls@latest

tput setaf 2 && echo 'Go Lang dev setup is ready!!!' && tput sgr0
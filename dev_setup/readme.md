
# Development Setup

## Installation

run `$ ./setup.sh`

script will install these for you:

- [Homebrew](https://brew.sh/), Mac OS package manager.
- [Visual Studio Code](https://code.visualstudio.com/), recommended GoLang IDE.
- [Go VS Code extension](https://marketplace.visualstudio.com/items?itemName=ms-vscode.Go)
- [Go Lang v1.13](https://golang.org/doc/go1.13)
- [dlv], debugger tool
- [gopls](https://github.com/golang/tools/tree/master/gopls), Go Language Server.

## Language Server

Service that provides language development features like auto completion, go to definition, linter, formatter and the like to IDEs.

- collates language development features on a single tool
- behaviour of tooling is more consistent between IDEs
- saves development setup time
- processing is outsourced to the server
- ensures development follows the standard


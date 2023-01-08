# go-dev-containers-template

This Dev Container is created based on [microsoft/vscode-remote-try-go](https://github.com/microsoft/vscode-remote-try-go) and [microsoft/vscode-dev-containers](https://github.com/microsoft/vscode-dev-containers).

If you want to use databases/any other middlewares like postgres or redis, you can set it up with docker-compose. In detail, check [https://containers.dev/guide/dockerfile](https://containers.dev/guide/dockerfile).

## Installed Tools

These tools are already installed.

### Common

- git

### Languages

- Go(1.19)
- NVM(Node.js can be installed manually)
- yarn

### Packages

#### Go

- dlv
- golangci-lint
- golint
- gomodifytags
- go-outline
- gopkgs
- goplay
- gopls
- gotests
- impl
- revive
- staticcheck

#### JS

- prettier

## Usage

### 1) Clone Repo

```
$ git clone git@github.com:YuheiNakasaka/go-dev-containers-template.git ./your-project
```

### 2) Modify go.mod

```
# module github.com/YuheiNakasaka/go-dev-containers-template
module github.com/YOUR-NAME/YOUR-PROHECT

go 1.19
```

Also the same modification should be applied to `server.go`.

### 3) Reopen

Reopen the VSCode and waiting for a while. You will get the environment for golang development.

## License

MIT

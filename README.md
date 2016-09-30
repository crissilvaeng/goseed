# Goddamned

![license](https://img.shields.io/badge/license-MIT-blue.svg) ![go-version](https://img.shields.io/badge/go%20version-1.7-orange.svg) ![go-vendor](https://img.shields.io/badge/dependencies-go%20vendor-ff69b4.svg)

The objective of this project is to practice writing applications Go. Covering aspects such as testing, documentation, restful api and deploys in cloud environments.

## Use from the source

Assuming a recent version installed of Go and the variables `GOROOT` and `GOPATH` already defined, the set of commands below should recreate this project locally on a machine. Otherwise, [this installation guide](https://golang.org/doc/install) available in official of the [Golang](https://golang.org) site will help.

In addition, to manage possible dependencies this project uses [Godep](https://github.com/tools/godep) whose [installation instructions are available here](https://github.com/tools/godep).

**On Windows**

```bash
> go get -u github.com/crissilvaeng/goddamned
> cd %GOPATH%\github.com\crissilvaeng\goddamned
> godep restore
> go get
> goddamned
```

**On Linux**

```bash
$ go get -u github.com/crissilvaeng/goddamned
$ cd $GOPATH/github.com/crissilvaeng/goddamned
$ godep restore
$ go get
$ goddamned
```

## Change Log

To see the changes in this project, see the [change log file](CHANGELOG.md).

## Contribution

To contribute to this project, see the file [contribution tips](CONTRIBUTING.md) and [code of conduct](CONDUCT.md).

## License

The MIT License (MIT). To see the details of this license, see the [license file](LICENSE.md).

:octocat: